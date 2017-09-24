//
//
// I tried for a long time to make tests,
// but they wouldn't work because I don't understand it enough :(
//
package main

import (
	"testing"
	"net/http"
	"time"
)
/*
type  testJson struct{
	Comitter string `json:"login"`
	Comitts int `json:"contributions"`
}


func TestGetComitter(t *testing.T){

	// Test Json
	tj := testJson{"gitster", 18497}

	// Test URL
	tu := "http://api.github.com/repos/git/git/contributors"

	// Client
	tc := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, tu, nil)
	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Testing")

	comitter := getOwner(tu, tc)

	if comitter != tj{
		t.Fatalf("Error ", tj, comitter)
	}

}
*/


func TestGetOwner(t *testing.T){
	out := "git"
	client := http.Client{
		Timeout: time.Second *2,
	}
	owner := getOwner("http://api.github.com/repos/git/git/", client)

	if owner != out{
		t.Fatalf("Error got '%s' instead of '%s'", owner, out)
	}

}
