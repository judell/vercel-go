package deployment_test

import (
	"os"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/deployment"
)

type testCase struct {
	name    string
	handler *deployment.DeploymentHandler
}

var (
	token string
	tests []testCase
)

func init() {
	token = os.Getenv("VERCEL_TOKEN")
	if token == "" {
		panic("VERCEL_TOKEN must be defined")
	}
	teamid := os.Getenv("VERCEL_TEAM_ID")
	if teamid == "" {
		panic("VERCEL_TEAM_ID must be defined")
	}

	tests = []testCase{
		{name: "team", handler: deployment.New(api.New(api.NewClientConfig{Token: token}), teamid)},
	}
}

func TestListDeployments(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.handler.List(deployment.ListDeploymentsRequest{})

			require.NoError(t, err)
			require.True(t, len(res.Deployments) > 0)
		})
	}
}

func TestListDeploymentsPaginated(t *testing.T) {
	max := 227
	req := deployment.ListDeploymentsRequest{Limit: 100}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			total := 0
			for {
				res, err := tt.handler.List(req)
				require.NoError(t, err)
				for i := 0; i <= len(res.Deployments); i++ {
					total += 1
					if total == max {
						break
					}
				}
				req.Until = int(res.Pagination.Next)
				if total == max {
					break
				}
			}

			require.True(t, total == max)
		})
	}
}
