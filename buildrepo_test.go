package buildrepocore

import (
	"buildrepo-core/internal/inference"
	"log"
	"testing"
)

func Test_GetInstructions(t *testing.T) {

	uri := "https://github.com/icelain/radio"
	res, err := GetInstructions(uri)

	if err != nil {

		log.Fatal(err)

	}

	inference.Request([]byte(res))

}
