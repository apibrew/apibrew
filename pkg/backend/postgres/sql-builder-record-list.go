package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/backend/sqlbuilder"
	"github.com/tislib/data-handler/pkg/errors"
	"github.com/tislib/data-handler/pkg/logging"
	"github.com/tislib/data-handler/pkg/model"
	"github.com/tislib/data-handler/pkg/service/annotations"
	"github.com/tislib/data-handler/pkg/types"
	"github.com/tislib/data-handler/pkg/util"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type colDetails struct {
	colName      string
	path         string
	def          string
	alias        string
	property     *model.ResourceProperty
	propertyType model.ResourcePropertyType
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
	runner     QueryRunner
	colList    []colDetails
	joins      []joinDetails

	resource          *model.Resource
	query             *model.BooleanExpression
	Limit             uint32
	Offset            uint64
	UseHistory        bool
	ResolveReferences []string
	Schema            abs.Schema
	logger            *log.Entry
	builder           *sqlbuilder.SelectBuilder
	resultChan        chan<- *model.Record
	packRecords       bool
}

func (r *recordLister) Prepare() errors.ServiceError {
	r.logger = log.WithFields(logging.CtxFields(r.ctx))
	r.tableName = getFullTableName(r.resource.SourceConfig, r.UseHistory)
	r.tableAlias = "t"

	r.builder = sqlbuilder.Select()
	r.builder.SetFlavor(sqlbuilder.PostgreSQL)

	r.builder.From(r.tableName + " as " + r.tableAlias)

	r.expandProps("t", r.resource)

	if r.query != nil {
		var where string
		where, err := r.applyCondition(r.resource, r.query)
		if err != nil {
			return err
		}
		r.builder.Where(where)
	}

	if r.Limit == 0 {
		r.Limit = 100
	}

	if r.Limit > 10000 && r.resultChan == nil {
		r.Limit = 10000
	}

	r.builder.Limit(int(r.Limit))
	r.builder.Offset(int(r.Offset))

	return nil
}

func (r *recordLister) Exec() (result []*model.Record, total uint32, err errors.ServiceError) {
	if err := r.Prepare(); err != nil {
		return nil, 0, err
	}

	total, err = r.ExecCount()
	if err != nil || total == 0 {
		return
	}

	selectBuilder := r.builder
	selectBuilder.Select(r.prepareCols()...)

	for _, jd := range r.joins {
		selectBuilder.JoinWithOption(sqlbuilder.LeftJoin, fmt.Sprintf("%s as %s", jd.targetTable, jd.targetTableAlias), fmt.Sprintf("%s.%s = %s", jd.targetTableAlias, jd.targetColumn, jd.sourcePath))
	}

	sqlQuery, args := selectBuilder.Build()

	rows, sqlErr := r.runner.Query(sqlQuery, args...)
	err = handleDbError(r.ctx, sqlErr)

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
		record.DataType = model.DataType_USER

		err = r.scanRecord(record, rows)
		if err != nil {
			return
		}

		if r.resultChan != nil {
			r.resultChan <- record
		} else {
			result = append(result, record)
		}

		if !r.packRecords && annotations.IsEnabled(r.resource, annotations.DoPrimaryKeyLookup) {
			err := computeRecordFromProperties(r.ctx, r.resource, record)

			if err != nil {
				return nil, 0, err
			}
		}

		if r.packRecords {
			for _, prop := range r.resource.Properties {
				record.PropertiesPacked = append(record.PropertiesPacked, record.Properties[prop.Name])
			}
			record.Properties = nil
		}
	}

	return
}

func (r *recordLister) ExecCount() (total uint32, err errors.ServiceError) {
	countBuilder := r.builder.Copy()
	countBuilder.Select("count(*)")

	countQuery, args := countBuilder.Build()
	r.logger.Tracef("countQuery: %s", countQuery)

	countRow := r.runner.QueryRowContext(r.ctx, countQuery, args...)
	err = handleDbError(r.ctx, countRow.Scan(&total))

	return
}

