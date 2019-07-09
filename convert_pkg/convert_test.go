package main

import (
	"errors"
	"os"
	"testing"
)

func TestShortenUrl_Execute(t *testing.T) {
	firebaseApi := os.Getenv("FIREBASE_API_KEY")
	urlPrefix := os.Getenv("FIREBASE_URL_PREFIX")
	origUrl := "https://google.com/?someKey=someVal&thisKey=thisVal"

	input := ShortenUrlInput{
		FirebaseApiKey:    firebaseApi,
		FirebaseUrlPrefix: urlPrefix,
		UrlToShorten:      origUrl,
	}

	shorten := ShortenUrl{}
	resp, err := shorten.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output, ok := resp.(*ShortenUrlOutput); ok {
		if output.ShortUrl == "" {
			t.Error(errors.New("url is empty"))
			t.FailNow()
		}
		return
	}
	t.Error("response was not s ShortenUrlResponse")
	t.Fail()
}
