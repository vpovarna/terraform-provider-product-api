package provider

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vpovarna/go-mux-api/client"
	"github.com/vpovarna/go-mux-api/server"
)

func resourceProduct() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "An unique ID for the product",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the item",
				ValidateFunc: validateName,
			},
			"price": {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Product price",
			},
		},
		Create: resourceCreateProduct,
		Read:   resourceReadProduct,
		Update: resourceUpdateProduct,
		Delete: resourceDeleteProduct,
		Exists: resourceExistsProduct,
	}
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string

	value, ok := v.(string)

	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs

}

func resourceCreateProduct(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	product := server.Product{
		ID:    d.Get("id").(int),
		Name:  d.Get("name").(string),
		Price: d.Get("price").(float64),
	}

	err := apiClient.NewProduct(&product)

	if err != nil {
		return err
	}
	id := strconv.Itoa(product.ID)
	d.SetId(id)
	return nil
}

func resourceReadProduct(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	productID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	product, err := apiClient.GetProduct(productID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding product with ID %d", productID)
		}
	}

	id := strconv.Itoa(product.ID)
	d.SetId(id)
	d.Set("name", product.Name)
	d.Set("price", product.Price)
	return nil
}

func resourceUpdateProduct(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	product := server.Product{
		ID:    d.Get("id").(int),
		Name:  d.Get("name").(string),
		Price: d.Get("price").(float64),
	}

	err := apiClient.UpdateProduct(&product)
	if err != nil {
		return err
	}

	return nil
}

func resourceDeleteProduct(d *schema.ResourceData, m interface{}) error {
	apiClinet := m.(*client.Client)

	id := d.Id()

	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = apiClinet.DeleteProduct(productID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsProduct(d *schema.ResourceData, m interface{}) (bool, error) {

	apiClient := m.(*client.Client)
	id := d.Id()
	productID, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}

	_, err = apiClient.GetProduct(productID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
