[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/rpip/paystack-go) [![Build Status](https://travis-ci.org/rpip/paystack-go.svg?branch=master)](https://travis-ci.org/rpip/paystack-go) 

# Go library for the Paystack API.

paystack-go is a Go client library for accessing the Paystack API.

Where possible, the services available on the client groups the API into logical chunks and correspond to the structure of the Paystack API documentation at https://developers.paystack.co/v1.0/reference.

## Usage
Reference paystack-go in your go program:
``` go
import "github.com/rpip/paystack-go"
```

Initialize new Paystack client:
``` go
apiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"

// The second parameter is an optional HTTP client, allowing overriding of the HTTP client to use. This is useful if you're running in a Google AppEngine environment where the http.DefaultClient is not available.
client := paystack.NewClient(apiKey)
```
### Transfers
Create a TransferRecipient:
``` go
transferRecipient := &TransferRecipient{
    Type:          "Nuban",
    Name:          "Customer 1",
    Description:   "Demo customer",
    AccountNumber: "0100000010",
    BankCode:      "044",
    Currency:      "NGN",
    Metadata:      map[string]interface{}{"job": "Plumber"},
}

recipient, err := client.Transfer.CreateRecipient(transferRecipient)
// You can store the RecipientCode(recipient.RecipientCode) and retrieve as desired for transfers
```
Initiate transfer:
``` go
transferRequest := &TransferRequest{
    Source:    "balance", // Funds to be transferred from your PayStack balance
    Amount:    30, // In least denomination (Kobo if NGN, pesewas if GHS)
    Recipient: recipient1.RecipientCode,
    Currency: "NGN" // Optional. Defaults to NGN
    Reason:    "Delivery pickup", // Optional

}

transfer, err := client.Transfer.Initiate(transferRequest)
if err != nil {
    // do something with error
}
```
### Customers
``` go
customer := &Customer{
    FirstName: "User123",
    LastName:  "AdminUser",
    Email:     "user123@gmail.com",
    Phone:     "+23400000000000000",
}
// create the customer
customer, err := client.Customer.Create(customer)
if err != nil {
    // do something with error
}

// Get customer by ID
customer, err := client.Customers.Get(customer.ID)
```

See the test files for more examples.

## Docker

Test this library in a docker container:

```bash
# PAYSTACK_KEY is an environment variable that should be added to your rc file. i.e .bashrc
$ make docker && docker run -e PAYSTACK_KEY -i -t paystack:latest
```

## TODO
- [ ] Maybe support request context?
- [ ] Test on App Engine

## CONTRIBUTING
Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. Please make sure the build succeeds and the test suite passes.
