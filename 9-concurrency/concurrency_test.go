package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	sleepFor := random.Intn(1000)
	fmt.Printf("sleeping for %v\n", sleepFor)
	time.Sleep(time.Millisecond * time.Duration(sleepFor))
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
