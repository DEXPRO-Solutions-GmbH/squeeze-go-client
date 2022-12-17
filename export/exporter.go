package export

import (
	"fmt"
	"github.com/dexpro-solutions-gmbh/squeeze-go-client"
)

// Exporter is a crude exporter which enables exporting of multiple api resources which normally you would have
// to call multiple api endpoints for.
type Exporter struct {
	client *squeeze_go_client.Client
}

func NewExporter(client *squeeze_go_client.Client) *Exporter {
	return &Exporter{client: client}
}

func (e *Exporter) ExportDocumentClass(className string) (*DocumentClassExport, error) {
	if classes, err := e.client.DocumentClass.GetDocumentClasses(); err != nil {
		return nil, fmt.Errorf("failed to get document classes: %w", err)
	} else {
		for _, class := range classes {
			if class.Name == className {
				return e.exportDocumentClass(class)
			}
		}
	}

	return nil, fmt.Errorf("document class %s not found", className)
}

func (e *Exporter) exportDocumentClass(class *squeeze_go_client.DocumentClassDto) (*DocumentClassExport, error) {
	export := &DocumentClassExport{
		DocumentClass: class,
		Fields:        nil,
		FieldGroups:   nil,
		Tables:        nil,
	}

	if fieldGroups, err := e.client.DocumentClass.GetAllFieldGroups(class.Id); err != nil {
		return nil, fmt.Errorf("failed to get field groups: %w", err)
	} else {
		export.FieldGroups = fieldGroups
	}

	if fields, err := e.client.DocumentClass.GetAllDocumentClassFields(class.Id); err != nil {
		return nil, fmt.Errorf("failed to get fields: %w", err)
	} else {
		export.Fields = fields
	}

	if tables, err := e.client.DocumentClass.GetAllDocumentClassTables(class.Id); err != nil {
		return nil, fmt.Errorf("failed to get tables: %w", err)
	} else {
		export.Tables = tables
	}

	for _, table := range export.Tables {
		if columns, err := e.client.DocumentClass.GetAllDocumentClassTableColumns(class.Id, table.Id); err != nil {
			return nil, fmt.Errorf("failed to get columns of table %s: %w", table.Name, err)
		} else {
			table.Columns = columns
		}
	}

	return export, nil
}
