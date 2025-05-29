package paystack

import "testing"

func TestDedicatedVirtualAccount(t *testing.T) {
	cust := &Customer{
		FirstName: "User123",
		LastName:  "AdminUser",
		Email:     "user1-deny@gmail.com",
		Phone:     "+2341000000000000",
	}
	customer1, _ := c.Customer.Create(cust)
	
	// Test CREATE
	dvaRequest := &DedicatedVirtualAccountRequest{
		Customer:      customer1.ID,
		PreferredBank: "test-bank",
	}

	dva, err := c.DedicatedVirtualAccount.Create(dvaRequest)
	if err != nil {
		t.Errorf("CREATE Dedicated Virtual Account returned error: %v", err)
	}

	if dva.AccountName == "" {
		t.Errorf("Expected account name to be set")
	}

	if dva.Bank.Name == "" {
		t.Errorf("Expected Bank name to be set")
	}

	// test ASSIGN
	assignDVARequest := &AssignDVARequest{
		Email:         "janedoe@test.com",
		FirstName:     "Jane",
		MiddleName:    "Karen",
		LastName:      "Doe",
		Phone:         "+2348100000000",
		PreferredBank: "test-bank",
		Country:       "NG",
	}

	dva1, err := c.DedicatedVirtualAccount.Assign(assignDVARequest)
	if err != nil {
		t.Errorf("ASSIGN Dedicated Virtual Account returned error: %v", err)
	}

	if dva1.AccountNumber == "" {
		t.Errorf("Expected account name to be set")
	}

	if dva1.Bank.Name == "" {
		t.Errorf("Expected Bank name to be set")
	}

	// Test LIST DVA
	filter := &DVAListFilter{}
	dvaList, err := c.DedicatedVirtualAccount.List(filter)
	if err != nil || !(len(dvaList.Values) > 0) || !(dvaList.Meta.Total > 0) {
		t.Errorf("Expected DVA list, got %d, returned error %v", len(dvaList.Values), err)
	}

	if dvaList.Values[0].AccountName == "" {
		t.Errorf("Expected Account name for first DVA in List to be set")
	}

	// Test FETCH
	sameDVA, err := c.DedicatedVirtualAccount.Get(dva1.Id)
	if err != nil {
		t.Errorf("GET DVA returned error: %v", err)
	}

	if sameDVA.AccountName != dva1.AccountName {
		t.Errorf("Expected Account Name to be %v, got %v", dva1.AccountName, sameDVA.AccountName)
	}

	// Test REQUERY
	req := &RequeryDVARequest{
		AccountNumber: "1234567890",
		ProviderSlug:  "example-provider",
		Date:          "2023-05-30",
	}
	_, err = c.DedicatedVirtualAccount.Requery(req)
		if err != nil {
		t.Errorf("REQUERY DVA returned error: %v", err)
	}

	// Test SPLIT
	splitRequest := &DVATransactionSplitRequest{
		Customer: 481193,
		PreferredBank: "wema-bank",
		SplitCode: "SPL_e7jnRLtzla",
	}
	dva2, err := c.DedicatedVirtualAccount.Split(splitRequest)
	if err != nil {
		t.Errorf("Failed to add Split to DVA: %v", err)
	}
	if dva2.SplitConfig.SplitCode == "" {
		t.Errorf("Expected Split Code to be set")
	}

	// Test REMOVE SPLIT
	noSplitDva2, err := c.DedicatedVirtualAccount.RemoveSplit(dva2.AccountNumber)
	if err != nil {
		t.Errorf("Failed to Remove Split from DVA: %v", err)
	}
	if noSplitDva2.SplitConfig.SplitCode != "" {
		t.Errorf("Expected Split Code to be removed")
	}


	// Test DEACTIVATE
	deactivatedDVA, err := c.DedicatedVirtualAccount.Deactivate(dva2.Id)
	if err != nil {
		t.Errorf("Failed to DEACTIVATE DVA: %v", err)
	}
	if deactivatedDVA.Assigned != false {
		t.Errorf("Expected DVA to be unassigned")
	}

	// Test DELETE
	providers, err := c.DedicatedVirtualAccount.GetBankProviders()
	if err != nil {
		t.Errorf("Failed to FETCH Bank providers: %v", err)
	}
	if providers[0].BankName == "" {
		t.Errorf("Expected Bank Name to be set for provider")
	}
}
