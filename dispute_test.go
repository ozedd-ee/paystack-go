package paystack

import "testing"

func TestDisputeService(t *testing.T) {
	// retrieve the dispute list
	options := &DisputeFilterOptions{}
	disputes, err := c.Dispute.List(options)
	if err != nil {
		t.Errorf("Error occurred while retrieving disputes: %v", err)
	}

	if !(len(disputes.Values) > 0) || !(disputes.Meta.Total > 0) {
		t.Skip("You currently have no disputes on your integration")
	}

	// fetch the split
	dispute1, err := c.Dispute.Get(disputes.Values[0].Id)
	if err != nil {
		t.Errorf("GET Dispute returned error: %v", err)
	}

	if dispute1.TransactionReference == "" {
		t.Error("Expected Dispute Transaction Reference to be set")
	}

	// list transaction disputes
	_, err = c.Dispute.ListTransactionDisputes(dispute1.Transaction.ID)
	if err != nil {
		t.Errorf("Failed to GET Dispute by transaction ID: %v", err)
	}

	// Test UPDATE Dispute
	newRefundAmount := 500000
	update := &UpdateDisputeRequest{
		RefundAmount: newRefundAmount,
	}
	updatedDispute, err := c.Dispute.Update(dispute1.Id, update)
	if err != nil {
		t.Errorf("Failed to UPDATE Dispute: %v", err)
	}
	if updatedDispute.RefundAmount != newRefundAmount {
		t.Errorf("Expected updated refund amount to be %v, got %v", newRefundAmount, updatedDispute.RefundAmount)
	}

	// Test AddDisputeEvidence
	evidence := &AddDisputeEvidenceRequest{
		CustomerEmail: "cus@gmail.com",
		CustomerName: "Mensah King",
		CustomerPhone: "0802345167",
		ServiceDetails: "claim for buying product",
		DeliveryAddress: "3a ladoke street ogbomoso",
    }
	disputeEvidence, err := c.Dispute.AddDisputeEvidence(dispute1.Id, evidence)
	if err != nil {
		t.Errorf("Unable to add Dispute evidence: %v", err)
	}
	if disputeEvidence.CustomerName == "" {
		t.Error("Expected Customer Name for dispute evidence to be set")
	}

	// Test GET UPLOAD URL
	upload, err := c.Dispute.GetUploadURL(dispute1.Id, "receipt.pdf")
	if err != nil {
		t.Errorf("Unable to get upload URL: %v", err)
	}
	if upload.SignedUrl == "" {
		t.Error("Expected Signed URL to be set")
	}

	// Test Export
	export, err := c.Dispute.Export(options)
	if err != nil {
		t.Errorf("Unable to export disputes: %v", err)
	}
	if export.Path == "" {
		t.Error("Expected export path to be set")
	}

	// Test ResolveDispute
	request := &ResolveDisputeRequest{
		Resolution: "merchant-accepted",
        Message: "Merchant accepted", 
        UploadedFilename: "qesp8a4df1xejihd9x5q", 
        RefundAmount: 300000, 
	}
	resolvedDispute, err := c.Dispute.ResolveDispute(dispute1.Id, request)
	if err != nil {
		t.Errorf("Unable to resolve dispute: %v", err)
	}
	if resolvedDispute.Resolution != request.Resolution {
		t.Errorf("Expected dispute resolution to be %v, got %v", request.Resolution, resolvedDispute.Resolution)
	}
}
