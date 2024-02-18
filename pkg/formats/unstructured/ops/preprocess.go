package ops

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/formats/jsonformat"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	writer2 "github.com/apibrew/apibrew/pkg/formats/writer"
	"github.com/apibrew/apibrew/pkg/formats/yamlformat"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

type preprocessor struct {
	dhClient client.Client
	writer   *writer2.Writer
}

func (p *preprocessor) preprocess(cfPath string, body unstructured.Unstructured) (interface{}, error) {
	visitedBody, err := WalkUnstructured(body, func(value interface{}) (interface{}, error) {
		if un, ok := value.(unstructured.Unstructured); ok {
			for key := range un {
				if isPreprocessorKey(key) {
					return p.runPreprocess(cfPath, value.(unstructured.Unstructured))
				}
			}
		}

		return value, nil
	})

	if err != nil {
		return nil, err
	}

	return visitedBody, nil
}

func (p *preprocessor) runPreprocess(cfPath string, un unstructured.Unstructured) (interface{}, error) {
	var err error
	var keys = unstructured.Keys(un)

	if util.ArrayContains(keys, "$file") {
		return p.runIncludeFile(cfPath, un)
	}

	if util.ArrayContains(keys, "$base64File") {
		return p.runIncludeBase64File(un)
	}

	if util.ArrayContains(keys, "$folder") {
		return p.runIncludeFolder(un)
	}

	if util.ArrayContains(keys, "$extend") {
		un, err = p.runPreprocessExtend(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(cfPath, un)
	}

	if util.ArrayContains(keys, "$include") {
		result, err := p.runPreprocessInclude(cfPath, un)

		if err != nil {
			return nil, err
		}

		return result, nil
	}

	if util.ArrayContains(keys, "$properties") {
		un, err = p.runPreprocessProperties(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(cfPath, un)
	}

	if util.ArrayContains(keys, "$override") {
		un, err = p.runPreprocessOverride(un)

		if err != nil {
			return nil, err
		}

		return p.runPreprocess(cfPath, un)
	}

	if util.ArrayContains(keys, "$syntax") {
		err = p.checkSyntax(cfPath, un)

		if err != nil {
			return nil, err
		}

		return un, nil
	}

	return un, nil
}

func (p *preprocessor) runPreprocessInclude(cfPath string, un unstructured.Unstructured) (interface{}, error) {
	filePath := un["$include"].(string)

	if strings.HasPrefix(filePath, "./") {
		filePath = p.resolveAbsolutePath(cfPath, filePath)
	}

	var result []unstructured.Unstructured
	var handler = func(data unstructured.Unstructured) error {
		processed, err := p.preprocess(filePath, data)

		if err != nil {
			return err
		}
		result = append(result, processed.(unstructured.Unstructured))
		return nil
	}

	var in io.Reader
	if strings.HasPrefix(filePath, "https://") {
		resp, err := http.Get(filePath)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		in = resp.Body
	} else {
		var err error
		in, err = os.Open(filePath)

		if err != nil {
			return nil, err
		}
	}

	if strings.HasSuffix(filePath, ".json") {
		if err := jsonformat.Parse(in, context.TODO(), handler); err != nil {
			return nil, err
		}
	} else {
		if err := yamlformat.Parse(in, context.TODO(), handler); err != nil {
			return nil, err
		}
	}

	if len(result) == 1 {
		return result[0], nil
	} else {
		return result, nil
	}
}

func (p *preprocessor) runIncludeFile(cfPath string, un unstructured.Unstructured) (string, error) {
	filePath := un["$file"].(string)

	if strings.HasPrefix(filePath, "./") {
		filePath = p.resolveAbsolutePath(cfPath, filePath)
	}

	var in io.Reader
	if strings.HasPrefix(filePath, "https://") {
		resp, err := http.Get(filePath)

		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		in = resp.Body
	} else {
		var err error
		in, err = os.Open(filePath)

		if err != nil {
			return "", err
		}
	}

	dat, err := io.ReadAll(in)

	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func (p *preprocessor) runIncludeBase64File(un unstructured.Unstructured) (string, error) {
	filePath := un["$base64File"]

	dat, err := os.ReadFile(filePath.(string))

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(dat), nil
}

func (p *preprocessor) runIncludeFolder(un unstructured.Unstructured) (string, error) {
	filePath := un["$folder"].(string)
	format := un["$format"].(string)

	if format == "" || format == "tar" {
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

	if namespace != "" {
		namespace = "ui" //default
	}

	var filters = make(map[string]string)

	if extend["$match"] != nil {
		match := extend["$match"].(unstructured.Unstructured)
		for key, value := range match {
			filters[key] = value.(string)
		}
	}

	records, _, err := p.dhClient.ListRecords(context.TODO(), service.RecordListParams{
		Namespace: namespace,
		Resource:  resource,
		Filters:   filters,
	})

	if err != nil {
		reportError(un, err)
		return nil, err
	}

	ref := extend["$ref"].(string)

	if len(records) == 0 {
		return un, fmt.Errorf("no record found for %s/%s/%v", namespace, resource, extend["$match"])
	}

	recordUn, err := p.writer.WriteRecord(namespace, resource, records[0])

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

func (p *preprocessor) checkSyntax(cfPath string, un unstructured.Unstructured) error {
	if un["$syntax"] != nil {
		syntax, err := p.runPreprocess(cfPath, un["$syntax"].(unstructured.Unstructured))

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
				return nil, fmt.Errorf("invalid property type %s", typeStr)
			}
			err := unstructured.FromProtoMessage(propertyUn, property)

			if err != nil {
				return nil, err
			}
		} else if valueUn, ok := value.(unstructured.Unstructured); ok {
			propertyUn = valueUn
		} else {
			return nil, fmt.Errorf("invalid property type %s", value)
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

func (p *preprocessor) resolveAbsolutePath(baseURL, relativePath string) string {
	// Remove './' from the beginning of the relative path, if present
	relativePath = strings.TrimPrefix(relativePath, "./")

	// Modify the path of the URL
	parts := strings.Split(baseURL, "/")

	pp := strings.Join(parts[:len(parts)-1], "/")

	return pp + "/" + relativePath
}

var preprocessKeywords = []string{
	"$extend", "$override", "$select", "$syntax", "$ref", "$merge", "$set", "$clear", "$append", "$include", "$expression", "$properties", "$file", "$base64File", "$folder",
}

func isPreprocessorKey(key string) bool {
	for _, keyword := range preprocessKeywords {
		if key == keyword {
			return true
		}
	}

	return false
}