func (r *recordLister) expandProps(path string, resource *model.Resource) {
	isInner := path != "t"
	gCol := func(name string, typ model.ResourcePropertyType) colDetails {
		return colDetails{
			resource:     resource,
			colName:      name,
			path:         path + "_" + name,
			def:          path + "." + name,
			alias:        path + "_" + name,
			propertyType: typ,
			required:     false,
		}
	}

	if !annotations.IsEnabled(r.resource, annotations.DoPrimaryKeyLookup) {
		r.colList = append(r.colList, gCol("id", model.ResourcePropertyType_TYPE_UUID))
	}

	if !annotations.IsEnabled(r.resource, annotations.DisableAudit) {
		r.colList = append(r.colList, gCol("created_on", model.ResourcePropertyType_TYPE_TIMESTAMP))
		r.colList = append(r.colList, gCol("updated_on", model.ResourcePropertyType_TYPE_TIMESTAMP))
		r.colList = append(r.colList, gCol("created_by", model.ResourcePropertyType_TYPE_STRING))
		r.colList = append(r.colList, gCol("updated_by", model.ResourcePropertyType_TYPE_STRING))
	}

	if !annotations.IsEnabled(r.resource, annotations.DisableVersion) {
		r.colList = append(r.colList, gCol("version", model.ResourcePropertyType_TYPE_INT32))
	}

	for _, prop := range resource.Properties {
		r.colList = append(r.colList, colDetails{
			resource:     resource,
			colName:      prop.Mapping,
			path:         path + "_" + prop.Mapping,
			def:          path + "." + prop.Mapping,
			alias:        path + "_" + prop.Mapping,
			property:     prop,
			required:     prop.Required && !isInner,
			propertyType: prop.Type,
		})

		if prop.Type == model.ResourcePropertyType_TYPE_REFERENCE {
			// check resource
			found := false
			for _, rr := range r.ResolveReferences {
				if rr == "*" || rr == prop.Name || strings.HasPrefix(rr, prop.Name+"/") {
					found = true
				}
			}
			if found {
				// locating referenced resource
				referencedResource := r.Schema.ResourceByNamespaceSlashName[r.resource.Namespace+"/"+prop.Reference.ReferencedResource]
				newPath := path + "__" + prop.Mapping
				r.expandProps(newPath, referencedResource)

				// add to joins
				r.joins = append(r.joins, joinDetails{
					targetTable:      referencedResource.SourceConfig.Entity,
					targetTableAlias: newPath,
					targetColumn:     "id",
					sourcePath:       path + "." + prop.Mapping,
				})
			}
		}
	}
}

