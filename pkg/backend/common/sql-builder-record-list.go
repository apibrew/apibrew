package common

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/backend/helper"
	"github.com/apibrew/apibrew/pkg/backend/sqlbuilder"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/types"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type colDetails struct {
	colName      string
	path         string
	def          string
	alias        string
	propertyType model.ResourceProperty_Type
	required     bool
	resource     *model.Resource
}

type joinDetails struct {
	targetTable      string
	targetTableAlias string
	targetColumn     string
	sourcePath       string
}

type recordLister struct {
	tableName  string
	tableAlias string
	ctx        context.Context
	runner     helper.QueryRunner
	colList    []colDetails
	joins      []joinDetails

	resource          *model.Resource
	query             *model.BooleanExpression
	Limit             uint32
	Offset            uint64
	ResolveReferences []string
	logger            *log.Entry
	builder           *sqlbuilder.SelectBuilder
	resultChan        chan<- *model.Record
	packRecords       bool
	backend           *sqlBackend
	aggregation       *model.Aggregation
	sorting           *model.Sorting
	propertyNameMap   map[string]*model.ResourceProperty
}

func (r *recordLister) Prepare() error {
	r.logger = log.WithFields(logging.CtxFields(r.ctx))
	r.tableName = r.backend.getFullTableName(r.resource.SourceConfig)
	r.tableAlias = "t"

	r.builder = sqlbuilder.Select()
	r.builder.SetFlavor(r.backend.options.GetFlavor())

	r.propertyNameMap = util.GetNamedMap(r.resource.Properties)

	r.builder.From(r.tableName + " as " + r.tableAlias)

	r.expandProps("t", r.resource)

	if r.query != nil {
		var where string
		where, err := r.applyCondition(r.resource, r.query)
		if err != nil {
			return err
		}

		if where != "" {
			r.builder.Where(where)
		}
	}

	if r.Limit == 0 {
		r.Limit = 100
	}

	if r.Limit > 10000 && r.resultChan == nil {
		r.Limit = 10000
	}

	r.builder.Limit(int(r.Limit))
	r.builder.Offset(int(r.Offset))

	if r.aggregation != nil {
		if err := r.prepareAggregation(); err != nil {
			return err
		}
	}

	return nil
}

func (r *recordLister) prepareAggregation() error {
	r.colList = nil

	if len(r.aggregation.Grouping) > 0 {
		var groupByCols []string
		var path = "t"

		for _, item := range r.aggregation.Grouping {
			var prop = r.propertyNameMap[item.Property]
			if prop == nil {
				return errors.RecordValidationError.WithDetails("Grouping property not exists: " + item.Property)
			}

			groupByCols = append(groupByCols, path+"."+r.quote(item.Property))

			r.colList = append(r.colList, colDetails{
				resource:     r.resource,
				colName:      item.Property,
				path:         "t_" + item.Property,
				def:          "t." + r.quote(item.Property),
				alias:        r.quote("t_" + item.Property),
				required:     true,
				propertyType: prop.Type,
			})
		}

		r.builder.GroupBy(groupByCols...)
	}

	for _, item := range r.aggregation.Items {
		var prop = r.propertyNameMap[item.Property]
		if prop == nil {
			return errors.RecordValidationError.WithDetails("Grouping property not exists: " + item.Property)
		}

		var fnName string
		var propType model.ResourceProperty_Type

		switch item.Algorithm {
		case model.AggregationItem_COUNT:
			fnName = "count"
			propType = model.ResourceProperty_INT32
		case model.AggregationItem_MIN:
			fnName = "min"
			propType = prop.Type
		case model.AggregationItem_MAX:
			fnName = "max"
			propType = prop.Type
		case model.AggregationItem_SUM:
			fnName = "sum"
			if prop.Type == model.ResourceProperty_INT32 || prop.Type == model.ResourceProperty_INT64 {
				propType = model.ResourceProperty_INT64
			} else {
				propType = model.ResourceProperty_FLOAT64
			}
		case model.AggregationItem_AVG:
			fnName = "avg"
			propType = model.ResourceProperty_FLOAT64
		}

		r.colList = append(r.colList, colDetails{
			resource:     r.resource,
			colName:      item.Name,
			path:         "t_" + item.Name,
			def:          fmt.Sprintf("%s(%s)", fnName, "t."+r.quote(item.Property)),
			alias:        r.quote("t_" + item.Property),
			required:     true,
			propertyType: propType,
		})
	}
	return nil
}

