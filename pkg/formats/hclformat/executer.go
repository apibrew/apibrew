package hclformat

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"io"
	"os"
	"strings"
)

type executor struct {
	params      ExecutorParams
	schema      *hcl.BodySchema
	resources   *stub.ListResourceResponse
	evalContext *hcl.EvalContext
}

func (e *executor) Restore(ctx context.Context, file *os.File) error {
	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	hclFile, diagnostics := hclsyntax.ParseConfig(data, file.Name(),
		hcl.Pos{Line: 1, Column: 1, Byte: 0})
	if diagnostics != nil && diagnostics.HasErrors() {
		e.reportHclErrors(diagnostics)
		return errors.New("invalid Hcl file")
	}

	bodyContent, diagnostics := hclFile.Body.Content(e.schema)

	if diagnostics != nil && diagnostics.HasErrors() {
		e.reportHclErrors(diagnostics)

		return errors.New("invalid Hcl file")
	}

	for _, blocks := range bodyContent.Blocks.ByType() {
		for _, block := range blocks {
			var msg proto.Message
			msg, err = e.ParseBlock(block)

			if err != nil {
				return err
			}

			data, err := protojson.Marshal(msg)

			log.Println(string(data))

			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

func (e *executor) ParseBlock(block *hcl.Block) (proto.Message, error) {
	// check system resources
	found := false
	for _, resource := range resources.GetAllSystemResources() {
		resourceMessage := resources.GetSystemResourceType(resource).ProtoReflect().New().Interface()
		blockType := util.ToSnakeCase(string(resources.GetSystemResourceType(resource).ProtoReflect().Descriptor().Name()))

		if blockType == block.Type {
			found = true
			err := e.parseBlock(block, resourceMessage)

			if err != nil {
				return nil, err
			}

			return resourceMessage, nil
		}
	}

	if !found {
		log.Println(block.Type + " not found")
	}
	return nil, nil
}

func (e *executor) parseBlock(block *hcl.Block, resourceMessage protoreflect.ProtoMessage) error {
	pr := resourceMessage.ProtoReflect()

	resourceSchema := prepareSystemResourceSchema(resourceMessage.ProtoReflect().Descriptor())
	bodyContent, diagnostics := block.Body.Content(resourceSchema)

	if diagnostics != nil && diagnostics.HasErrors() {
		e.reportHclErrors(diagnostics)
		return errors.New("invalid Hcl file")
	}

	fields := pr.Descriptor().Fields()

	for _, attr := range bodyContent.Attributes {
		field := fields.ByName(protoreflect.Name(util.SnakeCaseToCamelCase(attr.Name)))

		value, diags := attr.Expr.Value(e.evalContext)
		attr.Expr.Variables()

		if diags != nil {
			e.reportHclErrors(diags)
			return errors.New("invalid Hcl file")
		}

		val, err := e.getValue(value, field, pr.NewField(field))
		if err != nil {
			return err
		}
		pr.Set(field, val)
	}

	labelI := 0
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		hclBlock := proto.GetExtension(field.Options(), model.E_HclBlock)
		hclLabel := proto.GetExtension(field.Options(), model.E_HclLabel)

		if hclLabel != nil && hclLabel != "" {
			resourceMessage.ProtoReflect().Set(field, protoreflect.ValueOfString(block.Labels[labelI]))
			labelI++
		}

		if hclBlock != nil && hclBlock != "" {
			// locating block
			for _, subBlock := range bodyContent.Blocks {
				if subBlock.Type == hclBlock.(string) {
					if field.IsList() {
						l := pr.Get(field).List()

						if !l.IsValid() {
							l = pr.NewField(field).List()
						}

						subMessage := l.NewElement()

						err := e.parseBlock(subBlock, subMessage.Message().Interface())

						if err != nil {
							return err
						}

						l.Append(subMessage)

						pr.Set(field, protoreflect.ValueOfList(l))
					} else if field.IsMap() {

					} else {
						subMessage := pr.NewField(field).Message().Interface()
						err := e.parseBlock(subBlock, subMessage)

						if err != nil {
							return err
						}

						if subMessage != nil {
							resourceMessage.ProtoReflect().Set(field, protoreflect.ValueOfMessage(subMessage.ProtoReflect()))
						}
					}
				}
			}
		}
	}

	return nil
}

func (e *executor) getValue(value cty.Value, fieldDescriptor protoreflect.FieldDescriptor, newValue protoreflect.Value) (protoreflect.Value, error) {
	// Handle optional fields
	if fieldDescriptor.IsList() && value.IsNull() {
		return protoreflect.ValueOf(nil), nil
	}

	fieldType := fieldDescriptor.Kind()

	// Decode the value based on the field type
	switch fieldType {
	case protoreflect.BoolKind:
		if err := e.requireType(value, cty.Bool); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		return protoreflect.ValueOf(value.True()), nil
	case protoreflect.StringKind:
		if err := e.requireType(value, cty.String); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		return protoreflect.ValueOf(value.AsString()), nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if err := e.requireType(value, cty.Number); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		val, _ := value.AsBigFloat().Int64()
		return protoreflect.ValueOf(int32(val)), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if err := e.requireType(value, cty.Number); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		val, _ := value.AsBigFloat().Uint64()
		return protoreflect.ValueOf(uint32(val)), nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if err := e.requireType(value, cty.Number); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		val, _ := value.AsBigFloat().Int64()
		return protoreflect.ValueOf(val), nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if err := e.requireType(value, cty.Number); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		val, _ := value.AsBigFloat().Uint64()
		return protoreflect.ValueOf(val), nil
	case protoreflect.FloatKind:
		if err := e.requireType(value, cty.Number); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		val, _ := value.AsBigFloat().Float64()
		return protoreflect.ValueOf(float32(val)), nil
	case protoreflect.DoubleKind:
		if err := e.requireType(value, cty.Number); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		val, _ := value.AsBigFloat().Float64()
		return protoreflect.ValueOf(val), nil
	case protoreflect.BytesKind:
		if err := e.requireType(value, cty.String); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		return protoreflect.ValueOfBytes([]byte(value.AsString())), nil
	case protoreflect.EnumKind:
		if err := e.requireType(value, cty.String); err != nil {
			return protoreflect.ValueOf(nil), err
		}

		enumValueDescriptor := fieldDescriptor.Enum().Values().ByName(protoreflect.Name(strings.ToUpper(value.AsString())))
		if enumValueDescriptor == nil {
			return protoreflect.Value{}, fmt.Errorf("Invalid enum value: %q", value.AsString())
		}
		return protoreflect.ValueOfEnum(enumValueDescriptor.Number()), nil
	default:
		return protoreflect.Value{}, fmt.Errorf("Unsupported field type: %v", fieldType)
	}
}

func (e *executor) init() error {
	list, err := e.params.ResourceClient.List(context.TODO(), &stub.ListResourceRequest{})

	if err != nil {
		return err
	}

	e.resources = list

	e.schema = prepareSchema(list.Resources)

	e.evalContext = &hcl.EvalContext{
		Variables: nil,
		Functions: map[string]function.Function{
			"specialProperties": function.New(&function.Spec{
				Description: "",
				Params:      []function.Parameter{},
				VarParam:    &function.Parameter{},
				Type:        function.StaticReturnType(cty.List(cty.Object(nil))),
				Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
					return cty.ListVal([]cty.Value{
						cty.ObjectVal(map[string]cty.Value{
							//"name": cty.StringVal("EXAMPLE_STRING_VALUE_11123"),
						}),
					}), nil
				},
			}),
		},
	}

	return nil
}

func (e *executor) reportHclErrors(diagnostics hcl.Diagnostics) {
	for _, item := range diagnostics {
		log.Error(item.Error(), item.Summary)
	}

}

func (e *executor) requireType(value cty.Value, typ cty.Type) error {
	if typ != value.Type() {
		return errors.New(fmt.Sprintf("%s expected but %q found", typ.GoString(), value.Type().GoString()))
	}

	return nil
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ExecutorParams struct {
	Input          io.Reader
	ResourceClient stub.ResourceClient
	RecordClient   stub.RecordClient
	OverrideConfig OverrideConfig
	Token          string
	DoMigration    bool
	ForceMigration bool
	DataOnly       bool
}

func NewExecutor(params ExecutorParams) (Executor, error) {
	exec := &executor{
		params: params,
	}

	return exec, exec.init()
}
