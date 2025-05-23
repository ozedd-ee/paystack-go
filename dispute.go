package paystack

import (
	"fmt"
	"time"
)

type DisputeService service

type DisputeMessage struct {
	Sender    string `json:"sender,omitempty"`
	Body      string `json:"body,omitempty"`
	Dispute   int    `json:"dispute,omitempty"`
	Id        int    `json:"id,omitempty"`
	IsDeleted int    `json:"is_deleted,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type DisputeState struct {
	Id        int    `json:"id,omitempty"`
	Dispute   int    `json:"dispute,omitempty"`
	Status    string `json:"status,omitempty"`
	By        string `json:"by,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type Dispute struct {
	Currency               string           `json:"currency,omitempty"`
	Last4                  string           `json:"last4,omitempty"`
	Bin                    string           `json:"bin,omitempty"` // Verify data type
	TransactionReference   string           `json:"transaction_reference,omitempty"`
	MerchantTransactionRef string           `json:"merchant_transaction_reference,omitempty"`
	RefundAmount           int              `json:"refund_amount,omitempty"`
	Status                 string           `json:"status,omitempty"`
	Domain                 string           `json:"domain,omitempty"`
	Resolution             string           `json:"resolution,omitempty"`
	Category               string           `json:"category,omitempty"`
	Note                   string           `json:"note,omitempty"`
	Attachments            interface{}      `json:"attachments,omitempty"` // Verify data type
	Id                     int              `json:"id,omitempty"`
	Integration            int              `json:"integration,omitempty"`
	CreatedBy              string           `json:"created_by,omitempty"`
	Evidence               DisputeEvidence  `json:"evidence,omitempty"`
	ResolvedAt             string           `json:"resolvedAt,omitempty"`
	CreatedAt              string           `json:"createdAt,omitempty"`
	UpdatedAt              string           `json:"updatedAt,omitempty"`
	DueAt                  string           `json:"dueAt,omitempty"`
	Transaction            Transaction      `json:"transaction,omitempty"`
	Messages               []DisputeMessage `json:"messages,omitempty"`
	History                []DisputeState   `json:"history,omitempty"`
}

// DisputeList is a list object for disputes.
type DisputeList struct {
	Meta   ListMeta
	Values []Dispute `json:"data"`
}

type DisputeEvidence struct {
	CustomerEmail   string `json:"customer_email,omitempty"`
	CustomerName    string `json:"customer_name,omitempty"`
	CustomerPhone   string `json:"customer_phone,omitempty"`
	ServiceDetails  string `json:"service_details,omitempty"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	Dispute         int    `json:"dispute,omitempty"` // Dispute ID
	Id              int    `json:"id,omitempty"`      // Evidence ID
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

type UpdateDisputeRequest struct {
	RefundAmount     int    `json:"refund_amount,omitempty"`
	UploadedFilename string `json:"uploaded_filename,omitempty"` // Optional
}

type AddDisputeEvidenceRequest struct {
	CustomerEmail   string `json:"customer_email,omitempty"`
	CustomerName    string `json:"customer_name,omitempty"`
	CustomerPhone   string `json:"customer_phone,omitempty"`
	ServiceDetails  string `json:"service_details,omitempty"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
}

type ResolveDisputeRequest struct {
	Resolution       string `json:"resolution,omitempty"`
	Message          string `json:"message,omitempty"`
	UploadedFilename string `json:"uploaded_filename,omitempty"`
	RefundAmount     int    `json:"refund_amount,omitempty"`
	Evidence         int    `json:"evidence,omitempty"` // Evidence id
}

// All fields are optional
type DisputeFilterOptions struct {
	From        time.Time `json:"from,omitempty"`
	To          time.Time `json:"to,omitempty"`
	Transaction string    `json:"transaction,omitempty"` // Transaction ID
	Status      string    `json:"status,omitempty"`      // Filter dispute by status
}

type Upload struct {
	SignedUrl string `json:"signedUrl,omitempty"`
	FileName  string `json:"filename,omitempty"`
}

type Export struct {
	Path      string `json:"path,omitempty"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

// List disputes filed against you.
// For more details see https://paystack.com/docs/api/dispute/#list
func (s *DisputeService) List(options *DisputeFilterOptions) (*DisputeList, error) {
	return s.ListN(options, 10, 1)
}

// List disputes filed against you.
// For more details see https://paystack.com/docs/api/dispute/#list
func (s *DisputeService) ListN(options *DisputeFilterOptions, count, offset int) (*DisputeList, error) {
	url := paginateURL("/dispute", count, offset)
	disputes := &DisputeList{}
	err := s.client.Call("GET", url, options, disputes)
	return disputes, err
}

// Get details of Dispute with the specified id.
// For more details see https://paystack.com/docs/api/dispute/#fetch
func (s *DisputeService) Get(id int) (*Dispute, error) {
	url := fmt.Sprintf("/dispute/%d", id)
	dispute := &Dispute{}
	err := s.client.Call("GET", url, nil, dispute)
	return dispute, err
}

// Retrieve disputes for a particular transaction.
// For more details see https://paystack.com/docs/api/dispute/#transaction
func (s *DisputeService) ListTransactionDisputes(id int) (*Dispute, error) {
	url := fmt.Sprintf("/dispute/transaction/%d", id)
	dispute := &Dispute{}
	err := s.client.Call("GET", url, nil, dispute)
	return dispute, err
}

// Update details of a dispute on your integration.
// For more details see https://paystack.com/docs/api/dispute/#update
func (s *DisputeService) Update(id int, request *UpdateDisputeRequest) (*Dispute, error) {
	url := fmt.Sprintf("dispute/%d", id)
	dispute := &Dispute{}
	err := s.client.Call("PUT", url, request, dispute)
	return dispute, err
}

// Provide evidence for a dispute.
// For more details see https://paystack.com/docs/api/dispute/#evidence
func (s *DisputeService) AddDisputeEvidence(id int, request *AddDisputeEvidenceRequest) (*DisputeEvidence, error) {
	url := fmt.Sprintf("dispute/%d/evidence", id)
	evidence := &DisputeEvidence{}
	err := s.client.Call("POST", url, request, evidence)
	return evidence, err
}

// Resolve a dispute on your integration.
// For more details see https://paystack.com/docs/api/dispute/#resolve
func (s *DisputeService) ResolveDispute(id int, request *ResolveDisputeRequest) (*Dispute, error) {
	url := fmt.Sprintf("dispute/%d/resolve", id)
	dispute := &Dispute{}
	err := s.client.Call("PUT", url, request, dispute)
	return dispute, err
}

// Retrieve signed upload URL for dispute evidence documents.
// For more details see https://paystack.com/docs/api/dispute/#upload-url
func (s *DisputeService) GetUploadURL(id int, uploadFilename string) (*Upload, error) {
	url := fmt.Sprintf("dispute/:%d/upload_url?upload_filename=%s", id, uploadFilename)
	upload := &Upload{}
	err := s.client.Call("GET", url, nil, upload)
	return upload, err
}

// Export disputes available on your integration.
// For more details see https://paystack.com/docs/api/dispute/#export
func (s *DisputeService) Export(options *DisputeFilterOptions) (*Export, error) {
	url := "dispute/export"
	export := &Export{}
	err := s.client.Call("GET", url, options, export)
	return export, err
}
