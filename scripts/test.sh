#!/usr/bin/env bash
go test ./src/presentation/.. -cover
go test ./src/data/.. -cover
go test ./src/domain/usecase/.. -cover