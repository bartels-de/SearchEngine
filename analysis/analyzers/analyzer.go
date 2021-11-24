package analyzers

import (
	"SearchEngine/analysis"
	"SearchEngine/analysis/filters"
	"SearchEngine/analysis/tokenizers"
)

type Analyzer struct {
	Tokenizer tokenizers.Tokenizer
	Filters   []filters.Filter
}

func (a *Analyzer) Analyze(input string) []analysis.Term {
	terms := a.Tokenizer(input)

	for _, filter := range a.Filters {
		terms = applyFilter(filter, terms)
	}

	return terms
}

func applyFilter(filter filters.Filter, terms []analysis.Term) []analysis.Term {
	filteredTerms := make([]analysis.Term, 0, 5)

	for _, term := range terms {
		filteredTerms = append(filteredTerms, filter(term)...)
	}

	return filteredTerms
}