/////////////////////////////////////////////////////////////////////
// 		   AUTHOR: 	Brede Fritjof Klausen		  				  //
//	STUDENTNUMBER: 	473211						 				 //
// 		  SUBJECT: 	IMT2681 Cloud Technologies					//
//=============================================================//
//	SOURCES:												  //
// * 	Nata Niel (Code for getting languages)				 //
// *	https://blog.alexellis.io/golang-json-api-client/	//
/////////////////////////////////////////////////////////////
package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"
	"time"
	"os"
)

type Information struct {
	Project  string   `json:"project"`
	Owner    string	  `json:"owner"`
	Commiter string   `json:"commiter"`
	Commits  int 	  `json:"commits"`
	Language []string `json:"language"`
}

type Login struct{
	Login    string   `json:"login"`
}

type Repos struct{
	Login    Login 	  `json:"owner"`
}

type Comitter struct{
	Comitter string   `json:"login"`
	Commits  int 	  `json:"contributions"`
}

func getOwner(url string, gitClient http.Client)string{

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "assignment")

	res, getErr := gitClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	Rep := Repos{}
	jsonErr := json.Unmarshal(body, &Rep)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	return Rep.Login.Login
}

func getComitter(url string, gitClient http.Client)Comitter{

	comUrl := url + "/contributors"
	req, err := http.NewRequest(http.MethodGet, comUrl, nil)
	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "assignment")

	res, getErr := gitClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	Com := []Comitter{}
	jsonErr := json.Unmarshal(body, &Com)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}

	return Com[0]
}

func getLanguage(url string, gitClient http.Client)[]string{

	comUrl := url + "/languages"
	req, err := http.NewRequest(http.MethodGet, comUrl, nil)
	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "assignment")

	res, getErr := gitClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var Languages []string
	LanguageMap := make(map [string] int)
	jsonErr := json.Unmarshal(body, &LanguageMap)
	if jsonErr != nil{
		log.Fatal(jsonErr)
	}
	for key := range LanguageMap{
		Languages = append(Languages, key)
	}

	return Languages
}

func HandlerGitHub(w http.ResponseWriter, r *http.Request) {

	// DECLARE IT'S A JSON FILE
	http.Header.Add(w.Header(), "Content-type", "application/json")

	// SPLIT URL FOR EACH "/"
	parts := strings.Split(r.URL.Path, "/")

	// CREATE THE GENERAL URL
	url := "http://api." + parts[3] + "/repos/" + parts[4] + "/" + parts[5]

	// CREATE CLIENT
	gitClient := http.Client{
		Timeout: time.Second * 2,
	}

	// PROJECT
	project := "https://" + parts[3] + "/" + parts[4] + "/" + parts[5]

	// OWNER
	owner := getOwner(url, gitClient)

	// COMITTER & HIS/HERS COMMITS
	comitter := getComitter(url, gitClient)

	// LANGUAGE
	language := getLanguage(url, gitClient)

	// MAKE JSON \\
	jsonFile := Information{project, owner, comitter.Comitter, comitter.Commits, language}
	json.NewEncoder(w).Encode(jsonFile)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/projectinfo/v1/",	HandlerGitHub)
	http.ListenAndServe(":"+port, nil)
}