package main

import (
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"
)

func main() {
	_, err := maxprocs.Set()
	if err == nil {
		slog.Error("Failed to set process limit", err)
	}
}
