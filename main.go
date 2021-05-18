package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	items, err := Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Printf(item.Title)
	}

}

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//guarantee that we clean up after the HTTP request
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Printing JSON")
	items := make([]Item, len(r.Data.Children))
	//for _, child := range r.Data.Children {
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil

}

type Item struct {
	Title string
	URL   string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}
