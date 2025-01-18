package inference

import (
	"buildrepo-core/internal/gitmanager"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Command int

const (
	ReadFile Command = iota
	ReadDir
	Output
	Invalid
)

type InferenceResponse struct {
	Content  []byte
	Continue bool
}

const PROMPT string = `

	You will be provided with a direntry of a git repository. Directories will be marked with [DIR] and files will be marked with [FILE].
	Your task is to compute the terminal commands that are required to run the project whose direntry you're provided with.

	You can only communicate in the following way:
	LISTDIR {path of directory in the project, which you've confirmed exists through the provided direntry} -> Sending this message will provide you with the direntry on a listed directory you're aware about.
	READFILE {path file in the project, which you've confirmed exists through the provided direntry} -> Sending this message will provide you with the contents of the file.
	OUTPUT {output terminal command if you're convinced with your analysis of the project}

	Do not say anything other than the way you're instructed to communicate in.

`

func Request(dataInStream <-chan []byte, dataOutStream chan<- []byte) error {

	var GROQ_API_KEY = os.Getenv("GROQ_API_KEY")
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", GROQ_API_KEY))

	content := map[string]any{

		"model": "llama-3.3-70b-versatile",
		"messages": []map[string]string{{
			"role":    "user",
			"content": string(<-dataInStream),
		},
			{
				"role":    "system",
				"content": PROMPT,
			},
		},
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.groq.com/openai/v1/chat/completions", nil)
	if err != nil {
		return err
	}

	for {

		requestBody := &bytes.Buffer{}
		if err := json.NewEncoder(requestBody).Encode(content); err != nil {

			return err

		}

		req.Body = io.NopCloser(requestBody)
		req.Header = headers

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {

			return err

		}

		defer resp.Body.Close()

		respBytes, err := io.ReadAll(resp.Body)

		dataOutStream <- respBytes

	}

}

func MatchCommand(cmd []byte) (Command, [][]byte) {

	args := bytes.Split(cmd, []byte(" "))

	switch string(args[0]) {

	case "READFILE":

		return ReadFile, args[1:]

	case "LISTDIR":
		return ReadDir, args[1:]

	case "OUTPUT":
		return Output, args[1:]

	}

	return Invalid, [][]byte{}

}

func HandleCommand(args [][]byte, cmdtype Command, repository *gitmanager.Repository) (InferenceResponse, error) {

	switch cmdtype {

	case ReadDir:
		dirEntry, err := gitmanager.ReadDir(string(args[0]), repository)
		if err != nil {

			return InferenceResponse{}, err

		}

		response := strings.Join(dirEntry, "")
		return InferenceResponse{Content: []byte(response), Continue: true}, nil

	case ReadFile:
		fileContent, err := gitmanager.ReadFile(string(args[0]), repository)
		if err != nil {
			return InferenceResponse{}, err
		}
		return InferenceResponse{Content: fileContent, Continue: true}, nil

	case Output:
		return InferenceResponse{Content: args[0], Continue: false}, nil

	case Invalid:
		return InferenceResponse{Content: []byte{}, Continue: false}, nil

	}

	return InferenceResponse{}, nil

}
