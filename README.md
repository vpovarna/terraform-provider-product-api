### Terraform provider
Terraform provider for: https://github.com/vpovarna/go-mux-api

### Install dependencies
To install dependencies, run:
```
$ go get -u github.com/vpovarna/go-mux-api
$ go get -u github.com/hashicorp/terraform/plugin
$ go get -u github.com/hashicorp/terraform/terraform
$ go get -u github.com/hashicorp/terraform/helper/schema
```

### Build the project
```
$ go build -v -o terraform-provider-product_v1.0.0
````
Copy the output file to terraform plugin folder.

```
$ cp terraform-provider-product_v1.0.0 ~/.terraform.d/plugins/
$ chmod +x ~/.terraform.d/plugins/terraform-provider-product_v1.0.0
```
