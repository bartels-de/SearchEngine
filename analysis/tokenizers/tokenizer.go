package tokenizers

import "SearchEngine/analysis"

type Tokenizer func(input string) []analysis.Term
