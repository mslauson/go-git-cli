/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	th "gitea.slauson.io/mslauson/ggit/thelper"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

// tagChoices is the list of choices for the tag command
var tagChoices = []list.Item{th.Item("Patch"), th.Item("Minor"), th.Item("Major")}

// initialTagChoices creates the initial bubble tea list model for the tag commandj
func initialTagChoices() th.ListModel {
	const width = 20
	const height = 14

	l := list.New(tagChoices, th.ItemDelegate{}, width, height)
	l.Title = "What Type of Tag?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = th.TitleStyle
	l.Styles.PaginationStyle = th.PaginationStyle
	l.Styles.HelpStyle = th.HelpStyle
	return th.ListModel{List: l}
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

		trimVal := strings.ReplaceAll(val.View(), "\n", " ")
		trimVal = strings.TrimSpace(trimVal)
		spaceReg := regexp.MustCompile(`\s+`)
		op := spaceReg.Split(trimVal, -1)

		log.Println("op:", op)
		log.Println("op:", op[1])

		switch op[1] {
		case "Patch":
			handlePatch(repo)
		}
	},
}

// Adds tagCmd to rootCmd
func init() {
	rootCmd.AddCommand(tagCmd)
}

// handlePatch handles the patch tag choice and increments the patch tag
func handlePatch(repo *git.Repository) {
	log.Println("patch")
	latestTag := incPatch(repo)

	log.Println("Creating tag", latestTag)
	createTag(repo, latestTag, "HEAD")
}

// incPatch increments the patch version of the latest tag
func incPatch(repo *git.Repository) string {
	tags, err := repo.Tags()
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}

	sort.Strings(versions)
	if len(versions) > 0 {
		latestTag = versions[len(versions)-1]
	}

	parts := strings.Split(latestTag, ".")
	patch, _ := strconv.Atoi(parts[2])
	patch++
	parts[2] = strconv.Itoa(patch)
	newTag := strings.Join(parts, ".")

	fmt.Println(newTag)
	return newTag
}

// createTag creates a new tag on HEAD
func createTag(repo *git.Repository, tag string, commit string) error {
	_, err := repo.CreateTag(
		tag,
		plumbing.NewHash(commit),
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
