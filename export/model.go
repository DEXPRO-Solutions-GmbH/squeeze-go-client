package export

import squeeze_go_client "github.com/dexpro-solutions-gmbh/squeeze-go-client"

type DocumentClassExport struct {
	DocumentClass *squeeze_go_client.DocumentClassDto     `json:"document_class"`
	Fields        []*squeeze_go_client.DocumentField      `json:"fields"`
	FieldGroups   []*squeeze_go_client.DocumentFieldGroup `json:"field_groups"`
	Tables        []*squeeze_go_client.DocumentTable      `json:"tables"`
}
