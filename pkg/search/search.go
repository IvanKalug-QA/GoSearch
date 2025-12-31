package search

import (
	"GoSearch/pkg/index"
	"sort"
	"strings"
)

func Search(s string, indexDoc map[string]map[int]struct{}, d []index.Document) []index.Document {
	text := strings.SplitSeq(s, " ")
	documentsId := map[int]struct{}{}
	documentOutput := make([]index.Document, 0, 3)
	for t := range text {
		lower := strings.ToLower(t)
		docId, ok := indexDoc[lower]
		if ok {
			for k := range docId {
				_, has := documentsId[k]
				if !has {
					documentsId[k] = struct{}{}
				}
			}
		}
	}
	sort.Slice(d, func(i, j int) bool { return d[i].ID < d[j].ID })
	for k := range documentsId {
		doc := sort.Search(len(d), func(i int) bool { return d[i].ID >= k })
		if doc < len(d) && d[doc].ID == k {
			documentOutput = append(documentOutput, d[doc])
		}
	}
	return documentOutput
}
