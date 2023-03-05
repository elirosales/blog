#!/bin/bash

go test -v ./service -coverprofile=./service/coverage.out
go tool cover -func=./service/coverage.out

echo "\n"

go test -v ./database/models -coverprofile=./database/models/coverage.out
go tool cover -func=./database/models/coverage.out