package output

import (
	"encoding/json"
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v2"
	"io"
)

var jsonMo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: true,
}

type yamlReader struct {
	writer            io.Writer
	hasMessageWritten bool
}

func (c *yamlReader) IsBinary() bool {
	return false
}

func (c *yamlReader) WriteResources(resources []*model.Resource) {
	for _, resource := range resources {
		c.writePrefix()
		body, err := jsonMo.Marshal(resource)

		check(err)

		var data map[string]interface{}

		err = json.Unmarshal(body, &data)

		data["type"] = "resource"

		check(err)

		out, err := yaml.Marshal(data)

		check(err)

		_, err = c.writer.Write(out)

		check(err)
	}
}

func (c *yamlReader) WriteRecords(resource *model.Resource, recordsChan chan *model.Record) {
	for record := range recordsChan {
		c.writePrefix()
		body, err := jsonMo.Marshal(record)

		check(err)

		var data map[string]interface{}

		err = json.Unmarshal(body, &data)

		data["type"] = "record"
		data["namespace"] = resource.Namespace
		data["resource"] = resource.Name

		check(err)

		out, err := yaml.Marshal(data)

		check(err)

		_, err = c.writer.Write(out)

		check(err)
	}
}

func (c *yamlReader) writePrefix() {
	if c.hasMessageWritten {
		c.writer.Write([]byte("---\n"))
	}

	c.hasMessageWritten = true
}
