package ops

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type preprocessor struct {
	dhClient client.DhClient
	writer   *Writer
}

func (p *preprocessor) preprocess(body unstructured.Unstructured) (unstructured.Unstructured, error) {
	visitedBody, err := WalkUnstructured(body, func(value interface{}) (interface{}, error) {
		if un, ok := value.(unstructured.Unstructured); ok {
			for key := range un {
				if isPreprocessorKey(key) {
					return p.runPreprocess(value.(unstructured.Unstructured))
				}
			}
		}

		return value, nil
	})

	if err != nil {
		return nil, err
	}

	return visitedBody.(unstructured.Unstructured), nil
}

func (p *preprocessor) runPreprocess(un unstructured.Unstructured) (interface{}, error) {
	var err error
	var keys = unstructured.Keys(un)

	if util.ArrayContains(keys, "$file") {
		return p.runIncludeFile(un)
	}

	if util.ArrayContains(keys, "$folder") {
		return p.runIncludeFolder(un)
	}

	if util.ArrayContains(keys, "$extend") {
		un, err = p.runPreprocessExtend(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(un)
	}

	if util.ArrayContains(keys, "$include") {
		un, err = p.runPreprocessInclude(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(un)
	}

	if util.ArrayContains(keys, "$properties") {
		un, err = p.runPreprocessProperties(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(un)
	}

	if util.ArrayContains(keys, "$override") {
		un, err = p.runPreprocessOverride(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(un)
	}

	if util.ArrayContains(keys, "$syntax") {
		err = p.checkSyntax(un)

		if err != nil {
			return nil, err
		}

		return un, nil
	}

	return un, nil
}

func (p *preprocessor) runPreprocessInclude(un unstructured.Unstructured) (unstructured.Unstructured, error) {
	return un, nil
}

func (p *preprocessor) runIncludeFile(un unstructured.Unstructured) (string, error) {
	filePath := un["$file"]

	dat, err := os.ReadFile(filePath.(string))

	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func (p *preprocessor) runIncludeFolder(un unstructured.Unstructured) (string, error) {
	filePath := un["$folder"].(string)
	format := un["$format"].(string)

	if format == "tar" {
		var buffer = new(bytes.Buffer)
		err := makeTar(filePath, buffer)

		if err != nil {
			return "", err
		}

		return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
	} else {
		return "", errors.New("unsupported format: " + format)
	}
}

func (p *preprocessor) runPreprocessExtend(un unstructured.Unstructured) (unstructured.Unstructured, error) {
	extend := un["$extend"].(unstructured.Unstructured)

	namespace := extend["$namespace"].(string)
	resource := extend["$resource"].(string)
	match := extend["$match"].(unstructured.Unstructured)

	if namespace != "" {
		namespace = "ui" //default
	}

	var filters = make(map[string]string)

	if match != nil {
		for key, value := range match {
			filters[key] = value.(string)
		}
	}

	resp, err := p.dhClient.GetRecordClient().List(context.TODO(), &stub.ListRecordRequest{
		Namespace: namespace,
		Resource:  resource,
		Filters:   filters,
	})

	if err != nil {
		reportError(un, err)
		return nil, err
	}

	ref := extend["$ref"].(string)

	if len(resp.Content) == 0 {
		return un, errors.New(fmt.Sprintf("no record found for %s/%s/%v", namespace, resource, match))
	}

	recordUn, err := p.writer.WriteRecord(namespace, resource, resp.Content[0])

	if err != nil {
		reportError(un, err)
		return un, err
	}

	unstructured.MergeOut(un, recordUn, true)

	if ref != "" {
		updatedUn, err := ParseRef(un, ref)

		if err != nil {
			reportError(un, err)
			return nil, err
		}

		// keep directives
		for key, value := range un {
			if isPreprocessorKey(key) {
				updatedUn[key] = value
			}
		}

		un = updatedUn
	}

	unstructured.DeleteKey(un, "$extend")

	return un, nil
}

func reportError(un unstructured.Unstructured, err error) {
	data, _ := json.MarshalIndent(un, " ", "  ")
	fmt.Println(string(data))
	log.Error(err)
}

func (p *preprocessor) runPreprocessOverride(un unstructured.Unstructured) (unstructured.Unstructured, error) {
	override := un["$override"].(unstructured.Unstructured)
	selectParam := override["$select"].(string)
	merge, _ := override["$merge"].(unstructured.Unstructured)
	set, _ := override["$set"].(unstructured.Unstructured)

	reUn, err := ParseRef(un, selectParam)

	if err != nil {
		return nil, err
	}

	if merge != nil {
		unstructured.MergeInto(reUn, merge, true)
		unstructured.DeleteKey(un, "$merge")
	}

	if set != nil {
		unstructured.MergeInto(reUn, set, false)
		unstructured.DeleteKey(un, "$set")
	}

	unstructured.DeleteKey(un, "$override")

	return un, nil
}

func (p *preprocessor) checkSyntax(un unstructured.Unstructured) error {
	if un["$syntax"] != nil {
		syntax, err := p.runPreprocess(un["$syntax"].(unstructured.Unstructured))

		if err != nil {
			return err
		}

		unstructured.DeleteKey(un, "$syntax")

		var subType = &model.Resource{}

		err = unstructured.ToProtoMessage(syntax.(unstructured.Unstructured), subType)

		if err != nil {
			return err
		}

		record, err := unstructured.ToRecord(un)

		if err != nil {
			return err
		}

		err = validate.Records(subType, []*model.Record{record}, false)

		// recover syntax as it needs to be persisted
		un["$syntax"] = syntax

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *preprocessor) runPreprocessProperties(un unstructured.Unstructured) (unstructured.Unstructured, error) {
	propertiesDirective := un["$properties"].(unstructured.Unstructured)
	var properties []unstructured.Unstructured

	for key, value := range propertiesDirective {
		var propertyUn = make(unstructured.Unstructured)

		if typeStr, ok := value.(string); ok {
			var property = new(model.ResourceProperty)

			if typeId, ok := model.ResourceProperty_Type_value[strings.ToUpper(typeStr)]; ok {
				property.Type = model.ResourceProperty_Type(typeId)
			} else {
				return nil, errors.New(fmt.Sprintf("invalid property type %s", typeStr))
			}
			err := unstructured.FromProtoMessage(propertyUn, property)

			if err != nil {
				return nil, err
			}
		} else if valueUn, ok := value.(unstructured.Unstructured); ok {
			propertyUn = valueUn
		} else {
			return nil, errors.New(fmt.Sprintf("invalid property type %s", value))
		}

		propertyUn["name"] = key

		if strings.HasSuffix(key, "!") {
			propertyUn["name"] = key[:len(key)-1]
			propertyUn["required"] = true
		}

		properties = append(properties, propertyUn)
	}

	unstructured.DeleteKey(un, "$properties")
	un["properties"] = properties
	return un, nil
}

var preprocessKeywords = []string{
	"$extend", "$override", "$select", "$syntax", "$ref", "$merge", "$set", "$clear", "$append", "$include", "$expression", "$properties", "$file", "$folder",
}

func isPreprocessorKey(key string) bool {
	for _, keyword := range preprocessKeywords {
		if key == keyword {
			return true
		}
	}

	return false
}
