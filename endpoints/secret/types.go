package secret

type Secret struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The name of the secret. This is what you could use in your environment variables after a @.
	Name string `json:"name"`

	// The date when the secret was created (ISO 8601).
	Created string `json:"created"`

	// The date when the secret was created (timestamp).
	CreatedAt int `json:"createdAt"`

	// The unique identifier of the team the secret was created for.
	TeamID string `json:"teamId,omitempty"`

	// The unique identifier of the user who created the secret.
	UserID string `json:"userId,omitempty"`

	// The value of the secret.
	Value string `json:"value,omitempty"`

	// The unique identifier of the project which the secret belongs to.
	ProjectID string `json:"projectId,omitempty"`

	// Indicates whether the secret value can be decrypted after it has been created.
	Decryptable bool `json:"decryptable,omitempty"`
}

type ListSecretsRequest struct {
	// Maximum number of secrets to list from a request.
	// Required: No
	Limit int `json:"limit,omitempty"`

	// Get secrets created after this JavaScript timestamp.
	// Required: No
	Since int `json:"since,omitempty"`

	// Get secrets created before this JavaScript timestamp.
	// Required: No
	Until int `json:"until,omitempty"`
}

type ListSecretsResponse struct {
	Secrets    []Secret `json:"secrets"`
	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination"`
}

// Either `Name` or `id` must be provided
type GetSecretRequest struct {
	// The name of the secret.
	// Required: false
	Name string

	// The unique identifier of the secret.
	// Required: false
	Id string
}

type GetSecretResponse Secret

/*
type GetSecretResponse struct {
	// The unique identifier of the secret.
	Uid string `json:"uid"`

	// The name of the secret.
	Name string `json:"name"`

	// The team unique identifier to which the secret belongs to.
	Teamid string `json:"teamid"`

	// The unique identifier of the user who created the secret.
	Userid string `json:"userid"`

	// The date when the secret was created.
	Created string `json:"created"`

	// The date when the secret was created in milliseconds since the UNIX epoch.
	CreatedAt int `json:"createdAt"`
}
*/

type CreateSecretsRequest struct {
	// The name of the secret (max 100 characters).
	// Required: Yes
	Name string `json:"name"`
	// The value of the new secret.
	// Required: Yes
	Value string `json:"value"`
}

type CreateSecretsResponse struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The name of the secret.
	Name string `json:"name"`

	// The date when the secret was created.
	Created string `json:"created"`

	// The unique identifier of the user who created the secret.
	Userid string `json:"userid"`

	// A map with the value of the secret.
	Value struct {
		// The type of structure used to save the secret value (always Buffer).
		Type string `json:"type"`

		// 	A list of numbers which could be used to recreate the secret value.
		data []int `son:"data"`
	} `json:"value"`
}

type RenameSecretRequest struct {
	Name string

	// The new name
	// Required: Yes
	NewName string
}

type RenameSecretResponse struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The new name of the secret.
	Name string `json:"name"`

	// The date when the secret was created.
	Created string `json:"created"`

	// The old name of the secret.
	OldName string `json:"oldName"`
}

type DeleteSecretRequest struct {
	// The name of the secret
	// Required: Yes
	Name string `json:"name"`
}

type DeleteSecretResponse struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The new name of the secret.
	Name string `json:"name"`

	// The date when the secret was created.
	Created int `json:"created"`
}
