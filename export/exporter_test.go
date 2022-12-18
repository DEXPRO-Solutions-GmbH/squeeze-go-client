package export

import (
	squeeze_go_client "github.com/dexpro-solutions-gmbh/squeeze-go-client"
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExporter_ExportDocumentClass(t *testing.T) {
	c := squeeze_go_client.NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)

	exporter := NewExporter(c)
	export, err := exporter.ExportDocumentClass("Invoices")
	assert.NoError(t, err)
	assert.NotNil(t, export)

	assert.NotNil(t, export.DocumentClass)
	assert.NotEmpty(t, export.FieldGroups)
	assert.NotEmpty(t, export.Fields)
	assert.NotEmpty(t, export.Tables)
}
