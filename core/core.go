package core

type DocumentUID uint32

type Document struct {
	UID        DocumentUID
	Attributes map[string]string
}