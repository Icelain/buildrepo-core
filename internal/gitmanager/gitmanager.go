package gitmanager

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

type Repository struct {
	Path string
	git.Repository
}

func Clone(url string) (*Repository, error) {

	unique_uuid, err := uuid.NewV7()
	if err != nil {

		return &Repository{}, err

	}

	gitRepoPath := "/tmp/" + unique_uuid.String()

	repo, err := git.PlainClone(gitRepoPath, false, &git.CloneOptions{

		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		return &Repository{}, err
	}

	return &Repository{Repository: *repo, Path: gitRepoPath}, nil

}