func (r *recordLister) prepareSorting() error {
	var orderBy []string

	for _, item := range r.sorting.Items {
		var prop = r.propertyNameMap[item.Property]
		if prop == nil {
			return errors.RecordValidationError.WithDetails("Sorting property not exists: " + item.Property)
		}

		orderBy = append(orderBy, fmt.Sprintf("%s %s", r.quote("t_"+item.Property), item.Direction))
	}

	r.builder.OrderBy(orderBy...)

	return nil
}

func (r *recordLister) Exec() (result []*model.Record, total uint32, err error) {
	if err := r.Prepare(); err != nil {
		return nil, 0, err
	}

	total, err = r.ExecCount()
	if err != nil || total == 0 || r.Limit == 0 {
		return
	}

	selectBuilder := r.builder

	if r.sorting != nil {
		if err = r.prepareSorting(); err != nil {
			return nil, 0, err
		}
	}

	selectBuilder.Select(r.prepareCols()...)

	if r.sorting == nil && r.aggregation == nil {
		selectBuilder.OrderBy("t.id ASC")
	}

	for _, jd := range r.joins {
		selectBuilder.JoinWithOption(sqlbuilder.LeftJoin, fmt.Sprintf("%s as %s", r.quote(jd.targetTable), r.quote(jd.targetTableAlias)), fmt.Sprintf("%s.%s = %s", r.quote(jd.targetTableAlias), r.quote(jd.targetColumn), jd.sourcePath))
	}

	sqlQuery, args := selectBuilder.Build()

	rows, sqlErr := r.runner.Query(sqlQuery, args...)
	err = r.backend.handleDbError(r.ctx, sqlErr)

	if err != nil {
		return
	}

	defer func() {
		err2 := rows.Close()

		if err2 != nil {
			r.logger.Print(err2)
		}
	}()

	for rows.Next() {
		select {
		case <-r.ctx.Done():
			break
		default:
		}

		record := new(model.Record)

		err = r.scanRecord(record, rows)
		if err != nil {
			return
		}

		if r.resultChan != nil {
			r.resultChan <- record
		} else {
			result = append(result, record)
		}

		if r.packRecords {
			for _, prop := range r.resource.Properties {
				if helper.IsPropertyOmitted(prop) {
					continue
				}
				record.PropertiesPacked = append(record.PropertiesPacked, record.Properties[prop.Name])
			}
			record.Properties = nil
		}
	}

	return
}

func (r *recordLister) ExecCount() (total uint32, err error) {
	countBuilder := r.builder.Copy()
	countBuilder.Select("count(*)")
	countBuilder.Limit(1)
	countBuilder.Offset(0)

	countQuery, args := countBuilder.Build()
	r.logger.Tracef("countQuery: %s", countQuery)

	countRow := r.runner.QueryRowContext(r.ctx, countQuery, args...)
	err = r.backend.handleDbError(r.ctx, countRow.Scan(&total))

	return
}

