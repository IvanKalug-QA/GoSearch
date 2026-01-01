package index

import (
	"strings"
)

type Document struct {
	ID    int
	URL   string
	Title string
}

func CreateIndex(d []Document) map[string][]int {
	has_doc := make(map[string]map[int]struct{})
	IndexDocuments := make(map[string][]int)

	for _, doc := range d {
		splitText := strings.SplitSeq(doc.Title, " ")
		for t := range splitText {
			word := strings.ToLower(strings.TrimSpace(t))
			if has_doc[word] == nil {
				IndexDocuments[word] = append(IndexDocuments[word], doc.ID)
				has_doc[word] = make(map[int]struct{})
			} else {
				_, ok := has_doc[word][doc.ID]
				if !ok {
					IndexDocuments[word] = append(IndexDocuments[word], doc.ID)
				}
			}
			has_doc[word][doc.ID] = struct{}{}
		}
	}
	return IndexDocuments
}
