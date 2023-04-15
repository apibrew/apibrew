package hclformat

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/client"
	"github.com/tislib/data-handler/pkg/formats"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/resources"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/stub"
	"github.com/tislib/data-handler/pkg/util"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"os"
	"strings"
)

type executor struct {
	params      ExecutorParams
	resources   []*model.Resource
	evalContext *hcl.EvalContext
}

func (e *executor) Restore(ctx context.Context, file *os.File) error {
	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	hclFile, diags := hclsyntax.ParseConfig(data, file.Name(),
		hcl.Pos{Line: 1, Column: 1, Byte: 0})
	if diags != nil && diags.HasErrors() {
		e.reportHclErrors(diags)
		return errors.New("invalid Hcl file")
	}

	bodyContent, diags := hclFile.Body.Content(prepareRootSchema())

	if diags != nil && diags.HasErrors() {
		e.reportHclErrors(diags)

		return errors.New("invalid Hcl file")
	}

	for _, blocks := range bodyContent.Blocks.ByType() {
		for _, block := range blocks {
			err = e.ApplyRootBlock(ctx, block)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *executor) ApplyRootBlock(ctx context.Context, block *hcl.Block) error {
	var bodyContent *hcl.BodyContent
	var diags hcl.Diagnostics

	switch block.Type {
	case "schema":
		{
			bodyContent, diags = block.Body.Content(prepareSchemaSchema())

			if diags != nil && diags.HasErrors() {
				e.reportHclErrors(diags)

				return errors.New("invalid Hcl file")
			}
		}
	case "data":
		{
			if err := e.prepareResources(); err != nil {
				return err
			}

			bodyContent, diags = block.Body.Content(prepareDataSchema(e.resources))

			if diags != nil && diags.HasErrors() {
				e.reportHclErrors(diags)

				return errors.New("invalid Hcl file")
			}
		}
	}

	for _, blocks := range bodyContent.Blocks.ByType() {
		for _, item := range blocks {
			err := e.ApplyBlock(ctx, item)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *executor) ApplyBlock(ctx context.Context, block *hcl.Block) error {
	// check system resources
	for _, resource := range resources.GetAllSystemResources() {
		resourceMessage := resources.GetSystemResourceType(resource).ProtoReflect().New().Interface()
		blockType := util.ToSnakeCase(string(resources.GetSystemResourceType(resource).ProtoReflect().Descriptor().Name()))

		if blockType == block.Type {
			err := e.parseBlock(block, resourceMessage)

			if err != nil {
				return err
			}

			if resourceItem, ok := resourceMessage.(*model.Resource); ok {
				err = e.params.DhClient.ApplyResource(ctx, resourceItem, e.params.DoMigration, e.params.ForceMigration)
			} else {
				err = e.params.DhClient.Apply(ctx, resourceMessage)
			}

			if err != nil {
				log.Errorf("Cannot Apply: %v (%s/%s)", resourceMessage, resource.Namespace, resource.Name)
				return err
			}

			return nil
		}
	}

	if block.Type == "record" {
		namespace := block.Labels[0]
		resourceName := block.Labels[1]
		var resource *model.Resource

		for _, item := range e.resources {
			if item.Namespace == namespace && item.Name == resourceName {
				resource = item
			}
		}

		if resource == nil {
			return errors.New("Resource not found: " + resourceName)
		}

		record, err := e.parseBlockToRecord(block, resource, false)

		if err != nil {
			return err
		}

		err = e.params.DhClient.ApplyRecord(ctx, resource, record)

		if err != nil {
			log.Errorf("Cannot Apply record: %v (%s/%s)", record, resource.Namespace, resource.Name)
			return err
		}

		return nil
	} else {
		for _, resource := range e.resources {
			hclBlock := annotations.Get(resource, annotations.HclBlock)

			if hclBlock != "" && hclBlock == block.Type {
				record, err := e.parseBlockToRecord(block, resource, true)

				if err != nil {
					return err
				}

				err = e.params.DhClient.ApplyRecord(ctx, resource, record)

				if err != nil {
					log.Errorf("Cannot Apply record: %v (%s/%s)", record, resource.Namespace, resource.Name)
					return err
				}

				return nil
			}
		}
	}

	return errors.New("Unknown block: " + block.Type)
}

func (e *executor) parseBlockToRecord(block *hcl.Block, resource *model.Resource, parseLabels bool) (*model.Record, error) {
	var record = &model.Record{
		Properties: make(map[string]*structpb.Value),
	}

	bodyContent, diags := block.Body.Content(prepareResourceRecordSchema(resource, parseLabels))

	if diags != nil {
		e.reportHclErrors(diags)
		return nil, diags
	}

	for _, attr := range bodyContent.Attributes {
		var prop *model.ResourceProperty

		for _, item := range resource.Properties {
			if util.ToSnakeCase(item.Name) == attr.Name {
				prop = item
			}
		}

		val, diags := attr.Expr.Value(e.evalContext)

		if diags != nil {
			e.reportHclErrors(diags)
			return nil, diags
		}

		propVal, err := e.parseProperty(prop.Type, val)

		if err != nil {
			return nil, err
		}

		record.Properties[prop.Name] = propVal
	}

	for _, blockItem := range bodyContent.Blocks {
		var prop *model.ResourceProperty

		for _, item := range resource.Properties {
			if annotations.Get(item, annotations.HclBlock) == blockItem.Type {
				prop = item
			}
		}

		propVal, err := e.parseBlockProperty(prop, blockItem)

		if err != nil {
			return nil, err
		}

		record.Properties[prop.Name] = propVal
	}

	if parseLabels && len(block.Labels) > 0 {
		li := 0
		for _, item := range resource.Properties {
			hclLabel := annotations.IsEnabled(item, annotations.IsHclLabel)
			if hclLabel {
				record.Properties[item.Name] = structpb.NewStringValue(block.Labels[li])
				li++
			}
		}
	}

	return record, nil
}

func (e *executor) parseBlockProperty(prop *model.ResourceProperty, blockItem *hcl.Block) (*structpb.Value, error) {
	switch prop.Type {
	case model.ResourceProperty_OBJECT, model.ResourceProperty_STRUCT, model.ResourceProperty_REFERENCE:
		attributes, diags := blockItem.Body.JustAttributes()

		if diags != nil {
			e.reportHclErrors(diags)
			return nil, diags
		}

		var objData = make(map[string]interface{})

		for _, attr := range attributes {
			val, diags := attr.Expr.Value(e.evalContext)

			if diags != nil {
				e.reportHclErrors(diags)
				return nil, diags
			}

			objData[attr.Name] = val.AsString()
		}

		st, err := structpb.NewStruct(objData)

		if err != nil {
			return nil, err
		}

		return structpb.NewStructValue(st), nil
	default:
		panic("unknown property type: " + prop.Type.String())
	}
}

func (e *executor) parseBlock(block *hcl.Block, resourceMessage protoreflect.ProtoMessage) error {
	pr := resourceMessage.ProtoReflect()

	resourceSchema := prepareSystemResourceSchema(resourceMessage.ProtoReflect().Descriptor())
	bodyContent, diags := block.Body.Content(resourceSchema)

	if diags != nil && diags.HasErrors() {
		e.reportHclErrors(diags)
		return errors.New("invalid Hcl file")
	}

	fields := pr.Descriptor().Fields()

	for _, attr := range bodyContent.Attributes {
		field := fields.ByName(protoreflect.Name(util.SnakeCaseToCamelCase(attr.Name)))

		value, diags := attr.Expr.Value(e.evalContext)

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

		if hclLabel != "" {
			resourceMessage.ProtoReflect().Set(field, protoreflect.ValueOfString(block.Labels[labelI]))
			labelI++
		}

		if hclBlock != "" {
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
						attrs, diags := subBlock.Body.JustAttributes()

						if diags != nil && diags.HasErrors() {
							e.reportHclErrors(diags)
							return diags
						}
						newElem := pr.NewField(field)
						for key, value := range attrs {
							val, err := value.Expr.Value(e.evalContext)

							if err != nil {
								return err
							}
							newElem.Map().Set(protoreflect.ValueOf(key).MapKey(), protoreflect.ValueOf(val.AsString()))
						}
						pr.Set(field, newElem)
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

func (e *executor) parseProperty(resourcePropertyType model.ResourceProperty_Type, value cty.Value) (*structpb.Value, error) {
	switch resourcePropertyType {
	case model.ResourceProperty_INT64, model.ResourceProperty_INT32, model.ResourceProperty_FLOAT32, model.ResourceProperty_FLOAT64:
		if err := e.requireType(value, cty.Number); err != nil {
			return nil, err
		}

		val, _ := value.AsBigFloat().Float64()

		return structpb.NewNumberValue(val), nil
	case model.ResourceProperty_STRING, model.ResourceProperty_UUID, model.ResourceProperty_DATE, model.ResourceProperty_TIME, model.ResourceProperty_TIMESTAMP, model.ResourceProperty_ENUM, model.ResourceProperty_BYTES:
		if err := e.requireType(value, cty.String); err != nil {
			return nil, err
		}

		return structpb.NewStringValue(value.AsString()), nil
	case model.ResourceProperty_BOOL:
		if err := e.requireType(value, cty.Bool); err != nil {
			return nil, err
		}

		return structpb.NewBoolValue(value.True()), nil
	default:
		panic("unknown property type: " + resourcePropertyType.String())
	}
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

func (e *executor) prepareResources() error {
	resp, err := e.params.DhClient.GetResourceClient().List(context.TODO(), &stub.ListResourceRequest{})

	if err != nil {
		return err
	}

	e.resources = resp.Resources
	return nil
}

func (e *executor) reportHclErrors(diags hcl.Diagnostics) {
	for _, item := range diags {
		log.Error(item.Error(), item.Summary)
	}

}

func (e *executor) requireType(value cty.Value, typ cty.Type) error {
	if typ != value.Type() {
		return fmt.Errorf("%s expected but %q found", typ.GoString(), value.Type().GoString())
	}

	return nil
}

type OverrideConfig struct {
	Namespace  string
	DataSource string
}

type ExecutorParams struct {
	Input          io.Reader
	DhClient       client.DhClient
	OverrideConfig OverrideConfig
	Token          string
	DoMigration    bool
	ForceMigration bool
	DataOnly       bool
}

func NewExecutor(params ExecutorParams) (formats.Executor, error) {
	exec := &executor{
		params: params,
	}

	return exec, exec.init()
}
