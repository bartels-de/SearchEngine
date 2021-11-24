package indexing

import (
	"SearchEngine/analysis"
	"SearchEngine/core"
)

type TermIncidences struct {
Incidences map[core.DocumentUID]DocumentTermIncidences
}

type DocumentTermIncidences struct {
	Incidences []int
}

type InvertedIndex struct {
	InvertedIndex map[string]TermIncidences
	DocumentStore map[core.DocumentUID]core.Document
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		InvertedIndex: map[string]TermIncidences{},
		DocumentStore: map[core.DocumentUID]core.Document{},
	}
}

func (i *InvertedIndex) Index(terms []analysis.Term, documentUID core.DocumentUID) {
	i.addTermsToIndex(terms, documentUID)
}

func (i *InvertedIndex) Search(terms []analysis.Term) []core.DocumentUID {
	return i.query(terms)
}

func (i *InvertedIndex) addTermsToIndex(terms []analysis.Term, documentUID core.DocumentUID) {
	for _, term := range terms {
		i.addTermToIndex(term, documentUID)
	}
}

func (i *InvertedIndex) addTermToIndex(term analysis.Term, documentUID core.DocumentUID) {
	if _, ok := i.InvertedIndex[term.Term]; !ok {
		i.InvertedIndex[term.Term] = TermIncidences{
			Incidences: map[core.DocumentUID]DocumentTermIncidences{},
		}
	}

	if _, ok := i.InvertedIndex[term.Term].Incidences[documentUID]; !ok {
		i.InvertedIndex[term.Term].Incidences[documentUID] = DocumentTermIncidences{
			Incidences: []int{term.Position},
		}
	} else {
		i.InvertedIndex[term.Term].Incidences[documentUID] = DocumentTermIncidences{
			Incidences: append(
				i.InvertedIndex[term.Term].Incidences[documentUID].Incidences,
				term.Position,
			),
		}
	}
}

func (i *InvertedIndex) query(terms []analysis.Term) []core.DocumentUID {
	resultingDocumentsHash := map[core.DocumentUID]bool{}
	resultingDocuments := make([]core.DocumentUID, 0, 5)

	for _, term := range terms {
		if termIncidences, ok := i.InvertedIndex[term.Term]; ok {
			for documentUID := range termIncidences.Incidences {
				resultingDocumentsHash[documentUID] = true
			}
		}
	}

	for documentUID := range resultingDocumentsHash {
		resultingDocuments = append(resultingDocuments, documentUID)
	}

	return resultingDocuments
}