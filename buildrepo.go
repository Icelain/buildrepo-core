package buildrepocore

import (
	"buildrepo-core/internal/gitmanager"
	"path"
	"strings"
)

func GetInstructions(repositoryUrl string) (string, error) {

	repo, err := gitmanager.Clone(repositoryUrl)
	if err != nil {

		return "", err

	}

	worktree, err := repo.Worktree()
	if err != nil {

		return "", err

	}

	res := &strings.Builder{}
	files, err := worktree.Filesystem.ReadDir("./")
	if err != nil {

		return "", err

	}

	currentDir := "./"
	for _, f := range files {

		p := path.Join(currentDir, f.Name())
		res.WriteString(p)

		if f.IsDir() {
			res.WriteString(" [DIR]")
		} else {

			res.WriteString(" [FILE]")

		}

		res.WriteByte('\n')

	}

	s := res.String()
	s = s[:len(s)-1]

	return s, nil

}
