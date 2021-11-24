package analyzers

import (
	"SearchEngine/analysis/filters"
	"SearchEngine/analysis/tokenizers"
)

func NewSimpleAnalyzer() Analyzer {
	return Analyzer{
		Tokenizer: tokenizers.SimpleTokenizer,
		Filters: []filters.Filter{
			filters.LowercaseFilter,
		},
	}
}
