package gitmanager

import (
	"io"
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

func Delete(repository *Repository) error {

	return os.RemoveAll(repository.Path)

}

func ReadDir(path string, repository *Repository) ([]string, error) {

	wt, err := repository.Worktree()
	if err != nil {
		return []string{}, err
	}

	direntry := []string{}

	dir, err := wt.Filesystem.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	for _, v := range dir {

		direntry = append(direntry, v.Name())

	}

	return direntry, nil
}

func ReadFile(path string, repository *Repository) ([]byte, error) {

	wt, err := repository.Worktree()
	if err != nil {
		return []byte{}, err
	}

	file, err := wt.Filesystem.Open(path)
	if err != nil {
		return []byte{}, err
	}

	return io.ReadAll(file)

}
