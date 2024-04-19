terraform {
  required_providers {
    example = {
      source = "registry.terraform.io/koki-develop/example"
    }
  }
}

provider "example" {}

output "hello" {
  value = provider::example::hello("koki")
}
