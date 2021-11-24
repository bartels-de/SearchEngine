package filters

import (
	"SearchEngine/analysis"
	"strings"
)

func LowercaseFilter(term analysis.Term) []analysis.Term {
	return []analysis.Term{
		{Position: term.Position, Term: strings.ToLower(term.Term)},
	}
}
