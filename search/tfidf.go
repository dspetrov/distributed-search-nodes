package search

import (
	"dspetrov/distributed-search-nodes/model"
	"math"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
)

func calculateTermFrequency(words []string, term string) float64 {
	count := 0
	for _, word := range words {
		if strings.EqualFold(term, word) {
			count++
		}
	}

	termFrequency := float64(count) / float64(len(words))
	return termFrequency
}

func createDocumentData(words []string, terms []string) *model.DocumentData {
	dd := model.DocumentData{
		TermToFrequency: make(map[string]float64),
	}

	for _, term := range terms {
		termFreq := calculateTermFrequency(words, term)
		dd.TermToFrequency[term] = termFreq
	}

	return &dd
}

func GetDocumentsScores(terms []string, documentResults map[string]*model.DocumentData) *treemap.Map {
	scoreToDocuments := treemap.NewWith(func(a, b interface{}) int { return -1 * utils.Float64Comparator(a, b) })

	termToInverseDocumentFrequency := getTermToInverseDocumentFrequencyMap(terms, documentResults)

	for document, documentData := range documentResults {
		score := calculateDocumentScore(terms, documentData, termToInverseDocumentFrequency)
		addDocumentScoreToTreeMap(scoreToDocuments, score, document)
	}

	return scoreToDocuments
}

func addDocumentScoreToTreeMap(scoreToDoc *treemap.Map, score float64, document string) {
	var documentsWithCurrentScore []string
	scoreToDocValue, valueFound := scoreToDoc.Get(score)
	if valueFound {
		documentsWithCurrentScore = scoreToDocValue.([]string)
	}

	documentsWithCurrentScore = append(documentsWithCurrentScore, document)
	scoreToDoc.Put(score, documentsWithCurrentScore)
}

func calculateDocumentScore(terms []string, documentData *model.DocumentData, termToInverseDocumentFrequency map[string]float64) float64 {
	score := 0.0
	for _, term := range terms {
		termFrequency := documentData.TermToFrequency[term]
		inverseTermFrequency := termToInverseDocumentFrequency[term]
		score += termFrequency * inverseTermFrequency
	}

	return score
}

func getInverseDocumentFrequency(term string, documentResults map[string]*model.DocumentData) float64 {
	nt := 0
	for _, documentData := range documentResults {
		termFrequency := documentData.TermToFrequency[term]
		if termFrequency > 0 {
			nt++
		}
	}

	if nt == 0 {
		return 0
	}

	return math.Log10(float64(len(documentResults)) / float64(nt))
}

func getTermToInverseDocumentFrequencyMap(terms []string, documentResults map[string]*model.DocumentData) map[string]float64 {
	termToIDF := make(map[string]float64)
	for _, term := range terms {
		idf := getInverseDocumentFrequency(term, documentResults)
		termToIDF[term] = idf
	}

	return termToIDF
}

func getWordsFromDocument(lines []string) []string {
	var words []string
	for _, line := range lines {
		words = append(words, getWordsFromLine(line)...)
	}

	return words
}

func getWordsFromLine(line string) []string {
	r, _ := regexp.Compile("(\\.)+|(,)+|( )+|(-)+|(\\?)+|(!)+|(;)+(:)+|(/d)+|(/n)+")
	return r.Split(line, -1)
}
