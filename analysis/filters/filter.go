package filters

import "SearchEngine/analysis"

type Filter func(term analysis.Term) []analysis.Term
