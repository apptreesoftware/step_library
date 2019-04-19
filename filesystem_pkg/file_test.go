package main

import (
	"github.com/pkg/errors"
	"testing"
)

func TestFileMove(t *testing.T) {
	input := FileMoveInput{
		FilePath:    "/Users/alexis/Downloads/staged2.jpg",
		ToDirectory: "/Users/alexis",
	}

	fileMove := FileMove{}
	output, err := fileMove.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !output.Success {
		t.Error(errors.New("move failed"))
		t.FailNow()
	}
}