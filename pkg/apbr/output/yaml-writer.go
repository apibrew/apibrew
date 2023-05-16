package output

import (
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/model"
	log "github.com/sirupsen/logrus"
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

func (c *yamlWriter) IsBinary() bool {
	return false
}

func (c *yamlWriter) WriteResources(resources []*model.Resource) {
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

func (c *yamlWriter) WriteRecords(resource *model.Resource, total uint32, recordsChan chan *model.Record) {
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

func (c *yamlWriter) writePrefix() {
	if c.hasMessageWritten {
		if _, err := c.writer.Write([]byte("---\n")); err != nil {
			log.Fatal(err)
		}
	}

	c.hasMessageWritten = true
}
