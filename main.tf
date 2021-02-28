provider "product" {
  version = "1.0.0"
  address = "http://localhost"
  port    = "18010"
}

resource "product" "test" {
  id = 1
  name = "TestProduct"
  price = 11.00
}