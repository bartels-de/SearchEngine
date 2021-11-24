package indexing

import "SearchEngine/core"

type DocumentStore struct {
	StoreMap map[core.DocumentUID]core.Document
}

func NewDocumentStore() *DocumentStore {
	return &DocumentStore{
		StoreMap: map[core.DocumentUID]core.Document{},
	}
}

func (i *DocumentStore) Store(UID core.DocumentUID, document core.Document) {
	i.StoreMap[UID] = document
}

func (i *DocumentStore) Fetch(documentUID core.DocumentUID) core.Document {
	return i.StoreMap[documentUID]
}

func (i *DocumentStore) FetchDocuments(documentUIDS []core.DocumentUID) []core.Document {
	resultingDocuments := make([]core.Document, 0, len(documentUIDS))

	for _, documentUID := range documentUIDS {
		resultingDocuments = append(resultingDocuments, i.Fetch(documentUID))
	}

	return resultingDocuments
}