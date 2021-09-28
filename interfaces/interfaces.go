package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	inmemorydataprovider := InMemoryDataProvider{
		Data: map[int]string{
			1: "item1",
			2: "item2",
		},
	}
	apidataprovider := APIDataProvider{"https://jsonplaceholder.typicode.com/posts/"}

	var dataProvider IDataProvider

	dataProvider = inmemorydataprovider
	fmt.Println("Response from In Memory data provider:", dataProvider.GetItem(1))

	dataProvider = apidataprovider
	fmt.Println("Response from API data provider:", dataProvider.GetItem(1))
}

type IDataProvider interface {
	GetItem(id int) string
}

type InMemoryDataProvider struct {
	Data map[int]string
}

func (dp InMemoryDataProvider) GetItem(id int) string {
	return dp.Data[id]
}

type APIDataProvider struct {
	apiurl string
}

type post struct {
	Id    int    `json:"id"` //the field need to be exported, aka Capital letter first
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (dp APIDataProvider) GetItem(id int) string {
	//get request
	resp, getErr := http.Get(fmt.Sprintf("%s%d", dp.apiurl, id))
	if getErr != nil {
		log.Fatalln(getErr)
	}

	//read body
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	p := post{}
	if jsonErr := json.Unmarshal(body, &p); jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	//close body
	resp.Body.Close()

	return p.Title
}
