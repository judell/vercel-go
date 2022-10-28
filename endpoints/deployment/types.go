package deployment
	
type Deployment struct {
		UID     string `json:"uid"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		Created int64  `json:"created"`
		Source  string `json:"source"`
		State   string `json:"state"`
		Type    string `json:"type"`
		Creator struct {
			UID         string `json:"uid"`
			Email       string `json:"email"`
			Username    string `json:"username"`
			GithubLogin string `json:"githubLogin"`
		} `json:"creator"`
		InspectorURL string `json:"inspectorUrl"`
		Meta         struct {
			GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
			GithubCommitMessage     string `json:"githubCommitMessage"`
			GithubCommitOrg         string `json:"githubCommitOrg"`
			GithubCommitRef         string `json:"githubCommitRef"`
			GithubCommitRepo        string `json:"githubCommitRepo"`
			GithubCommitSha         string `json:"githubCommitSha"`
			GithubDeployment        string `json:"githubDeployment"`
			GithubOrg               string `json:"githubOrg"`
			GithubRepo              string `json:"githubRepo"`
			GithubCommitRepoID      string `json:"githubCommitRepoId"`
			GithubRepoID            string `json:"githubRepoId"`
			GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
			GithubPrID              string `json:"githubPrId"`
		} `json:"meta,omitempty"`
		Target              interface{} `json:"target"`
		AliasError          interface{} `json:"aliasError"`
		AliasAssigned       int64       `json:"aliasAssigned"`
		IsRollbackCandidate bool        `json:"isRollbackCandidate"`
		CreatedAt           int64       `json:"createdAt"`
		BuildingAt          int64       `json:"buildingAt"`
		Ready               int64       `json:"ready"`
	}

type ListDeploymentsRequest struct {
	// Maximum number of records to list from a request.
	// Required: No
	Limit int
}

type ListDeploymentsResponse struct {
	Deployments   []Deployment `json:"deployments"`
	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination"`
}