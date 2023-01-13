package output

import (
	"data-handler/model"
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v2"
	"io"
)

var jsonMo = protojson.MarshalOptions{
	Multiline:       true,
	EmitUnpopulated: true,
}

type yamlWriter struct {
	writer            io.Writer
	hasMessageWritten bool
}

func (c *yamlWriter) DescribeResource(resource *model.Resource) {
	c.WriteResources([]*model.Resource{resource})
}

func (c *yamlWriter) WriteResources(resources []*model.Resource) {
	for _, resource := range resources {
		c.writePrefix()
		body, err := jsonMo.Marshal(resource)

		check(err)

		var data interface{}

		err = json.Unmarshal(body, &data)

		check(err)

		out, err := yaml.Marshal(data)

		check(err)

		_, err = c.writer.Write(out)

		check(err)
	}
}

func (c *yamlWriter) WriteRecords(record []*model.Record) {

}

func (c *yamlWriter) writePrefix() {
	if c.hasMessageWritten {
		c.writer.Write([]byte("---\n"))
	}

	c.hasMessageWritten = true
}
