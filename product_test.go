package paystack

import (
	"strconv"
	"testing"
)

func TestProductCRUD(t *testing.T) {

	productRequest := &ProductRequest{
		Name:        "Puff Puff",
		Description: "Crispy flour ball with fluffy interior",
		Price:       "5000",
		Currency:    "NGN",
		Unlimited:   false,
		Quantity:    100,
	}

	// Test CREATE
	product, err := c.Product.Create(productRequest)
	if err != nil {
		t.Errorf("CREATE Product returned error: %v", err)
	}

	if product.ProductCode == "" {
		t.Errorf("Expected Product Code to be set")
	}

	if product.Name != productRequest.Name {
		t.Errorf("Expected Product name to be %v, got %v", productRequest.Name, product.Name)
	}

	// Test FETCH
	sameProduct, err := c.Product.Get(product.Id)
	if err != nil {
		t.Errorf("GET Product returned error: %v", err)
	}

	if sameProduct.Name != product.Name {
		t.Errorf("Expected Product Name to be %v, got %v", product.Name, sameProduct.Name)
	}

	if sameProduct.ProductCode != product.ProductCode {
		t.Errorf("Expected Product Code to be %v, got %v", product.ProductCode, sameProduct.ProductCode)
	}

	// retrieve the Product list
	products, err := c.Product.List()
	if err != nil || !(len(products.Values) > 0) || !(products.Meta.Total > 0) {
		t.Errorf("Expected Product list, got %d, returned error %v", len(products.Values), err)
	}

	// Test UPDATE Product
	updateRequest := &ProductRequest{
		Name:        "Puff Puff",
		Description: "Crispy flour ball with fluffy interior",
		Price:       "7000",
		Currency:    "NGN",
		Unlimited:   false,
		Quantity:    170,
	}
	updatedProduct, err := c.Product.Update(product.Id, updateRequest)
	if err != nil {
		t.Errorf("Failed to UPDATE Product: %v", err)
	}
	if updatedProduct.Quantity != updateRequest.Quantity {
		t.Errorf("Expected Product Quantity to be updated to %v, got %v", updatedProduct.Quantity, updateRequest.Quantity)
	}
	if strconv.Itoa(updatedProduct.Price) != updateRequest.Price {
		t.Errorf("Expected Product Quantity to be updated to %v, got %v", updatedProduct.Quantity, product.Quantity)
	}
}
