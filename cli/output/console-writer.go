package output

import (
	"data-handler/model"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
)

type consoleWriter struct {
	writer io.Writer
}

func (c consoleWriter) DescribeResource(resource *model.Resource) {
	const padding = 3
	w := tabwriter.NewWriter(c.writer, 0, 0, padding, ' ', 0)

	c.out(w, "Name: \t\t %s", resource.Name)
	c.out(w, "Workspace: \t\t %s", resource.Workspace)
	c.out(w, "Version: \t\t %d", resource.Version)
	c.out(w, "")

	c.out(w, "Source Config:")
	c.out(w, "  DataSource: \t\t %s", resource.SourceConfig.DataSource)
	c.out(w, "  Mapping: \t\t %s", resource.SourceConfig.Mapping)
	c.out(w, "")

	c.out(w, "AuditData:")
	c.out(w, "  Created By: \t\t %s", resource.AuditData.CreatedBy)
	c.out(w, "  Created On: \t\t %s", resource.AuditData.CreatedOn.AsTime().String())
	c.out(w, "  Updated By: \t\t %s", resource.AuditData.UpdatedBy)
	c.out(w, "  Updated On: \t\t %s", resource.AuditData.UpdatedOn.AsTime().String())
	c.out(w, "")

	c.out(w, "Properties:")

	var data [][]string

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Name", "Mapping", "Type", "Required", "Unique", "Primary", "Length"})
	c.configureTable(table)

	for _, item := range resource.Properties {
		mapping := item.SourceConfig.(*model.ResourceProperty_Mapping)

		typeStr := strings.ToLower(item.Type.String())[5:]

		data = append(data, []string{
			item.Name,
			mapping.Mapping.Mapping,
			typeStr,
			strconv.FormatBool(item.Required),
			strconv.FormatBool(item.Unique),
			strconv.FormatBool(item.Primary),
			strconv.Itoa(int(item.Length)),
		})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	c.out(w, "")
	c.out(w, "References:")

	table = tablewriter.NewWriter(w)
	table.SetHeader([]string{"Property", "Referenced Resource", "Cascade"})
	c.configureTable(table)

	data = [][]string{}

	for _, item := range resource.References {
		data = append(data, []string{
			item.PropertyName,
			item.ReferencedResource,
			strconv.FormatBool(item.Cascade),
		})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	w.Flush()
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

func (c consoleWriter) WriteResources(resources []*model.Resource) {
	var data [][]string

	table := tablewriter.NewWriter(c.writer)
	table.SetHeader([]string{"Name", "Workspace", "DataSource", "Mapping", "Version"})
	c.configureTable(table)

	for _, item := range resources {
		data = append(data, []string{
			item.Name,
			item.Workspace,
			item.SourceConfig.DataSource,
			item.SourceConfig.Mapping,
			string(item.Version),
		})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func (c consoleWriter) WriteRecords(record []*model.Record) {
	//TODO implement me
	panic("implement me")
}
