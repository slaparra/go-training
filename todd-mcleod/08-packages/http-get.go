package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//do a get http request to get the file words.txt
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")

	//if error
	if err != nil {
		log.Fatal(err)
	}

	//if no error, read the response body
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close() //we could use defer to close the response at the end of the func

	//if error on reading
	if err != nil {
		log.Fatal(err)
	}

	//print words on file words.txt
	fmt.Printf("%s", bs)
}
