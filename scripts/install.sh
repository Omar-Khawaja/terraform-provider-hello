#!/usr/bin/env bash

go build -o terraform-provider-hello_$(git describe --tags) ..
mv terraform-provider-hello_$(git describe --tags) ~/.terraform.d/plugins/
