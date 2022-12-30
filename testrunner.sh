#!/usr/bin/env bash

go test
go test -bench=. -benchmem
