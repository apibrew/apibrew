package flags

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"regexp"
	"strings"
)

type protoMessageParserFlags[T proto.Message] struct {
	pr               protoreflect.Message
	arrIndexMatchMap map[*regexp.Regexp]string
	normalizedFlags  map[string]bool
}

func (p *protoMessageParserFlags[T]) Declare(cmd *cobra.Command) {
	p.prepareFlags("", cmd, p.pr.Descriptor().Fields())
	cmd.PersistentFlags().SetNormalizeFunc(func(f *pflag.FlagSet, flagName string) pflag.NormalizedName {
		if p.normalizedFlags[flagName] {
			return pflag.NormalizedName(flagName)
		}

		p.normalizedFlags[flagName] = true

		for pattern, actualFlagName := range p.arrIndexMatchMap {
			if pattern.MatchString(flagName) && flagName != actualFlagName {
				var foundFlag pflag.Flag
				f.VisitAll(func(flag *pflag.Flag) {
					if flag.Name == actualFlagName {
						foundFlag = *flag
					}
				})

				typ := foundFlag.Value.Type()
				switch typ {
				case "string":
					f.String(flagName, "", "")
				case "bool":
					f.Bool(flagName, false, "")
				case "uint32":
					f.Uint32(flagName, 0, "")
				default:
					log.Fatal("Unknown type: " + foundFlag.Value.Type())
				}
			}
		}

		return pflag.NormalizedName(flagName)
	})
}

func (p *protoMessageParserFlags[T]) prepareFlags(path string, cmd *cobra.Command, fields protoreflect.FieldDescriptors) {
	var calls []func()

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		name := util.ToDashCase(string(field.Name()))
		kind := field.Kind().String()
		msg := field.Message()

		var fullName = name
		if msg != nil {
			fullName = string(msg.FullName())
		}

		if path != "" {
			name = path + "-" + name
		}

		if name == "audit-data" || name == "data-type" || name == "version" {
			continue
		}

		if strings.Contains(name, "[index]") {
			p.arrIndexMatchMap[regexp.MustCompile(strings.Replace("^"+name+"$", "[index]", "\\d", -1))] = name
		}

		if fullName == "google.protobuf.Timestamp" {
			cmd.PersistentFlags().String(name, "", "")
		} else if fullName == "google.protobuf.Value" {
			cmd.PersistentFlags().String(name, "", "")
		} else {
			switch kind {
			case "string":
				cmd.PersistentFlags().String(name, "", "")
			case "enum":
				var names []string
				for j := 0; j < field.Enum().Values().Len(); j++ {
					names = append(names, strings.ToLower(string(field.Enum().Values().Get(j).Name())))
				}
				cmd.PersistentFlags().String(name, "", strings.Join(names, ", "))
			case "message":
				if field.IsList() {
					calls = append(calls, func() {
						newName := name + "-[index]"
						p.prepareFlags(newName, cmd, field.Message().Fields())
					})
				} else if field.IsMap() {
					cmd.PersistentFlags().String(name, "", "")
				} else {
					calls = append(calls, func() {
						p.prepareFlags(name, cmd, field.Message().Fields())
					})
				}
			case "bool":
				cmd.PersistentFlags().Bool(name, false, "")
			case "uint32":
				cmd.PersistentFlags().Uint32(name, 0, "")
			case "int64":
				cmd.PersistentFlags().Int64(name, 0, "")
			case "int32":
				cmd.PersistentFlags().Int32(name, 0, "")
			case "double":
				cmd.PersistentFlags().Float64(name, 0, "")
			default:
				log.Fatal("Unknown type: " + kind)
			}
		}
	}

	for _, call := range calls {
		call()
	}
}

func (p *protoMessageParserFlags[T]) Parse(elem T, cmd *cobra.Command, args []string) {
	p.parseLocal("", p.pr.Descriptor().Fields(), elem, cmd, args)
}