func (r *recordLister) expandProps(path string, resource *model.Resource) {
	isInner := path != "t"

	for _, prop := range resource.Properties {
		if helper.IsPropertyOmitted(prop) {
			continue
		}

		r.colList = append(r.colList, colDetails{
			resource:     resource,
			colName:      prop.Name,
			path:         path + "_" + prop.Name,
			def:          r.quote(path) + "." + r.quote(prop.Name),
			alias:        r.quote(path + "_" + prop.Name),
			required:     prop.Required && !isInner,
			propertyType: prop.Type,
		})

		if annotations.IsEnabled(annotations.FromCtx(r.ctx), annotations.UseJoinTable) && prop.Type == model.ResourceProperty_REFERENCE {
			// check resource
			found := false
			for _, rr := range r.ResolveReferences {
				if rr == "*" || rr == prop.Name || strings.HasPrefix(rr, prop.Name+"/") {
					found = true
				}
			}
			if found {
				// locating referenced resource
				referenceNamespace := prop.Reference.Namespace
				if referenceNamespace == "" {
					referenceNamespace = resource.Namespace
				}
				referencedResource := r.backend.schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+prop.Reference.Resource]
				newPath := path + "__" + prop.Name

				// add to joins
				r.joins = append(r.joins, joinDetails{
					targetTable:      referencedResource.SourceConfig.Entity,
					targetTableAlias: newPath,
					targetColumn:     "id",
					sourcePath:       r.quote(path) + "." + r.quote(prop.Name),
				})

				r.expandProps(newPath, referencedResource)
			}
		}
	}
}

func (r *recordLister) quote(path string) string {
	return r.backend.options.Quote(path)
}

func (r *recordLister) scanRecord(record *model.Record, rows *sql.Rows) error {
	var rowScanFields []any
	var propertyPointers = make(map[string]interface{})

	for _, cd := range r.colList {
		propertyType := r.backend.options.TypeModifier(cd.propertyType)

		val := propertyType.Pointer(cd.required)

		if cd.propertyType == model.ResourceProperty_REFERENCE {
			if cd.required {
				val = new(string)
			} else {
				val = new(*string)
			}
		}

		propertyPointers[cd.path] = val

		rowScanFields = append(rowScanFields, val)
	}

	err := rows.Scan(rowScanFields...)

	if err == sql.ErrNoRows {
		return errors.RecordNotFoundError.WithDetails(fmt.Sprintf("namespace: %s; resource: %s", r.resource.Namespace, r.resource.Name))
	}

	if err != nil {
		return r.backend.handleDbError(r.ctx, err)
	}

	var serviceErr error

	record.Properties, serviceErr = r.mapRecordProperties(util.GetRecordId(record), r.resource, "t_", propertyPointers)
	if serviceErr != nil {
		return serviceErr
	}

	return nil
}

func (r *recordLister) mapRecordProperties(recordId string, resource *model.Resource, pathPrefix string, propertyPointers map[string]interface{}) (map[string]*structpb.Value, error) {
	properties := make(map[string]*structpb.Value)

	for _, cd := range r.colList {
		propertyType := r.backend.options.TypeModifier(cd.propertyType)
		propPointer := propertyPointers[cd.path]

		propV := types.Dereference(propPointer)
		if propV == nil {
			continue
		}

		val, err := DbDecode(cd.propertyType, propV)

		if err != nil {
			return nil, err
		}

		if r.aggregation != nil {
			for _, item := range r.aggregation.Items {
				if cd.path == "t_"+item.Name {
					v, err2 := propertyType.Pack(val)

					if err2 != nil {
						return nil, errors.InternalError.WithDetails(err2.Error())
					}

					properties[item.Name] = v
				}
			}
		}

		for _, prop := range resource.Properties {
			if helper.IsPropertyOmitted(prop) {
				continue
			}
			if pathPrefix+prop.Name == cd.path {

				if prop.Type == model.ResourceProperty_REFERENCE {
					resolveReference := false
					for _, rr := range r.ResolveReferences {
						if rr == "*" || rr == prop.Name || strings.HasPrefix(rr, prop.Name+"/") {
							resolveReference = true
						}
					}

					referenceNamespace := prop.Reference.Namespace
					if referenceNamespace == "" {
						referenceNamespace = resource.Namespace
					}
					referencedResource := r.backend.schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+prop.Reference.Resource]

					if referencedResource != nil && resolveReference {
						nv, err := r.mapRecordProperties(recordId, referencedResource, pathPrefix+"_"+prop.Name+"_", propertyPointers)
						if err != nil {
							return nil, err
						}

						properties[prop.Name] = structpb.NewStructValue(&structpb.Struct{Fields: nv})
						v1 := properties[prop.Name].GetStructValue().Fields["id"].GetStringValue()
						v2 := val.(map[string]interface{})["id"]
						if v1 != v2 {
							log.Print(properties[prop.Name], val)
						}
					} else {
						st, err := structpb.NewStruct(val.(map[string]interface{}))

						if err != nil {
							return nil, errors.InternalError.WithDetails(err.Error())
						}

						properties[prop.Name] = structpb.NewStructValue(st)
					}
				} else {
					v, err2 := propertyType.Pack(val)

					if err2 != nil {
						return nil, errors.InternalError.WithDetails(err2.Error())
					}

					properties[prop.Name] = v
				}
				break
			}
		}
	}

	return properties, nil
}

