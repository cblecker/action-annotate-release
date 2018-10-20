package main // import "github.com/cblecker/action-annotate-release"

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"encoding/json"
	"github.com/google/go-github/v18/github"
	"golang.org/x/oauth2"
)

func main() {
	// pull in the event data file path from the environment
	jsonFilePath := os.Getenv("GITHUB_EVENT_PATH")
	// open the json event file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	// unmarshal the json into a release event
	var releaseEvent github.ReleaseEvent
	err = json.Unmarshal(jsonByte, &releaseEvent)
	if err != nil {
		fmt.Println(err)
	}

	// pull in token from the environment
	ghToken := os.Getenv("GITHUB_TOKEN")

	// set up connection to GitHub
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// pull in data to annotate to release
	annotationFilePath := os.Getenv("ANNOTATION_FILE")
	// Open the annotation file
	annotationFile, err := os.Open(annotationFilePath)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our annotation so that we can parse it later on
	defer annotationFile.Close()

	// read our opened annotation as a byte array.
	annotationByte, _ := ioutil.ReadAll(annotationFile)
	// read annotation file into a string
	annotateBody := fmt.Sprintf("```\n%s\n```", string(annotationByte[:]))

	// get the current release data from GitHub
	release, _, err := client.Repositories.GetRelease(ctx, *releaseEvent.Repo.Owner.Login, *releaseEvent.Repo.Name, *releaseEvent.Release.ID)
	if err != nil {
		fmt.Println(err)
	}

	// format and set up the new body
	newReleaseBody := strings.Join([]string{*release.Body, annotateBody}, "\n")
	release.Body = &newReleaseBody

	// edit the current release on GitHub to append the body
	_, _, err = client.Repositories.EditRelease(ctx, *releaseEvent.Repo.Owner.Login, *releaseEvent.Repo.Name, *releaseEvent.Release.ID, release)
	if err != nil {
		fmt.Println(err)
	}
}
