package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const REQUEST_METHOD = "GET"

func SynonymReq(Input string) *Resp {
	client := new(http.Client)

	req, err := http.NewRequest(REQUEST_METHOD, fmt.Sprintf("https://dictionaryapi.com/api/v3/references/thesaurus/json/%s?key=%s", Input, os.Getenv("THESAURUS_API_KEY")), nil)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := new(Resp)

	err = json.Unmarshal([]byte(body), &r)
	if err != nil{
		log.Fatalln("Failed to unmarshal response: ", err)
	}

	return r;
}