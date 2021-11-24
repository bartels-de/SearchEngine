package tokenizers

import (
	"SearchEngine/analysis"
	"strings"
)

func SimpleTokenizer(input string) []analysis.Term{
	terms := strings.Split(input, " ")
	tokens := make([]analysis.Term, 0, 20)

	for position, term := range terms{
		tokens = append(tokens, analysis.Term{
			Position: position +1,
			Term: term,
		})
	}

	return tokens
}
