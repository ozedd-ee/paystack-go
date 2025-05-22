package paystack

import (
	"fmt"
	"testing"
)

func TestRefund(t *testing.T) {
	txn := &TransactionRequest{
		Email:     "user123@gmail.com",
		Amount:    600000,
		Reference: "Txn-" + fmt.Sprintf("%d", makeTimestamp()),
	}
	resp, err := c.Transaction.Initialize(txn)
	if err != nil {
		t.Error(err)
	}

	if resp["reference"] == "" {
		t.Error("Missing transaction reference")
	}

	txn1, err := c.Transaction.Verify(resp["reference"].(string))

	if err != nil {
		t.Error(err)
	}

	if txn1.Reference == "" {
		t.Errorf("Missing transaction reference")
	}

	request := &RefundRequest{
		Transaction: txn1.Reference,
	}

	refund, err := c.Refund.CreateRefund(request)
	if err != nil {
		t.Errorf("CREATE Refund returned error: %v", err)
	}

	if refund.Id == 0 {
		t.Errorf("Expected Refund ID to be set")
	}

	if refund.Amount != int(txn1.Amount) {
		t.Errorf("Expected refund amount to be %v, got %v", txn1.Amount, refund.Amount)
	}

	sameRefund, err := c.Refund.Get(refund.Id)
	if err != nil {
		t.Errorf("GET Refund returned error: %v", err)
	}

	if sameRefund.Id != refund.Id {
		t.Errorf("Expected Refund Id to be %v, got %v", refund.Id, sameRefund.Id)
	}

	refunds, err := c.Refund.List()
	if err != nil || !(len(refunds.Values) > 0) || !(refunds.Meta.Total > 0) {
		t.Errorf("Expected refund list, got %d, returned error %v", len(refunds.Values), err)
	}
}