func (r *recordLister) applyCondition(resource *model.Resource, query *model.BooleanExpression) (string, error) {
	if and, ok := query.Expression.(*model.BooleanExpression_And); ok {
		if len(and.And.Expressions) == 0 {
			return "", nil
		}

		expressions, err := util.ArrayMapWithError(and.And.Expressions, func(t *model.BooleanExpression) (string, error) {
			return r.applyCondition(resource, t)
		})
		if err != nil {
			return "", err
		}
		return r.builder.And(expressions...), nil
	}

	if and, ok := query.Expression.(*model.BooleanExpression_Or); ok {
		if len(and.Or.Expressions) == 0 {
			return "", nil
		}

		expressions, err := util.ArrayMapWithError(and.Or.Expressions, func(t *model.BooleanExpression) (string, error) {
			return r.applyCondition(resource, t)
		})
		if err != nil {
			return "", err
		}
		return r.builder.Or(expressions...), nil
	}

	if and, ok := query.Expression.(*model.BooleanExpression_Not); ok {
		exp, err := r.applyCondition(resource, and.Not)

		if err != nil {
			return "", err
		}

		return fmt.Sprintf("not (%s)", exp), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_Equal); ok {
		left, right, err := r.applyExpressionPair(resource, equ.Equal)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s = %s", left, right), nil
	}

	if isn, ok := query.Expression.(*model.BooleanExpression_IsNull); ok {
		left, err := r.applyExpression(resource, isn.IsNull)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s is null", left), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_GreaterThan); ok {
		left, right, err := r.applyExpressionPair(resource, equ.GreaterThan)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s > %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_GreaterThanOrEqual); ok {
		left, right, err := r.applyExpressionPair(resource, equ.GreaterThanOrEqual)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s >= %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_LessThan); ok {
		left, right, err := r.applyExpressionPair(resource, equ.LessThan)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s < %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_LessThanOrEqual); ok {
		left, right, err := r.applyExpressionPair(resource, equ.LessThanOrEqual)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s <= %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_In); ok {
		left, right, err := r.applyExpressionPair(resource, equ.In)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s in (%s)", left, right), nil
	}

	if _, ok := query.Expression.(*model.BooleanExpression_RegexMatch); ok {
		panic("not implemented")
	}

	if query.Expression == nil {
		return "", errors.RecordValidationError.WithDetails("Empty expression is sent")
	}

	return "", errors.LogicalError.WithDetails("Unknown boolean expression type: " + query.String())
}

