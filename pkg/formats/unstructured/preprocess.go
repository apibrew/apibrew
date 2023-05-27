package unstructured

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/stub"
	log "github.com/sirupsen/logrus"
)

type preprocessor struct {
	dhClient client.DhClient
	writer   *Writer
}

func (p *preprocessor) preprocess(body Unstructured) (Unstructured, error) {
	visitedBody, err := WalkUnstructured(body, func(value interface{}) (interface{}, error) {
		if un, ok := value.(Unstructured); ok {
			for key := range un {
				if isPreprocessorKey(key) {
					return p.runPreprocess(value.(Unstructured))
				}
			}
		}

		return value, nil
	})

	if err != nil {
		return nil, err
	}

	return visitedBody.(Unstructured), nil
}

func (p *preprocessor) runPreprocess(un Unstructured) (interface{}, error) {
	var err error
	for key := range un {
		if key == "$extend" {
			un, err = p.runPreprocessExtend(un)

			if err != nil {
				return nil, err
			}
		}
		if key == "$override" {
			un, err = p.runPreprocessOverride(un)

			if err != nil {
				return nil, err
			}
		}
	}

	return un, nil
}

func (p *preprocessor) runPreprocessExtend(un Unstructured) (Unstructured, error) {
	extend := un["$extend"].(Unstructured)

	namespace := extend["$namespace"].(string)
	resource := extend["$resource"].(string)
	match := extend["$match"].(Unstructured)

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

	un.MergeInto(recordUn)

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

	delete(un, "$extend")

	return un, nil
}

func reportError(un Unstructured, err error) {
	data, _ := json.MarshalIndent(un, " ", "  ")
	fmt.Println(string(data))
	log.Error(err)
}

func (p *preprocessor) runPreprocessOverride(un Unstructured) (Unstructured, error) {
	override := un["$override"].(Unstructured)
	selectParam := override["$select"].(string)
	set := override["$set"].(Unstructured)

	reUn, err := ParseRef(un, selectParam)

	if err != nil {
		return nil, err
	}

	if set != nil {
		reUn.MergeInto(set)
		delete(un, "$set")
	}

	delete(un, "$override")

	return un, nil
}

var preprocessKeywords = []string{
	"$extend", "$override", "$select", "$syntax", "$ref", "$set", "$clear", "$append",
}

func isPreprocessorKey(key string) bool {
	for _, keyword := range preprocessKeywords {
		if key == keyword {
			return true
		}
	}

	return false
}
