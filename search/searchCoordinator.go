package search

import (
	"dspetrov/distributed-search-nodes/clusterManagement"
	"dspetrov/distributed-search-nodes/model"
	"dspetrov/distributed-search-nodes/networking"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/emirpasic/gods/maps/treemap"
	proto "github.com/gogo/protobuf/proto"
)

const (
	SEARCH_ENDPOINT = "/search"
	BOOKS_DIRECTORY = "./resources/books/"
)

type SearchCoordinator struct {
	workersServiceRegistry *clusterManagement.ServiceRegistry
	client                 *networking.WebClient
	documents              []string
}

func NewSearchCoordinator(workersServiceRegistry *clusterManagement.ServiceRegistry, client *networking.WebClient) *SearchCoordinator {
	sc := SearchCoordinator{
		workersServiceRegistry: workersServiceRegistry,
		client:                 client,
		documents:              readDocumentsList(),
	}

	return &sc
}

func (sc SearchCoordinator) HandleRequest(requestPayload []byte) []byte {
	var request model.Request
	if err := proto.Unmarshal(requestPayload, &request); err != nil {
		fmt.Println(err)
		return []byte{}
	}

	response := sc.createResponse(request)

	responseBytes, err := proto.Marshal(&response)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}

	return responseBytes
}

func (sc SearchCoordinator) GetEndpoint() string {
	return SEARCH_ENDPOINT
}

func (sc SearchCoordinator) createResponse(searchRequest model.Request) model.Response {
	fmt.Println("Received search query:", searchRequest.SearchQuery)

	searchTerms := getWordsFromLine(searchRequest.SearchQuery)

	workers := sc.workersServiceRegistry.GetAllServiceAddresses()

	if len(workers) == 0 {
		fmt.Println("No search workers currently available")
		return model.Response{}
	}

	tasks := sc.createTasks(len(workers), searchTerms)
	results := sc.sendTasksToWorkers(workers, tasks)

	sortedDocuments := sc.aggregateResults(results, searchTerms)
	searchResponse := model.Response{
		RelevantDocuments: sortedDocuments,
	}

	return searchResponse
}

func (sc SearchCoordinator) aggregateResults(results []model.Result, terms []string) []*model.Response_DocumentStats {
	allDocumentsResults := make(map[string]*model.DocumentData)

	for _, result := range results {
		for doc, docData := range result.GetDocumentToDocumentData() {
			allDocumentsResults[doc] = docData
		}
	}

	fmt.Println("Calculating score for all the documents")
	scoreToDocuments := GetDocumentsScores(terms, allDocumentsResults)

	return sc.sortDocumentsByScore(scoreToDocuments)
}

func (sc SearchCoordinator) sortDocumentsByScore(scoreToDocuments *treemap.Map) []*model.Response_DocumentStats {
	var sortedDocumentsStatsList []*model.Response_DocumentStats

	for _, key := range scoreToDocuments.Keys() {
		score := key.(float64)
		documents, _ := scoreToDocuments.Get(score)
		for _, document := range documents.([]string) {
			fi, err := os.Stat(document)
			if err != nil {
				panic(err)
			}

			documentStats := model.Response_DocumentStats{
				Score:        score,
				DocumentName: fi.Name(),
				DocumentSize: fi.Size(),
			}

			sortedDocumentsStatsList = append(sortedDocumentsStatsList, &documentStats)
		}
	}

	return sortedDocumentsStatsList
}

func (sc SearchCoordinator) sendTasksToWorkers(workers []string, tasks []model.Task) []model.Result {
	workersCount := len(workers)
	ch := make(chan model.Result)

	for i := 0; i < workersCount; i++ {
		payload, err := proto.Marshal(&tasks[i])
		if err != nil {
			panic(err)
		}

		go sc.client.SendTask(workers[i], payload, ch)
	}

	results := make([]model.Result, workersCount)
	for i := 0; i < workersCount; i++ {
		results[i] = <-ch
	}

	fmt.Printf("Received %v results\n", workersCount)

	return results
}

func (sc SearchCoordinator) createTasks(numberOfWorkers int, searchTerms []string) []model.Task {
	workersDocuments := splitDocumentList(numberOfWorkers, sc.documents)

	tasks := []model.Task{}

	for _, documentsForWorker := range workersDocuments {
		task := model.Task{
			SearchTerms: searchTerms,
			Documents:   documentsForWorker,
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func splitDocumentList(numberOfWorkers int, documents []string) [][]string {
	numberOfDocumentsPerWorker := (len(documents) + numberOfWorkers - 1) / numberOfWorkers

	workersDocuments := [][]string{}

	for i := 0; i < numberOfWorkers; i++ {
		firstDocumentIndex := i * numberOfDocumentsPerWorker
		lastDocumentIndexExclusive := min(firstDocumentIndex+numberOfDocumentsPerWorker, len(documents))

		if firstDocumentIndex >= lastDocumentIndexExclusive {
			break
		}

		currentWorkerDocuments := documents[firstDocumentIndex:lastDocumentIndexExclusive]

		workersDocuments = append(workersDocuments, currentWorkerDocuments)
	}

	return workersDocuments
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func readDocumentsList() []string {
	files, err := ioutil.ReadDir(BOOKS_DIRECTORY)
	if err != nil {
		panic(err)
	}

	var result []string
	for _, f := range files {
		result = append(result, BOOKS_DIRECTORY+"/"+f.Name())
	}

	return result
}
