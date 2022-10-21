package deployment

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type DeploymentHandler struct {
	vercelClient api.VercelClient
	teamid       string
}

func New(vercelClient api.VercelClient, teamid string) *DeploymentHandler {
	return &DeploymentHandler{
		vercelClient,
		teamid,
	}
}

func (h *DeploymentHandler) ListDeployments() (res ListDeploymentsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v6/deployments/", &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return ListDeploymentsResponse{}, fmt.Errorf("unable to fetch deployments: %w", err)
	}
	return res, nil
}


// Get a list of the Deployments you currently have under your account.
// By default, it returns the last 20 Deployments. The rest can be retrieved
// using the pagination options described below.
// NOTE: The order is always based on the updatedAt field of the Deployment.
func (h *DeploymentHandler) List() (res ListDeploymentsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v6/Deployments", &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return ListDeploymentsResponse{}, fmt.Errorf("Unable to fetch Deployments: %w", err)
	}
	return res, nil
}

