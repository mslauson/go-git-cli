package ggit

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing"
)

func getTag(repo *git.Repository) string {
	inputTag, ok := os.LookupEnv("INPUT_TAG")
	if !ok {
		return incPatch(repo)
	}
	return inputTag
}

func incPatch(repo *git.Repository) string {
	tags, err := repo.Tags()
	if err != nil {
		panic(err)
	}

	var latestTag string
	var versions []string
	re := regexp.MustCompile(`\d+\.\d+\.\d+`) // Matches semantic versioning

	err = tags.ForEach(func(t *plumbing.Reference) error {
		version := re.FindString(t.Name().String())
		if version != "" {
			versions = append(versions, version)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	sort.Strings(versions)
	if len(versions) > 0 {
		latestTag = versions[len(versions)-1]
	}

	// Increment the patch version
	parts := strings.Split(latestTag, ".")
	patch, _ := strconv.Atoi(parts[2])
	patch++
	parts[2] = strconv.Itoa(patch)
	newTag := strings.Join(parts, ".")
	return newTag
	// Create the new tag
}

func createTag(repo *git.Repository, tag string, commit string) {
	_, err := repo.CreateTag(
		tag,
		plumbing.NewHash(commit),
		nil,
	) // Replace "abc123" with your commit hash
	if err != nil {
		panic(err)
	}
}
