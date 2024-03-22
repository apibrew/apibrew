package output

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/olekukonko/tablewriter"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
)

type consoleWriter struct {
	writer   io.Writer
	describe bool
}

func (c consoleWriter) IsBinary() bool {
	return false
}

func (c consoleWriter) DescribeResource(resource *resource_model.Resource) {
	const padding = 3
	w := tabwriter.NewWriter(c.writer, 0, 0, padding, ' ', 0)

	c.out(w, "Name: \t\t %s", resource.Name)
	c.out(w, "Namespace: \t\t %s", resource.Namespace)
	c.out(w, "Version: \t\t %d", resource.Version)
	c.out(w, "")

	c.out(w, "Source Config:")
	c.out(w, "  DataSource: \t\t %s", resource.DataSource)
	c.out(w, "  Catalog: \t\t %s", resource.Catalog)
	c.out(w, "  Entity: \t\t %s", resource.Entity)
	c.out(w, "")

	if resource.AuditData != nil {
		c.out(w, "AuditData:")
		c.out(w, "  Created By: \t\t %s", resource.AuditData.CreatedBy)
		c.out(w, "  Created On: \t\t %s", resource.AuditData.CreatedOn.String())
		c.out(w, "  Updated By: \t\t %s", resource.AuditData.UpdatedBy)
		c.out(w, "  Updated On: \t\t %s", resource.AuditData.UpdatedOn.String())
		c.out(w, "")
	}

	if len(resource.Annotations) > 0 {
		c.out(w, "Annotations:")
		for key, value := range resource.Annotations {
			c.out(w, fmt.Sprintf("%s:\t%s", key, value))
		}
		c.out(w, "")
	}

	c.out(w, "Properties:")

	var data [][]string

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Name", "Type", "Required", "Unique", "Length", "Annotations"})
	c.configureTable(table)

	for itemName, item := range resource.Properties {

		typeStr := strings.ToLower(string(item.Type))

		data = append(data, []string{
			itemName,
			typeStr,
			strconv.FormatBool(item.Required),
			strconv.FormatBool(item.Unique),
			strconv.Itoa(int(item.Length)),
			annotations.ToString(annotations.FromMap(item.Annotations)),
		})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	_ = w.Flush()

	if len(resource.Indexes) > 0 {
		table = tablewriter.NewWriter(w)

		c.out(c.writer, "")
		c.out(c.writer, "Indexes:")

		data = [][]string{}
		table.SetHeader([]string{"IndexType", "Unique", "Properties", "Annotations"})
		c.configureTable(table)

		for _, item := range resource.Indexes {
			data = append(data, []string{
				string(*item.IndexType),
				strconv.FormatBool(util.DePointer(item.Unique, false)),
				strings.Join(util.ArrayMapToString(item.Properties, func(t resource_model.ResourceIndexProperty) string {
					return t.Name
				}), ", "),
				annotations.ToString(annotations.FromMap(item.Annotations)),
			})
		}

		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	}

	c.out(w, "")
	_ = w.Flush()
}

func (c consoleWriter) out(w io.Writer, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format+"\n", a...)
}

func (c consoleWriter) configureTable(table *tablewriter.Table) {
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
}

func (c consoleWriter) WriteResource(resources ...*resource_model.Resource) error {
	if c.describe {
		for _, resource := range resources {
			c.DescribeResource(resource)
		}
	} else {
		c.ShowResourceTable(resources)
	}

	return nil
}

func (c consoleWriter) ShowResourceTable(resources []*resource_model.Resource) {
	var data [][]string

	table := tablewriter.NewWriter(c.writer)
	table.SetHeader([]string{"Id", "Name", "Namespace", "DataSource", "Catalog", "Entity", "Version"})
	c.configureTable(table)

	for _, item := range resources {
		data = append(data, []string{
			item.Id.String(),
			item.Name,
			item.Namespace.Name,
			item.DataSource.Name,
			util.DePointer(item.Catalog, ""),
			util.DePointer(item.Entity, ""),
			strconv.Itoa(int(item.Version)),
		})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func (c consoleWriter) WriteRecords(resource *resource_model.Resource, total uint32, records []unstructured.Unstructured) error {
	table := tablewriter.NewWriter(c.writer)
	var columns []string

	for propName := range resource.Properties {
		columns = append(columns, util.ToDashCase(propName))
	}

	table.SetHeader(columns)
	c.configureTable(table)

	var i = 0
	for _, item := range records {
		var row []string

		for propName := range resource.Properties {
			row = append(row, fmt.Sprintf("%v", item[propName]))
		}
		i++

		table.Append(row)

		if i%1000 == 0 {
			table.Render()
		}
	}

	table.Render()

	return nil
}
