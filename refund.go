package paystack

import "fmt"

type RefundService service

type Refund struct {
	Transaction    Transaction `json:"transaction,omitempty"`
	Integration    int         `json:"integration,omitempty"`
	DeductedAmount int         `json:"deducted_amount,omitempty"`
	Channel        interface{} `json:"channel,omitempty"` // TODO: Confirm data type
	MerchantNote   string      `json:"merchant_note,omitempty"`
	CustomerNote   string      `json:"customer_note,omitempty"`
	Status         string      `json:"status,omitempty"`
	RefundedBy     string      `json:"refunded_by,omitempty"`
	ExpectedAt     string      `json:"expected_at,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	Domain         string      `json:"domain,omitempty"`
	Amount         int         `json:"amount,omitempty"`
	FullyDeducted  bool        `json:"fully_deducted,omitempty"`
	Id             int         `json:"id,omitempty"`
	CreatedAt      string      `json:"createdAt,omitempty"`
	UpdatedAt      string      `json:"updatedAt,omitempty"`
}

type RefundRequest struct {
	Transaction  string `json:"transaction,omitempty"`   // Transaction reference or id
	Amount       int    `json:"amount,omitempty"`        // Optional: Defaults to original transaction amount
	Currency     string `json:"currency,omitempty"`      // Optional
	CustomerNote string `json:"customer_note,omitempty"` // Optional
	MerchantNote string `json:"merchant_note,omitempty"` // Optional
}

// RefundList is a list object for Splits.
type RefundList struct {
	Meta   ListMeta
	Values []Refund `json:"data"`
}

// Create and manage transaction refunds.
// For more details see https://paystack.com/docs/api/refund/#refunds
func (s *RefundService) CreateRefund(request *RefundRequest) (*Refund, error) {
	url := "/refund"
	refund := &Refund{}
	err := s.client.Call("POST", url, request, refund)
	return refund, err
}

// List refunds available on your integration
// For more details see https://paystack.com/docs/api/refund/#list
func (s *RefundService) List() (*RefundList, error) {
	return s.ListN(10, 1)
}

// List refunds available on your integration
// For more details see https://paystack.com/docs/api/refund/#list
func (s *RefundService) ListN(count, offset int) (*RefundList, error) {
	url := paginateURL("/refund", count, offset)
	refunds := &RefundList{}
	err := s.client.Call("GET", url, nil, refunds)
	return refunds, err
}

// Get details of a refund on your integration
// For more details see https://paystack.com/docs/api/refund/#fetch
func (s *RefundService) Get(id int) (*Refund, error) {
	url := fmt.Sprintf("/refund/%d", id)
	refund := &Refund{}
	err := s.client.Call("GET", url, nil, refund)
	return refund, err
}
