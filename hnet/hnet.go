package hnet

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"timestamps/filters"
)

//===================== URL SOURCE SELECTION ========================

func Url_selection(url string) {

	if strings.Contains(url, "https://github.com/rwxrob/boost/") {

		fmt.Println("You selected GitHub")

		filters.Filter_github(Get_url(url))

	} else if strings.Contains(url, "https://www.youtube.com/watch?v=") {

		fmt.Println("You selected YOUTUBE")

		filters.Filter_youtube(Get_url(url))

	} else {
		fmt.Println(" invalid URL ")
	}

}

//========================== GET URL ================================

func Get_url(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("The status code we got is:", resp.StatusCode)
		fmt.Println("The status code text we got is:", http.StatusText(resp.StatusCode))
		os.Exit(0)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	str := string(body)
	return str
}

//===================================================================
//===================================================================
