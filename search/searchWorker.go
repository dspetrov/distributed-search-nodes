package search

import (
	"bufio"
	"dspetrov/distributed-search/model"
	"fmt"
	"os"

	proto "github.com/gogo/protobuf/proto"
)

const TASK_ENDPOINT = "/task"

type SearchWorker struct {
}

func (s SearchWorker) HandleRequest(requestPayload []byte) []byte {
	var task model.Task
	if err := proto.Unmarshal(requestPayload, &task); err != nil {
		panic(err)
	}

	result := s.createResult(task)

	resultBytes, err := proto.Marshal(&result)
	if err != nil {
		panic(err)
	}

	return resultBytes
}

func (s SearchWorker) createResult(task model.Task) model.Result {
	documents := task.GetDocuments()
	fmt.Println("Received %v documents to process", len(documents))

	result := model.Result{}

	for _, document := range documents {
		words := s.parseWordsFromDocument(document)
		documentData := createDocumentData(words, task.GetSearchTerms())
		result.DocumentToDocumentData[document] = documentData
	}

	return result
}

func (s SearchWorker) parseWordsFromDocument(document string) []string {
	lines := s.getLinesFromDocument(document)
	words := getWordsFromDocument(lines)
	return words
}

func (s SearchWorker) getLinesFromDocument(documentPath string) []string {
	lines := []string{}

	file, err := os.Open(documentPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func (s SearchWorker) GetEndpoint() string {
	return TASK_ENDPOINT
}