func (r *recordLister) scanRecord(record *model.Record, rows *sql.Rows) errors.ServiceError {
	var rowScanFields []any
	var propertyPointers = make(map[string]interface{})

	for _, cd := range r.colList {
		propertyType := types.ByResourcePropertyType(cd.propertyType)

		val := propertyType.Pointer(cd.required)

		if cd.propertyType == model.ResourcePropertyType_TYPE_REFERENCE {
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
		return handleDbError(r.ctx, err)
	}

	if !annotations.IsEnabled(r.resource, annotations.DisableAudit) {
		record.AuditData = &model.AuditData{}
	}

	if propertyPointers["t_id"] != nil && !annotations.IsEnabled(r.resource, annotations.DoPrimaryKeyLookup) {
		record.Id = (**(propertyPointers["t_id"].(**uuid.UUID))).String()
	}

	if propertyPointers["t_created_on"] != nil && *propertyPointers["t_created_on"].(**time.Time) != nil {
		record.AuditData.CreatedOn = timestamppb.New(**propertyPointers["t_created_on"].(**time.Time))
	}

	if propertyPointers["t_updated_on"] != nil && *propertyPointers["t_updated_on"].(**time.Time) != nil {
		record.AuditData.UpdatedOn = timestamppb.New(**propertyPointers["t_updated_on"].(**time.Time))
	}

	if propertyPointers["t_created_by"] != nil && *propertyPointers["t_created_by"].(**string) != nil {
		record.AuditData.CreatedBy = **propertyPointers["t_created_by"].(**string)
	}

	if propertyPointers["t_updated_by"] != nil && *propertyPointers["t_updated_by"].(**string) != nil {
		record.AuditData.CreatedBy = **propertyPointers["t_updated_by"].(**string)
	}

	if propertyPointers["t_version"] != nil && *propertyPointers["t_version"].(**int32) != nil {
		record.Version = uint32(**propertyPointers["t_version"].(**int32))
	}

	var serviceErr errors.ServiceError

	record.Properties, serviceErr = r.mapRecordProperties(record.Id, r.resource, "t_", propertyPointers)
	if serviceErr != nil {
		return serviceErr
	}

	if !annotations.IsEnabled(r.resource, annotations.DoPrimaryKeyLookup) {
		delete(record.Properties, "id")
	}

	return nil
}

func (r *recordLister) mapRecordProperties(recordId string, resource *model.Resource, pathPrefix string, propertyPointers map[string]interface{}) (map[string]*structpb.Value, errors.ServiceError) {
	properties := make(map[string]*structpb.Value)

	if propertyPointers[pathPrefix+"id"] != nil && !annotations.IsEnabled(r.resource, annotations.DoPrimaryKeyLookup) {
		id := (**(propertyPointers[pathPrefix+"id"].(**uuid.UUID))).String()
		properties["id"] = structpb.NewStringValue(id)
	}

	for _, cd := range r.colList {
		if cd.property == nil {
			continue
		}
		propertyType := types.ByResourcePropertyType(cd.propertyType)
		propPointer := propertyPointers[cd.path]

		propV := types.Dereference(propPointer)
		if propV == nil {
			continue
		}

		val, err := DbDecode(cd.property, propV)

		if err != nil {
			return nil, err
		}

		for _, prop := range resource.Properties {
			if pathPrefix+prop.Mapping == cd.path {

				if prop.Type == model.ResourcePropertyType_TYPE_REFERENCE {
					resolveReference := false
					for _, rr := range r.ResolveReferences {
						if rr == "*" || rr == prop.Name || strings.HasPrefix(rr, prop.Name+"/") {
							resolveReference = true
						}
					}

					referencedResource := r.Schema.ResourceByNamespaceSlashName[r.resource.Namespace+"/"+prop.Reference.ReferencedResource]

					if referencedResource != nil && resolveReference {
						nv, err := r.mapRecordProperties(recordId, referencedResource, pathPrefix+"_"+prop.Mapping+"_", propertyPointers)
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

		//v1 := fmt.Sprintf("%v", val)
		//v2 := fmt.Sprintf("%v", properties[cd.property.Name].AsInterface())
		//if v1 != v2 {
		//	fmt.Print("\n", val, "  <->  ", properties[cd.property.Name].AsInterface(), "\n")
		//	log.Print("Diffferent vals")
		//}

		r.logger.Tracef("%s[%s]=%s [%s](%s)", recordId, cd.path, val, cd.property.Name, properties[cd.property.Name])
	}

	return properties, nil
}

func (r *recordLister) applyCondition(resource *model.Resource, query *model.BooleanExpression) (string, errors.ServiceError) {
	if and, ok := query.Expression.(*model.BooleanExpression_And); ok {
		if len(and.And.Expressions) == 0 {
			return "", nil
		}

		expressions, err := util.ArrayMapWithError(and.And.Expressions, func(t *model.BooleanExpression) (string, errors.ServiceError) {
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

		expressions, err := util.ArrayMapWithError(and.Or.Expressions, func(t *model.BooleanExpression) (string, errors.ServiceError) {
			return r.applyCondition(resource, t)
		})
		if err != nil {
			return "", err
		}
		return r.builder.Or(expressions...), nil
	}

	if and, ok := query.Expression.(*model.BooleanExpression_Not); ok {
		return r.applyCondition(resource, and.Not)
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_Equal); ok {
		left, err := r.applyExpression(resource, equ.Equal.Left)
		if err != nil {
			return "", err
		}
		right, err := r.applyExpression(resource, equ.Equal.Right)
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
		left, err := r.applyExpression(resource, equ.GreaterThan.Left)
		if err != nil {
			return "", err
		}
		right, err := r.applyExpression(resource, equ.GreaterThan.Right)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s > %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_GreaterThanOrEqual); ok {
		left, err := r.applyExpression(resource, equ.GreaterThanOrEqual.Left)
		if err != nil {
			return "", err
		}
		right, err := r.applyExpression(resource, equ.GreaterThanOrEqual.Right)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s >= %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_LessThan); ok {
		left, err := r.applyExpression(resource, equ.LessThan.Left)
		if err != nil {
			return "", err
		}
		right, err := r.applyExpression(resource, equ.LessThan.Right)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s < %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_LessThanOrEqual); ok {
		left, err := r.applyExpression(resource, equ.LessThanOrEqual.Left)
		if err != nil {
			return "", err
		}
		right, err := r.applyExpression(resource, equ.LessThanOrEqual.Right)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s <= %s", left, right), nil
	}

	if equ, ok := query.Expression.(*model.BooleanExpression_In); ok {
		left, err := r.applyExpression(resource, equ.In.Left)
		if err != nil {
			return "", err
		}
		right, err := r.applyExpression(resource, equ.In.Right)
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

	panic("unknown boolean expression type: " + query.String())
}

func (r *recordLister) applyExpression(resource *model.Resource, query *model.Expression) (string, errors.ServiceError) {
	if query.Expression == nil {
		return "", errors.PropertyNotFoundError.WithDetails("expression is empty")
	}

	var additionalProperties = []string{
		"id", "version",
	}

	if propEx, ok := query.Expression.(*model.Expression_Property); ok {
		for _, ap := range additionalProperties {
			if ap == propEx.Property {
				return fmt.Sprintf("t." + ap), nil
			}
		}
		property := locatePropertyByName(resource, propEx.Property)

		if property == nil {
			return "", errors.PropertyNotFoundError.WithDetails(propEx.Property)
		}

		return fmt.Sprintf("t." + property.Mapping), nil
	}

	if propEx, ok := query.Expression.(*model.Expression_Value); ok {
		if propEx.Value.GetListValue() != nil {
			list := propEx.Value.GetListValue()
			var c []string
			for _, val := range list.Values {
				c = append(c, r.builder.Var(val.AsInterface()))
			}
			return strings.Join(c, ","), nil
		} else {
			return r.builder.Var(propEx.Value.AsInterface()), nil
		}
	}

	panic("unknown expression type: " + query.String())
}

func (r *recordLister) prepareCols() []string {
	var cols []string

	for _, cd := range r.colList {
		cols = append(cols, fmt.Sprintf("%s as %s", cd.def, cd.alias))
	}

	return cols
}

func recordList(ctx context.Context, runner QueryRunner, params abs.ListRecordParams) (result []*model.Record, total uint32, err errors.ServiceError) {
	return (&recordLister{
		ctx:               ctx,
		runner:            runner,
		resource:          params.Resource,
		query:             params.Query,
		Limit:             params.Limit,
		Offset:            params.Offset,
		UseHistory:        params.UseHistory,
		ResolveReferences: params.ResolveReferences,
		Schema:            *params.Schema,
		resultChan:        params.ResultChan,
		packRecords:       params.PackRecords,
	}).Exec()
}

func DbDecode(property *model.ResourceProperty, val interface{}) (interface{}, errors.ServiceError) {
	if property.Type == model.ResourcePropertyType_TYPE_OBJECT || property.Type == model.ResourcePropertyType_TYPE_ENUM || property.Type == model.ResourcePropertyType_TYPE_MAP || property.Type == model.ResourcePropertyType_TYPE_LIST {
		var data = new(interface{})
		err2 := json.Unmarshal([]byte(val.(string)), data)

		if err2 != nil {
			return nil, errors.InternalError.WithDetails(err2.Error())
		}

		val = *data
	} else if property.Type == model.ResourcePropertyType_TYPE_REFERENCE {
		return types.ReferenceType{
			"id": val,
		}, nil
	}

	return val, nil
}
