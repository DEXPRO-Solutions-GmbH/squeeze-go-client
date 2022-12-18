package examples

import (
	"fmt"
	squeeze_go_client "github.com/dexpro-solutions-gmbh/squeeze-go-client"
	"github.com/rodaine/table"
)

type checkJobsResult struct {
	unlockActive  bool
	cleanupActive bool
}

func CheckJobs(tenants, apiKeys []string) {
	if len(tenants) != len(apiKeys) {
		panic("tenants and apiKeys must have the same length")
	}

	tbl := table.New("Tenant", "Cleanup", "Unlock")

	for i := 0; i < len(tenants); i++ {
		tenant := tenants[i]
		apiKey := apiKeys[i]

		basePath := fmt.Sprintf("https://%s/api/v2", tenant)

		c := squeeze_go_client.NewClient(basePath)
		c.ApiKey = apiKey
		jobs, err := checkJobs(c)
		if err != nil {
			panic(err)
		}

		tbl.AddRow(tenant, jobs.cleanupActive, jobs.unlockActive)
	}

	tbl.Print()
}

func checkJobs(client *squeeze_go_client.Client) (*checkJobsResult, error) {
	jobs, err := client.Jobs.GetAllJobs()
	if err != nil {
		return nil, err
	}

	result := &checkJobsResult{}

	for _, job := range jobs {
		if job.Name == "(Internal) cleanup" {
			result.cleanupActive = job.Active
		} else if job.Name == "(Internal) unlock-documents" {
			result.unlockActive = job.Active
		}
	}

	return result, nil
}
