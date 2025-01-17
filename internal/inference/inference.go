package inference

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const PROMPT string = `

	You will be provided with a direntry of a git repository. Directories will be marked with [DIR] and files will be marked with [FILE].
	Your task is to compute the terminal commands that are required to run the project whose direntry you're provided with.

	You can only communicate in the following way:
	LISTDIR {path of directory in the project, which you've confirmed exists through the provided direntry} -> Sending this message will provide you with the direntry on a listed directory you're aware about.
	READFILE {path file in the project, which you've confirmed exists through the provided direntry} -> Sending this message will provide you with the contents of the file.
	OUTPUT {output terminal command if you're convinced with your analysis of the project}

`

var GROQ_API_KEY = os.Getenv("GROQ_API_KEY")

func Request(data []byte) error {

	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", GROQ_API_KEY))

	content := map[string]any{

		"model": "llama-3.3-70b-versatile",
		"messages": []map[string]string{{
			"role":    "user",
			"content": "Explain the importance of fast language models",
		},
			{
				"role":    "system",
				"content": PROMPT,
			},
		},
	}

	requestBody := &bytes.Buffer{}
	if err := json.NewEncoder(requestBody).Encode(content); err != nil {

		return err

	}

	req, err := http.NewRequest(http.MethodPost, "https://api.groq.com/openai/v1/chat/completions", requestBody)
	if err != nil {
		return err
	}

	req.Header = headers

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		return err

	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)

	fmt.Println(respBytes)

	return nil

}
