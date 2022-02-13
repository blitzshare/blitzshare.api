//go:build tools
// +build tools

package tools

import (
	_ "github.com/cucumber/godog/cmd/godog"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
	_ "github.com/swaggo/swag/cmd/swag"
	_ "github.com/vektra/mockery/v2"
	_ "golang.org/x/tools/cmd/goimports"
)
