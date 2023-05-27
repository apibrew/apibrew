package unstructured

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/client"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/validate"
	"github.com/apibrew/apibrew/pkg/stub"
	"github.com/apibrew/apibrew/pkg/util"
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
	var keys = un.Keys()

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

func (p *preprocessor) runPreprocessInclude(un Unstructured) (Unstructured, error) {
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

	un.MergeOut(recordUn, true)

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

	un.DeleteKey("$extend")

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
	merge, _ := override["$merge"].(Unstructured)
	set, _ := override["$set"].(Unstructured)

	reUn, err := ParseRef(un, selectParam)

	if err != nil {
		return nil, err
	}

	if merge != nil {
		reUn.MergeInto(merge, true)
		un.DeleteKey("$merge")
	}

	if set != nil {
		reUn.MergeInto(set, false)
		un.DeleteKey("$set")
	}

	un.DeleteKey("$override")

	return un, nil
}

func (p *preprocessor) checkSyntax(un Unstructured) error {
	if un["$syntax"] != nil {
		syntax := un["$syntax"].(Unstructured)

		un.DeleteKey("$syntax")

		var subType = &model.Resource{}

		err := syntax.ToProtoMessage(subType)

		if err != nil {
			return err
		}

		record, err := un.ToRecord()

		if err != nil {
			return err
		}

		err = validate.ValidateRecords(subType, []*model.Record{record}, false)

		// recover syntax as it needs to be persisted
		un["$syntax"] = syntax

		if err != nil {
			return err
		}
	}

	return nil
}

var preprocessKeywords = []string{
	"$extend", "$override", "$select", "$syntax", "$ref", "$merge", "$set", "$clear", "$append", "$include", "$expression",
}

func isPreprocessorKey(key string) bool {
	for _, keyword := range preprocessKeywords {
		if key == keyword {
			return true
		}
	}

	return false
}
