package main

import (
	"create/createRelease"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	releaseDir, tarballPath, force, err := parseArgs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	imageName := "cloudfoundry/windows2016fs"

	imageTagPath := filepath.Join(releaseDir, "src", "code.cloudfoundry.org", "windows2016fs", "1803", "IMAGE_TAG")

	versionDataPath := filepath.Join(releaseDir, "VERSION")

	releaseCreator := new(createRelease.ReleaseCreator)
	err = releaseCreator.CreateRelease(imageName, releaseDir, tarballPath, imageTagPath, versionDataPath, force)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func parseArgs() (string, string, bool, error) {
	var releaseDir, tarballPath string
	var force bool
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)

	flagSet.StringVar(&releaseDir, "releaseDir", "", "")
	flagSet.StringVar(&tarballPath, "tarball", "", "")
	flagSet.BoolVar(&force, "force", false, "set -force=true to build a dev release")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return "", "", false, err
	}

	if releaseDir == "" {
		return "", "", false, errors.New("missing required flag 'releaseDir'")
	}

	return releaseDir, tarballPath, force, nil
}
