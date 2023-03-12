package flags

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
)

type recordParserFlags struct {
	resource  *model.Resource
	parserMap map[string]func() (*structpb.Value, error)
}

func funct[T any](p *T) func() (*structpb.Value, error) {
	return func() (*structpb.Value, error) {
		if p == nil {
			return nil, nil
		}
		return structpb.NewValue(*p)
	}
}

func (p *recordParserFlags) Declare(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	for _, prop := range p.resource.Properties {
		name := util.ToDashCase(prop.Name)

		switch prop.Type {
		case model.ResourceProperty_BOOL:
			p.parserMap[prop.Name] = funct(flags.Bool(name, false, ""))
		case model.ResourceProperty_STRING:
			p.parserMap[prop.Name] = funct(flags.String(name, "", ""))
		case model.ResourceProperty_FLOAT32:
			p.parserMap[prop.Name] = funct(flags.Float32(name, 0, ""))
		case model.ResourceProperty_FLOAT64:
			p.parserMap[prop.Name] = funct(flags.Float64(name, 0, ""))
		case model.ResourceProperty_INT32:
			p.parserMap[prop.Name] = funct(flags.Int32(name, 0, ""))
		case model.ResourceProperty_INT64:
			p.parserMap[prop.Name] = funct(flags.Int64(name, 0, ""))
		case model.ResourceProperty_BYTES:
			p.parserMap[prop.Name] = funct(flags.BytesHex(name, []byte{}, ""))
		case model.ResourceProperty_UUID:
			p.parserMap[prop.Name] = funct(flags.String(name, "", ""))
		case model.ResourceProperty_DATE:
			p.parserMap[prop.Name] = funct(flags.String(name, "", ""))
		case model.ResourceProperty_TIME:
			p.parserMap[prop.Name] = funct(flags.String(name, "", ""))
		case model.ResourceProperty_TIMESTAMP:
			p.parserMap[prop.Name] = funct(flags.String(name, "", ""))
		case model.ResourceProperty_OBJECT:
			p.parserMap[prop.Name] = funct(flags.String(name, "", "")) //todo fix
		case model.ResourceProperty_MAP:
			p.parserMap[prop.Name] = funct(flags.String(name, "", "")) //todo fix
		case model.ResourceProperty_LIST:
			p.parserMap[prop.Name] = funct(flags.String(name, "", "")) //todo fix
		case model.ResourceProperty_REFERENCE:
			p.parserMap[prop.Name] = funct(flags.String(name, "", "")) //todo fix
		case model.ResourceProperty_ENUM:
			p.parserMap[prop.Name] = funct(flags.String(name, "", "")) // todo fix
		}
	}
}

func (p *recordParserFlags) Parse(elem *model.Record, cmd *cobra.Command, args []string) {
	err := cmd.PersistentFlags().Parse(args)

	if err != nil {
		log.Fatal(err)
	}

	elem.Properties = make(map[string]*structpb.Value)

	for _, prop := range p.resource.Properties {
		name := util.ToDashCase(prop.Name)

		if !cmd.Flag(name).Changed {
			continue
		}

		val, err := p.parserMap[prop.Name]()

		if err != nil {
			log.Fatal(err)
		}

		elem.Properties[prop.Name] = val
	}
}

func NewRecordParserFlags(resource *model.Resource) FlagHelper[*model.Record] {
	return &recordParserFlags{
		resource:  resource,
		parserMap: make(map[string]func() (*structpb.Value, error)),
	}
}
