package search

import (
	"GoSearch/pkg/index"
	"fmt"
	"sort"
	"strings"
)

type DocScore struct {
	ID    int
	Count int
}

func Search(s string, indexDoc map[string][]int, d []index.Document) []string {
	words := strings.Fields(strings.ToLower(s))
	docId := make(map[int]int)
	documentOutput := make([]string, 0, 3)
	for _, w := range words {
		if dId, ok := indexDoc[w]; ok {
			for _, d := range dId {
				docId[d]++
			}
		}
	}
	docScore := make([]DocScore, 0, len(docId))
	for id, count := range docId {
		docScore = append(docScore, DocScore{ID: id, Count: count})
	}
	sort.Slice(docScore, func(i, j int) bool { return docScore[i].Count > docScore[j].Count })
	sort.Slice(d, func(i, j int) bool { return d[i].ID < d[j].ID })
	for _, k := range docScore {
		doc := sort.Search(len(d), func(i int) bool { return d[i].ID >= k.ID })
		if doc < len(d) && d[doc].ID == k.ID {
			documentOutput = append(documentOutput, fmt.Sprintf("TITLE: %v, URL: %v\n", d[doc].Title, d[doc].URL))
		}
	}
	return documentOutput
}
