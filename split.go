package paystack

import "fmt"

// SplitService handles operations related to transaction Splits
// For more details see https://paystack.com/docs/api/split/
type SplitService service

// Represents a Paystack Split payment
type Split struct {
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

// SplitRequest represents a request to create a transaction Split
type SplitRequest struct {
	Name             string                      `json:"name,omitempty"`
	Type             string                      `json:"type,omitempty"`
	Currency         string                      `json:"currency,omitempty"`
	Subaccounts      []BeneficiaryAccountRequest `json:"subaccounts,omitempty"`
	BearerType       string                      `json:"bearer_type,omitempty"`       // Any of "subaccount", "account", "all-proportional","all"
	BearerSubAccount string                      `json:"bearer_subaccount,omitempty"` // SubAccountCode of bearer, if SubAccount
}

// SplitList is a list object for Splits.
type SplitList struct {
	Meta   ListMeta
	Values []Split `json:"data"`
}

// Represents a request to update a split
type SplitUpdateRequest struct {
	Name             string `json:"name,omitempty"`
	Active           bool   `json:"active,omitempty"`
	BearerType       string `json:"bearer_type,omitempty"`       // Any of "subaccount", "account", "all-proportional","all".
	BearerSubAccount string `json:"bearer_subaccount,omitempty"` // SubAccountCode of bearer, if SubAccount.
}

// BeneficiaryAccount represents a SubAccount paired with its allocated share of the Split
type BeneficiaryAccount struct {
	Subaccount SubAccount `json:"subaccount,omitempty"`
	Share      int        `json:"share,omitempty"`
}

// Represents a SubAccount code paired with its allocated share of the split. Used in requests to create Splits.
type BeneficiaryAccountRequest struct {
	SubAccountCode string `json:"subaccount,omitempty"`
	Share          int `json:"share,omitempty"`
}

// Create a split payment on your integration
// For more details see https://paystack.com/docs/api/split/#create
func (s *SplitService) CreateSplit(request *SplitRequest) (*Split, error) {
	url := "/split"
	response := &Split{}
	err := s.client.Call("POST", url, request, response)
	return response, err
}

// List available transaction Splits
// For more details see https://paystack.com/docs/api/split/#list
func (s *SplitService) List() (*SplitList, error) {
	return s.ListN(10, 1)
}

// List available transaction Splits
// For more details see https://paystack.com/docs/api/split/#list
func (s *SplitService) ListN(count, offset int) (*SplitList, error) {
	url := paginateURL("/split", count, offset)
	splits := &SplitList{}
	err := s.client.Call("GET", url, nil, splits)
	return splits, err
}

// Get details of Split with the specified id
// For more details see https://paystack.com/docs/api/split/#fetch
func (s *SplitService) Get(id int) (*Split, error) {
	url := fmt.Sprintf("/split/%d", id)
	split := &Split{}
	err := s.client.Call("GET", url, nil, split)
	return split, err
}

// Update a transaction split details on your integration
// For more details see https://paystack.com/docs/api/split/#update
func (s *SplitService) Update(id int, request *SplitUpdateRequest) (*Split, error) {
	url := fmt.Sprintf("split/%d", id)
	split := &Split{}
	err := s.client.Call("PUT", url, request, split)
	return split, err
}

// Add a Subaccount to a Transaction Split, or update the share of an existing Subaccount in a Transaction Split
// For more details see https://paystack.com/docs/api/split/#add-subaccount
func (s *SplitService) UpdateSubAccounts(splitID int, subAccountCode string, share int) (*Split, error) {
	url := fmt.Sprintf("split/%d/subaccount/add", splitID)
	split := &Split{}
	requestData := map[string]interface{}{
		"subaccount": subAccountCode,
		"share":      share,
	}
	err := s.client.Call("POST", url, requestData, split)
	return split, err
}

// Remove a subaccount from a transaction split
// For more details see https://paystack.com/docs/api/split/#remove-subaccount
func (s *SplitService) RemoveSubAccount(splitID int, subAccountCode string) error {
	url := fmt.Sprintf("split/%d/subaccount/remove", splitID)
	err := s.client.Call("POST", url, subAccountCode, nil)
	return err
}