func (p *protoMessageParserFlags[T]) parseLocal(path string, fields protoreflect.FieldDescriptors, elem proto.Message, cmd *cobra.Command, args []string) {
	var l = fields.Len()
	for i := 0; i < l; i++ {
		field := fields.Get(i)
		name := util.ToDashCase(string(field.Name()))
		kind := field.Kind().String()
		msg := field.Message()

		var fullName = name
		if msg != nil {
			fullName = string(msg.FullName())
		}

		if path != "" {
			name = path + "-" + name
		}

		if name == "audit-data" || name == "data-type" || name == "version" {
			continue
		}

		if fullName == "google.protobuf.Timestamp" {
			val, err := cmd.PersistentFlags().GetString(name)

			if err != nil {
				continue
			}

			elem.ProtoReflect().Set(field, protoreflect.ValueOfString(val))
		} else if fullName == "google.protobuf.Value" {
			_, err := cmd.PersistentFlags().GetString(name)

			if err != nil {
				continue
			}

			//elem.ProtoReflect().Set(field, protoreflect.ValueOfString(val))
		} else {
			switch kind {
			case "string":
				val, err := cmd.PersistentFlags().GetString(name)

				if err != nil {
					continue
				}

				elem.ProtoReflect().Set(field, protoreflect.ValueOfString(val))
			case "enum":
				valX, err := cmd.PersistentFlags().GetString(name)

				if err != nil {
					continue
				}

				instance := field.Enum().Values().ByName(protoreflect.Name(strings.ToUpper(valX)))

				if valX != "" && instance == nil {
					log.Fatalf("Unknown value for enum: %s => %s", field.Enum().Name(), valX)
				}

				if instance != nil {
					elem.ProtoReflect().Set(field, protoreflect.ValueOf(instance.Number()))
				}
			case "message":
				var i = 0
				if field.IsList() {
					var l = elem.ProtoReflect().NewField(field).List()

					for {
						newName := fmt.Sprintf("%s-%d", name, i)

						var found = false
						cmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
							if strings.HasPrefix(flag.Name, newName) {
								found = true
							}
						})

						if found {
							el := l.NewElement()
							instance := el.Message().Interface()

							p.parseLocal(newName, field.Message().Fields(), instance, cmd, args)

							l.Append(el)
						} else {
							break
						}

						i++
					}
					elem.ProtoReflect().Set(field, protoreflect.ValueOfList(l))
				} else if field.IsMap() {
					var l = elem.ProtoReflect().NewField(field).Map()

					elem.ProtoReflect().Set(field, protoreflect.ValueOfMap(l))
				} else {
					var found = false
					var newName = "--" + name + "-"
					for _, arg := range args {
						if strings.HasPrefix(arg, newName) {
							found = true
						}
					}

					if found {
						var l = elem.ProtoReflect().NewField(field)
						instance := l.Message().Interface()
						p.parseLocal(name, field.Message().Fields(), instance, cmd, args)
						elem.ProtoReflect().Set(field, l)
					}
				}
			case "bool":
				val, err := cmd.Flags().GetBool(name)

				if err != nil {
					continue
				}

				elem.ProtoReflect().Set(field, protoreflect.ValueOf(val))
			case "uint32":
				val, err := cmd.PersistentFlags().GetUint32(name)

				if err != nil {
					continue
				}

				elem.ProtoReflect().Set(field, protoreflect.ValueOf(val))
			case "int64":
				val, err := cmd.PersistentFlags().GetInt64(name)

				if err != nil {
					continue
				}

				elem.ProtoReflect().Set(field, protoreflect.ValueOf(val))
			case "int32":
				val, err := cmd.PersistentFlags().GetInt32(name)

				if err != nil {
					continue
				}

				elem.ProtoReflect().Set(field, protoreflect.ValueOf(val))
			case "double":
				val, err := cmd.PersistentFlags().GetFloat64(name)

				if err != nil {
					continue
				}

				elem.ProtoReflect().Set(field, protoreflect.ValueOf(val))
			default:
				log.Fatal("Unknown type: " + kind)
			}
		}
	}
}

func NewProtoMessageParserFlags[T proto.Message](pr protoreflect.Message) FlagHelper[T] {
	return &protoMessageParserFlags[T]{
		pr:               pr,
		arrIndexMatchMap: make(map[*regexp.Regexp]string),
		normalizedFlags:  make(map[string]bool),
	}
}
