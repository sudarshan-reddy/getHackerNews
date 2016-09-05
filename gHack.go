package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func errHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type story struct {
	Title string
}

func getTitles(address string) string {
	var stories story
	resp, err := http.Get(address)
	errHandle(err)
	body, err := ioutil.ReadAll(resp.Body)
	errHandle(err)
	Jserr := json.Unmarshal(body, &stories)
	errHandle(Jserr)
	return stories.Title
}

func main() {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	errHandle(err)
	body, err := ioutil.ReadAll(resp.Body)
	errHandle(err)
	apis := strings.Split(string(body), ",")
	i, err := strconv.Atoi(os.Args[1])
	errHandle(err)
	getString := "https://hacker-news.firebaseio.com/v0/item/" +
		strings.Trim(apis[i], "[ ") +
		".json?print=json"
	fmt.Println(getTitles(getString))
}
