package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type BookStruct struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type BooksRepo struct {
	Docs   []BookStruct `json:"docs"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
	Page   int          `json:"page"`
	Pages  int          `json:"pages"`
}

func main() {
	filename := flag.String("filename", "books", "File name to save the books")
	flag.Parse()
	var books BooksRepo
	resp, _ := http.Get("https://the-one-api.dev/v2/book")
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &books)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*filename)
	f, err := os.Create(*filename + ".csv")
	defer f.Close()

	if err != nil {

		log.Fatal(err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, book := range books.Docs {
		if err := w.Write([]string{book.Id, book.Name}); err != nil {
			log.Fatal(err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
