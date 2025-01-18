package buildrepocore

import (
	"buildrepo-core/internal/env"
	"buildrepo-core/internal/gitmanager"
	"buildrepo-core/internal/inference"
	"log"
	"testing"
)

func Test_GetInstructions(t *testing.T) {

	env.Init()

	uri := "https://github.com/icelain/radio"

	repo, err := gitmanager.Clone(uri)
	if err != nil {
		t.Fatal(err)
	}

	res, err := GetInstructions(repo)

	if err != nil {

		log.Fatal(err)

	}

	sendChannel := make(chan []byte)
	recvChannel := make(chan []byte)

	go inference.Request(sendChannel, recvChannel)

	for {

		sendChannel <- []byte(res)
		resp := <-recvChannel

		cmd, args := inference.MatchCommand(resp)
		apiResp, err := inference.HandleCommand(args, cmd, repo)
		if err != nil {

			t.Fatal(err)

		}

		res = string(apiResp.Content)

	}

}
