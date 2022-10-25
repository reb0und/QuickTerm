package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	os.Setenv("THESAURUS_API_KEY", "3794032c-4efe-4e3b-bc2a-6daceb97794e")

	for {
		log.Print("Enter Term (Synonym): ")
		input := new(string)
		fmt.Scan(input)
		var s *Resp = SynonymReq(*input)
		for _, item := range *s {
			if item.Meta.ID == *input {
				for _, item := range item.Meta.Syns {
					for _, i := range item {
						log.Println(i)
					}
				}
			}

		}
	}

}

type Resp []ResponseElement

type ResponseElement struct {
	Meta struct {
		ID      string `json:"id"`
		UUID    string `json:"uuid"`
		Src     string `json:"src"`
		Section string `json:"section"`
		Target  struct {
			Tuuid string `json:"tuuid"`
			Tsrc  string `json:"tsrc"`
		} `json:"target"`
		Stems     []string   `json:"stems"`
		Syns      [][]string `json:"syns"`
		Ants      [][]string `json:"ants"`
		Offensive bool       `json:"offensive"`
	} `json:"meta"`
	Hwi struct {
		Hw string `json:"hw"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Shortdef []string `json:"shortdef"`
}
