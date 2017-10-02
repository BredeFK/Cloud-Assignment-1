package main

import (
	"testing"
	"net/http"
	"time"
)


func TestGetComitter(t *testing.T){

	// Test Json
	tj := Comitter{"gitster", 18497}

	// Test URL
	tu := "http://api.github.com/repos/git/git"

	// Client
	tc := http.Client{
		Timeout: time.Second * 2,
	}

	comitter := getComitter(tu, tc)

	if comitter.Comitter != tj.Comitter{
		t.Fatalf("Error got '%s' instead of '%s'",comitter.Comitter, tj.Comitter)
	}

}



func TestGetOwner(t *testing.T){
	out := "git"
	client := http.Client{
		Timeout: time.Second *2,
	}
	owner := getOwner("http://api.github.com/repos/git/git", client)

	if owner != out{
		t.Fatalf("Error got '%s' instead of '%s'", owner, out)
	}

}
