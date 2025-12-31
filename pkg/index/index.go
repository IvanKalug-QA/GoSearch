package index

import (
	"strings"
)

type Document struct {
	ID    int
	URL   string
	Title string
}

func CreateIndex(d []Document) map[string]map[int]struct{} {
	IndexDocuments := make(map[string]map[int]struct{}, 0)

	for _, doc := range d {
		splitText := strings.SplitSeq(doc.Title, " ")
		for t := range splitText {
			word := strings.ToLower(strings.TrimSpace(t))
			if IndexDocuments[word] == nil {
				IndexDocuments[word] = make(map[int]struct{})
			}
			IndexDocuments[word][doc.ID] = struct{}{}
		}
	}
	return IndexDocuments
}
