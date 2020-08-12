terraform {
  required_providers {
    hello = {
      # this source assumes you have installed the 
      # provider locally as specified in the README
      # Change this if you are pulling from
      # the Terraform registry
      source  = "example.com/example/hello"
      version = "<version of provider (ex: 0.3.0)>"
    }
  }
  required_version = ">= 0.13"
}
