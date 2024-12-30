package buildrepocore

import (
	"log"
	"testing"
)

func Test_GetInstructions(t *testing.T) {

	uri := "https://github.com/ollama/ollama"
	res, err := GetInstructions(uri)

	if err != nil {

		log.Fatal(err)

	}

	log.Println(res)

}
