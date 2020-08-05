provider "hello" {}

resource "hello_ping" "example_ping" {
  name  = "my-ping"
  count = 2
}

resource "hello_pong" "example_pong" {
  name = "my-pong"
}

resource "hello_append_item" "example" {
  name  = "adding-apple"
  item  = "apple"
  count = 2
}
