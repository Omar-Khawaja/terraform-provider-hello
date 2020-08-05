#!/usr/bin/env bash

go build -o ../terraform-provider-hello_$(git describe --tags --abbrev=0) ..
