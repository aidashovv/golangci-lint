package loglint

import (
	loglintpkg "github.com/aidashovv/loglint/pkg/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = loglintpkg.Analyzer

func New() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"loglint",
		"linter for checking logs",
		[]*analysis.Analyzer{Analyzer},
		nil,
	)
}
