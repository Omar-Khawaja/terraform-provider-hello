## About

Demo Terraform provider for
[hello-service](https://github.com/Omar-Khawaja/hello-service)

## Build and Installation

Run `make install` to build the plugin and install it in the necessary directory
for Terraform to detect it.

If you want to just build the plugin and move it to a different directory
yourself, run `make build` to build the plugin.

## Usage

You will need to run `terraform init` before using the provider. See the
[examples](./examples) directory for sample Terraform config.