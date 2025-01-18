package buildrepocore

import (
	"buildrepo-core/internal/env"
	"buildrepo-core/internal/inference"
	"log"
	"testing"
)

func Test_GetInstructions(t *testing.T) {

	env.Init()

	uri := "https://github.com/icelain/radio"
	res, err := GetInstructions(uri)

	if err != nil {

		log.Fatal(err)

	}

	sendChannel := make(chan []byte)
	recvChannel := make(chan []byte)

	go inference.Request(sendChannel, recvChannel)

	for {

		sendChannel <- []byte(res)
		resp := <-recvChannel

		cmd := inference.MatchCommand(resp)

	}

}
