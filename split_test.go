package paystack

import "testing"

func TestSplitCRUD(t *testing.T) {

	subAccount := &SubAccount{
		BusinessName:     "Sunshine Studios",
		SettlementBank:   "057",
		AccountNumber:    "0000000000",
		PercentageCharge: 12.8,
	}

	// create the subAccount
	subAccount, err := c.SubAccount.Create(subAccount)
	if err != nil {
		t.Errorf("CREATE SubAccount returned error: %v", err)
	}

	subaccount := BeneficiaryAccountRequest{
		SubAccountCode: subAccount.SubAccountCode,
		Share:          20,
	}
	splitRequest := &SplitRequest{
		Name:        "Halfsies",
		Type:        "percentage",
		Currency:    "NGN",
		Subaccounts: []BeneficiaryAccountRequest{subaccount},
	}

	split, err := c.Split.CreateSplit(splitRequest)
	if err != nil {
		t.Errorf("CreateSplit returned error: %v", err)
	}

	if split.SplitCode == "" {
		t.Errorf("Expected SplitCode to be set")
	}

	if split.Name != "Halfsies" {
		t.Errorf("Expected Split name to be %v, got %v", splitRequest.Name, split.Name)
	}

	// fetch the split
	sameSplit, err := c.Split.Get(split.SplitID)
	if err != nil {
		t.Errorf("GET Spilt returned error: %v", err)
	}

	if sameSplit.Name != split.Name {
		t.Errorf("Expected Split Name to be %v, got %v", split.Name, sameSplit.Name)
	}

	// retrieve the Split list
	splits, err := c.Split.List()
	if err != nil || !(len(splits.Values) > 0) || !(splits.Meta.Total > 0) {
		t.Errorf("Expected Split list, got %d, returned error %v", len(splits.Values), err)
	}

	// Test UPDATE Split
	update := &SplitUpdateRequest{
		Name: "Royalty",
		Active: true,
	}
	updatedSplit, err := c.Split.Update(sameSplit.SplitID, update)
	if err != nil {
		t.Errorf("Failed to UPDATE Split: %v", err)
	}
	if updatedSplit.Name != update.Name {
		t.Errorf("Expected Split Name to be updated to %v, got %v", update.Name, updatedSplit.Name)
	}

	// Test UPDATE Split SubAccounts
	newShare := 50
	updatedSplit, err = c.Split.UpdateSubAccounts(split.SplitID, subAccount.SubAccountCode, newShare)
	if err != nil {
		t.Errorf("Failed to UPDATE Split SubAccounts: %v", err)
	}
	if updatedSplit.Subaccounts[0].Share != newShare {
		t.Errorf("Expected Split SubAccount share to be updated to %v, got %v", newShare, updatedSplit.Subaccounts[0].Share)
	}

	// Test DELETE
	err = c.Split.RemoveSubAccount(split.SplitID, subAccount.SubAccountCode)
	if err != nil {
		t.Errorf("Failed to REMOVE Split SubAccount: %v", err)
	}
}
