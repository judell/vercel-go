package team

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type TeamHandler struct {
	vercelClient api.VercelClient
}

func New(vercelClient api.VercelClient) *TeamHandler {
	return &TeamHandler{
		vercelClient,
	}
}

// List all teams
func (h *TeamHandler) ListTeams(req ListTeamsRequest) (res ListTeamsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v2/teams", &res)
	/*
		if h.teamid != "" {
			apiRequest.Query.Add("teamId", h.teamid)
		}
	*/
	if req.Limit != 0 {
		apiRequest.Query.Add("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Until != 0 {
		apiRequest.Query.Add("until", fmt.Sprintf("%d", req.Until))
	}
	if req.Since != 0 {
		apiRequest.Query.Add("since", fmt.Sprintf("%d", req.Since))
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return ListTeamsResponse{}, fmt.Errorf("Unable to fetch teams from vercel: %w", err)
	}
	return res, nil
}

// Return the team info
func (h *TeamHandler) Get(req GetTeamRequest) (res GetTeamResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v2/teams", &res)

	// Teams can be references by their id or slug.
	// If id the url looks like this: .../<id>
	// Or when using a slug: ...?slug=<slug>
	if req.ID != "" {
		apiRequest.Path = fmt.Sprintf("%s/%s", apiRequest.Path, req.ID)
	} else if req.Slug != "" {

		apiRequest.Query.Add("slug", req.Slug)
	} else {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team: Either `id` or `Slug` must be defined")
	}

	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team from vercel: %w", err)
	}
	return res, nil
}
