package team

import "time"

// Either `id` or `Slug` must be defined
type GetTeamRequest struct {
	// The team unique identifier. Always prepended by team_.
	// Required: false
	ID string

	// The team slug. A slugified version of the name.
	Slug string
}

type Billing struct {
	Address struct {
		State      string `json:"state"`
		Country    string `json:"country"`
		City       string `json:"city"`
		PostalCode string `json:"postalCode"`
		Line1      string `json:"line1"`
	} `json:"address"`
	Cancelation interface{} `json:"cancelation"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Period      struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
	} `json:"period"`
	Plan         string      `json:"plan"`
	Tax          interface{} `json:"tax"`
	Language     string      `json:"language,omitempty"`
	Currency     string      `json:"currency"`
	Trial        interface{} `json:"trial"`
	InvoiceItems struct {
		Pro struct {
			Price    int  `json:"price"`
			Quantity int  `json:"quantity"`
			Hidden   bool `json:"hidden"`
		} `json:"pro"`
		TeamSeats struct {
			Hidden   bool `json:"hidden"`
			Price    int  `json:"price"`
			Quantity int  `json:"quantity"`
		} `json:"teamSeats"`
		ConcurrentBuilds struct {
			Hidden   bool `json:"hidden"`
			Price    int  `json:"price"`
			Quantity int  `json:"quantity"`
		} `json:"concurrentBuilds"`
	} `json:"invoiceItems"`
	Addons        interface{} `json:"addons"`
	Platform      string      `json:"platform"`
	ProgramType   string      `json:"programType,omitempty"`
	Overdue       interface{} `json:"overdue"`
	Subscriptions []struct {
		ID     string      `json:"id"`
		Trial  interface{} `json:"trial"`
		Period struct {
			Start int64 `json:"start"`
			End   int64 `json:"end"`
		} `json:"period"`
		Frequency struct {
			Interval      string `json:"interval"`
			IntervalCount int    `json:"intervalCount"`
		} `json:"frequency"`
		Discount interface{} `json:"discount"`
		Items    []struct {
			ID        string `json:"id"`
			Priceid   string `json:"priceid"`
			Productid string `json:"productid"`
			Amount    int    `json:"amount"`
			Quantity  int    `json:"quantity"`
		} `json:"items"`
	} `json:"subscriptions"`
	Controls      interface{} `json:"controls,omitempty"`
	PurchaseOrder string      `json:"purchaseOrder,omitempty"`
	Status        string      `json:"status,omitempty"`
	Contract      interface{} `json:"contract,omitempty"`
}

type ResourceConfig struct {
	NodeType         string `json:"nodeType,omitempty"`
	ConcurrentBuilds int    `json:"concurrentBuilds,omitempty"`
	AwsAccountType   string `json:"awsAccountType,omitempty"`
	AwsAccountIds    string `json:"awsAccountIds,omitempty"`
	CfZoneName       string `json:"cfZoneName,omitempty"`
}

type Team struct {
	ID         string    `json:"id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Creatorid  string    `json:"creatorid"`
	Created    time.Time `json:"created"`
	CreatedAt  int64     `json:"createdAt"`
	UpdatedAt  int64     `json:"updatedAt"`
	Avatar     string    `json:"avatar"`
	Membership struct {
		Role      string `json:"role"`
		Confirmed bool   `json:"confirmed"`
		Created   int64  `json:"created"`
		CreatedAt int64  `json:"createdAt"`
		Uid       string `json:"uid"`
		UpdatedAt int64  `json:"updatedAt"`
	} `json:"membership"`
	PlatformVersion         interface{}    `json:"platformVersion"`
	InviteCode              string         `json:"inviteCode"`
	Billing                 Billing        `json:"billing"`
	Description             interface{}    `json:"description"`
	Profiles                []interface{}  `json:"profiles"`
	StagingPrefix           string         `json:"stagingPrefix"`
	ResourceConfig          ResourceConfig `json:"resourceConfig"`
	PreviewDeploymentSuffix interface{}    `json:"previewDeploymentSuffix"`
	SoftBlock               interface{}    `json:"softBlock"`
	AllowProjectTransfers   bool           `json:"allowProjectTransfers"`
}

type GetTeamResponse Team

type ListTeamsRequest struct {
	// Maximum number of teams to list from a request.
	// Required: No
	Limit int `json:"limit,omitempty"`

	// Get teams created after this JavaScript timestamp.
	// Required: No
	Since int `json:"since,omitempty"`

	// Get teams created before this JavaScript timestamp.
	// Required: No
	Until int `json:"until,omitempty"`
}

type ListTeamsResponse struct {
	Teams      []Team `json:"teams"`
	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination"`
}
