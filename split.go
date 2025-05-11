package paystack

// SplitService handles operations related to transaction Splits
// For more details see https://paystack.com/docs/api/split/
type SplitService service

type SplitRequest struct {
	Name             string               `json:"name,omitempty"`
	Type             string               `json:"type,omitempty"`
	Currency         string               `json:"currency,omitempty"`
	Subaccounts      []BeneficiaryAccount `json:"subaccounts,omitempty"`
	BearerType       string               `json:"bearer_type,omitempty"`       // Any of "subaccount" | "account" | "all-proportional" | "all"
	BearerSubAccount string               `json:"bearer_subaccount,omitempty"` // SubAccountCode of bearer, if SubAccount
}

type BeneficiaryAccount struct {
	Subaccount SubAccount `json:"subaccount,omitempty"`
	Share          int    `json:"share,omitempty"`
}

type CreateSplitResponse struct {
	Status  bool              `json:"status,omitempty"`
	Message string            `json:"message,omitempty"`
	Data    SplitResponseData `json:"data,omitempty"`
}

type SplitResponseData struct {
	SplitID          int                  `json:"id,omitempty"`
	Name             string               `json:"name,omitempty"`
	Type             string               `json:"type,omitempty"`
	Currency         string               `json:"currency,omitempty"`
	Integration      int                  `json:"integration,omitempty"`
	Domain           string               `json:"domain,omitempty"`
	SplitCode        string               `json:"split_code,omitempty"`
	Active           bool                 `json:"active,omitempty"`
	BearerType       string               `json:"bearer_type,omitempty"`
	CreatedAt        string               `json:"created_at,omitempty"`
	UpdatedAt        string               `json:"updated_at,omitempty"`
	IsDynamic        bool                 `json:"is_dynamic,omitempty"`
	Subaccounts      []BeneficiaryAccount `json:"subaccounts,omitempty"`
	TotalSubAccounts int                  `json:"total_subaccounts,omitempty"`
}

func (s *SplitService) CreateSplit(request *SplitRequest) (*CreateSplitResponse, error) {
	url := "/split"
	response := &CreateSplitResponse{}
	err := s.client.Call("POST", url, request, response)
	return response, err
}
