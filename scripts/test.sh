#!/usr/bin/env bash

./scripts/mocks.sh

go test ./src/presentation/.. -cover
go test ./src/data/.. -cover
go test ./src/domain/usecase/.. -cover