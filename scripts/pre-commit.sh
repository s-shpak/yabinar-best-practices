#!/bin/sh
set -e

./scripts/go-mod-pre-commit.sh
./scripts/golangci-lint-pre-commit.sh
