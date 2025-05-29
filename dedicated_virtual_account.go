package paystack

import (
	"fmt"
	"net/url"
)

type DedicatedVirtualAccountService service

type DedicatedVirtualAccount struct {
	Bank          Bank        `json:"bank,omitempty"`
	AccountName   string      `json:"account_name,omitempty"`
	AccountNumber string      `json:"account_number,omitempty"`
	Assigned      bool        `json:"assigned,omitempty"`
	Currency      string      `json:"currency,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Active        bool        `json:"active,omitempty"`
	Id            int         `json:"id,omitempty"`
	CreatedAt     string      `json:"created_at,omitempty"`
	UpdatedAt     string      `json:"updated_at,omitempty"`
	Assignment    Assignment  `json:"assignment,omitempty"`
	Customer      Customer    `json:"customer,omitempty"`
	SplitConfig   Split       `json:"split_config,omitempty"`
}

type Assignment struct {
	Integration  int    `json:"integration,omitempty"`
	AssigneeId   int    `json:"assignee_id,omitempty"`
	AssigneeType string `json:"assignee_type,omitempty"`
	Expired      bool   `json:"expired,omitempty"`
	AccountType  string `json:"account_type,omitempty"`
	AssignedAt   string `json:"assigned_at,omitempty"`
}

type DedicatedVirtualAccountRequest struct {
	Customer      int    `json:"customer,omitempty"`       // Customer ID
	PreferredBank string `json:"preferred_bank,omitempty"` // Optional: We currently support Wema Bank and Titan Paystack.
	SubAccount    string `json:"subaccount,omitempty"`     // Optional
	SplitCode     string `json:"split_code,omitempty"`     // Optional
	FirstName     string `json:"first_name,omitempty"`     // Optional
	LastName      string `json:"last_name,omitempty"`      // Optional
	Phone         string `json:"phone,omitempty"`          // Optional
}

type AssignDVARequest struct {
	Email         string `json:"email,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Country       string `json:"country,omitempty"`
	SubAccount    string `json:"subaccount,omitempty"` // Optional
	SplitCode     string `json:"split_code,omitempty"` // Optional
	AccountNumber string `json:"account_number,omitempty"`
	Bvn           string `json:"bvn,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
}

// DVAList is a list object for Dedicated Virtual Accounts.
type DVAList struct {
	Meta   ListMeta
	Values []DedicatedVirtualAccount `json:"data"`
}

// Filter for retrieving DVA list. All fields are optional
type DVAListFilter struct {
	Active       bool   `json:"active,omitempty"`
	Currency     string `json:"currency,omitempty"`
	ProviderSlug string `json:"provider_slug,omitempty"`
	BankId       string `json:"bank_id,omitempty"`
	Customer     string `json:"customer,omitempty"`
}

type RequeryDVARequest struct {
	AccountNumber string `json:"account_number,omitempty"`
	ProviderSlug  string `json:"provider_slug,omitempty"`
	Date          string `json:"date,omitempty"` // Optional
}

type DVATransactionSplitRequest struct {
	Customer      int `json:"customer,omitempty"`   // Customer ID or code
	SubAccount    string `json:"subaccount,omitempty"` // Subaccount code of the account you want to split the transaction with
	SplitCode     string `json:"split_code,omitempty"` // Split code consisting of the lists of accounts you want to split the transaction with
	PreferredBank string `json:"preferred_bank,omitempty"`
}

type BankProvider struct {
	ProviderSlug string `json:"provider_slug,omitempty"`
	BankId       int    `json:"bank_id,omitempty"`
	BankName     string `json:"bank_name,omitempty"`
	Id           int    `json:"id,omitempty"`
}

// Create a dedicated virtual account for an existing customer.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#create
func (s *DedicatedVirtualAccountService) Create(request *DedicatedVirtualAccountRequest) (*DedicatedVirtualAccount, error) {
	url := "/dedicated_account"
	dva := &DedicatedVirtualAccount{}
	err := s.client.Call("POST", url, request, dva)
	return dva, err
}

// Create a customer, validate the customer, and assign a DVA to the customer. The process is asynchronous - listen for response using webhooks.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#assign
func (s *DedicatedVirtualAccountService) Assign(request *AssignDVARequest) (*DedicatedVirtualAccount, error) {
	url := "/dedicated_account"
	dva := &DedicatedVirtualAccount{}
	err := s.client.Call("POST", url, request, dva)
	return dva, err
}

// List dedicated virtual accounts available on your integration.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#list
func (s *DedicatedVirtualAccountService) List(filter *DVAListFilter) (*DVAList, error) {
	return s.ListN(filter, 10, 1)
}

// List dedicated virtual accounts available on your integration.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#list
func (s *DedicatedVirtualAccountService) ListN(filter *DVAListFilter, count, offset int) (*DVAList, error) {
	url := paginateURL("/dedicated_account", count, offset)
	dvaList := &DVAList{}
	err := s.client.Call("GET", url, nil, dvaList)
	return dvaList, err
}

// Get details of a dedicated virtual account on your integration.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#fetch
func (s *DedicatedVirtualAccountService) Get(id int) (*DedicatedVirtualAccount, error) {
	url := fmt.Sprintf("/dedicated_account/%d", id)
	dva := &DedicatedVirtualAccount{}
	err := s.client.Call("GET", url, nil, dva)
	return dva, err
}

// Requery Dedicated Virtual Account for new transactions.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#requery
func (s *DedicatedVirtualAccountService) Requery(request *RequeryDVARequest) (*DedicatedVirtualAccount, error) {
	url := fmt.Sprintf("/dedicated_account/requery?account_number=%s&provider_slug=%s&date=%s", request.AccountNumber, request.ProviderSlug, request.Date)
	dva := &DedicatedVirtualAccount{}
	err := s.client.Call("GET", url, nil, dva)
	return dva, err
}

// Deactivate a dedicated virtual account on your integration.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#deactivate
func (s *DedicatedVirtualAccountService) Deactivate(id int) (*DedicatedVirtualAccount, error) {
	url := fmt.Sprintf("/dedicated_account/:%d", id)
	dva := &DedicatedVirtualAccount{}
	err := s.client.Call("DELETE", url, nil, dva)
	return dva, err
}

// Split a dedicated virtual account transaction with one or more accounts.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#add-split
func (s *DedicatedVirtualAccountService) Split(request *DVATransactionSplitRequest) (*DedicatedVirtualAccount, error) {
	url := "/dedicated_account"
	dva := &DedicatedVirtualAccount{}
	err := s.client.Call("POST", url, request, dva)
	return dva, err
}

// Remove split payments for transactions on a dedicated virtual account
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#remove-split
func (s *DedicatedVirtualAccountService) RemoveSplit(acct string) (*DedicatedVirtualAccount, error) {
	u := "/dedicated_account/split"
	dva := &DedicatedVirtualAccount{}
	req := url.Values{}
	req.Add("account_number", acct)
	err := s.client.Call("DELETE", u, req, dva)
	return dva, err
}

// Get available bank providers for a dedicated virtual account
// For more details see https://paystack.com/docs/api/dedicated-virtual-account/#providers
func (s *DedicatedVirtualAccountService) GetBankProviders() ([]BankProvider, error) {
	url := "/dedicated_account/available_providers"
	providers := []BankProvider{}
	err := s.client.Call("GET", url, nil, providers)
	return providers, err
}