func (r *recordLister) applyExpressionPair(resource *model.Resource, pair *model.PairExpression) (string, string, error) {
	var left string
	var right string
	var property *model.ResourceProperty

	if propEx, ok := pair.Left.Expression.(*model.Expression_Property); ok {
		property = util.LocatePropertyByName(resource, propEx.Property)

		if property == nil {
			return "", "", errors.PropertyNotFoundError.WithDetails(propEx.Property)
		}

		left = fmt.Sprintf("t." + r.quote(property.Name))
	} else {
		return "", "", errors.LogicalError.WithDetails("Only property expression is allowed on the left part: " + pair.Left.String())
	}

	if propEx, ok := pair.Right.Expression.(*model.Expression_Value); ok {
		if property.Type == model.ResourceProperty_REFERENCE {
			if propEx.Value.GetStructValue() != nil {
				properties := propEx.Value.GetStructValue().Fields
				if properties["id"] != nil {
					right = r.applyValue(properties["id"])
				} else {
					referenceNamespace := property.Reference.Namespace

					if referenceNamespace == "" {
						referenceNamespace = resource.Namespace
					}

					referencedResource := r.backend.schema.ResourceByNamespaceSlashName[referenceNamespace+"/"+property.Reference.Resource]

					if referencedResource == nil {
						return "", "", errors.LogicalError.WithDetails("Referenced resource not found: " + referenceNamespace + "/" + property.Reference.Resource)
					}

					innerSql, err := r.backend.resolveReference(properties, r.builder.Var, referencedResource)

					if err != nil {
						return "", "", err
					}

					right = innerSql
				}
			} else {
				right = r.applyValue(propEx.Value)
			}
		} else {
			right = r.applyValue(propEx.Value)
		}
	} else {
		return "", "", errors.LogicalError.WithDetails("Only value expression is allowed on the right part: " + pair.Right.String())
	}

	return left, right, nil
}

func (r *recordLister) applyValue(value *structpb.Value) string {
	if value.GetListValue() != nil {
		list := value.GetListValue()
		var c []string
		for _, val := range list.Values {
			c = append(c, r.builder.Var(val.AsInterface()))
		}
		return strings.Join(c, ",")
	} else {
		return r.builder.Var(value.AsInterface())
	}
}

func (r *recordLister) applyExpression(resource *model.Resource, query *model.Expression) (string, error) {
	if query.Expression == nil {
		return "", errors.PropertyNotFoundError.WithDetails("expression is empty")
	}

	if propEx, ok := query.Expression.(*model.Expression_Property); ok {
		property := util.LocatePropertyByName(resource, propEx.Property)

		if property == nil {
			return "", errors.PropertyNotFoundError.WithDetails(propEx.Property)
		}

		return fmt.Sprintf("t." + r.quote(property.Name)), nil
	} else {
		return "", errors.LogicalError.WithDetails("Only property expression is allowed: " + query.String())
	}
}

func (r *recordLister) prepareCols() []string {
	var cols []string

	for _, cd := range r.colList {
		cols = append(cols, fmt.Sprintf("%s as %s", cd.def, cd.alias))
	}

	return cols
}

func (p *sqlBackend) recordList(ctx context.Context, runner helper.QueryRunner, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) (result []*model.Record, total uint32, err error) {
	return (&recordLister{
		ctx:               ctx,
		runner:            runner,
		resource:          resource,
		query:             params.Query,
		aggregation:       params.Aggregation,
		sorting:           params.Sorting,
		Limit:             params.Limit,
		Offset:            params.Offset,
		ResolveReferences: params.ResolveReferences,
		resultChan:        resultChan,
		backend:           p,
	}).Exec()
}

func DbDecode(propertyType model.ResourceProperty_Type, val interface{}) (interface{}, error) {
	if propertyType == model.ResourceProperty_STRUCT || propertyType == model.ResourceProperty_OBJECT || propertyType == model.ResourceProperty_MAP || propertyType == model.ResourceProperty_LIST {
		var data = new(interface{})
		err2 := json.Unmarshal([]byte(val.(string)), data)

		if err2 != nil {
			return nil, errors.InternalError.WithDetails(err2.Error())
		}

		val = *data
	} else if propertyType == model.ResourceProperty_REFERENCE {
		return map[string]interface{}{
			"id": val,
		}, nil
	}

	return val, nil
}
