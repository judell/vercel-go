package user

import "github.com/chronark/vercel-go/endpoints/team"

type UserResponse struct {
	User User `json:"user"`
}

type User struct {
	Uid             string       `json:"uid"`
	Email           string       `json:"email"`
	Name            string       `json:"name"`
	Username        string       `json:"username"`
	Avatar          string       `json:"avatar"`
	PlatformVersion int          `json:"platformVersion"`
	Billing         team.Billing `json:"billing"`
	Bio             string       `json:"bio"`
	Website         string       `json:"website"`
	Profiles        []struct {
		Service string `json:"service"`
		Link    string `json:"link"`
	} `json:"profiles"`
	SoftBlock struct {
		BlockedAt int    `json:"blockedAt"`
		Reason    string `json:"reason"`
	} `json:"softBlock,omitempty"`
	RemoteCaching struct {
		Enabled bool `json:"enabled,omitempty"`
	} `json:"remoteCaching,omitempty"`
	Date                     string              `json:"date"`
	HasTrialAvailable        bool                `json:"hasTrialAvailable"`
	StagingPrefix            string              `json:"stagingPrefix"`
	ResourceConfig           team.ResourceConfig `json:"resourceConfig"`
	ImportFlowGitNamespace   interface{}         `json:"importFlowGitNamespace,omitempty"`
	ImportFlowGitNamespaceID interface{}         `json:"importFlowGitNamespaceId,omitempty"`
	ImportFlowGitProvider    string              `json:"importFlowGitProvider,omitempty"`
}
