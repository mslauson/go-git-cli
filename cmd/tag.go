/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	th "gitea.slauson.io/mslauson/ggit/thelper"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

var tagChoices = []string{"Patch", "Minor", "Major"}

func initialTagChoices() th.TeaChoices {
	return th.TeaChoices{
		// Our to-do list is a grocery list
		Choices: tagChoices,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		Selected: make(map[int]struct{}),
	}
}

// tagCmd represents the tag command
// This is used for tagging the repo
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tag helper",
	Long:  `tag helper for tagging the repo`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen("./")

		p := tea.NewProgram(initialTagChoices())
		val, err := p.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		switch val.View() {
		case "Patch":
			handlePatch(repo)
		}
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}

func handlePatch(repo *git.Repository) {
	latestTag := incPatch(repo)

	createTag(repo, latestTag, "HEAD")
}

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
