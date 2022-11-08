//go:build mage

package main

import (
	"context"

	"github.com/magefile/mage/sh"

	"github.com/eunomie/dague/types"
	//mage:import
	_ "github.com/eunomie/dague/target/lint"
	//mage:import
	"github.com/eunomie/dague/target/golang"

	"github.com/magefile/mage/mg"
)

var (
	SnykDesktopVersion = "1.1025.0"
	SnykImageDigest    = "sha256:b979e1827473ce7675439213a918687ac532481c3370818bce61884735bdb09d"
	commit             string
	tagName            string
	envVars            = map[string]string{
		"CGO_ENABLED": "0",
		"GO111MODULE": "on",
	}
	buildFlags []string
)

func init() {
	var err error
	commit, err = sh.OutCmd("git", "rev-parse", "--short", "HEAD")()
	if err != nil {
		panic(err)
	}

	tagName, err = sh.Output("git", "describe", "--always", "--dirty", "--abbrev=10")
	if err != nil {
		panic(err)
	}

	buildFlags = []string{
		"-trimpath",
		"-ldflags=-s -w" +
			" -X 'github.com/docker/scan-cli-plugin/internal.GitCommit=" + commit + "'" +
			" -X 'github.com/docker/scan-cli-plugin/internal.Version=" + tagName + "'" +
			" -X 'github.com/docker/scan-cli-plugin/internal/provider.ImageDigest=" + SnykDesktopVersion + "'" +
			" -X 'github.com/docker/scan-cli-plugin/internal/provider.SnykDesktopVersion=" + SnykImageDigest + "'",
	}

}

type Build mg.Namespace

// Local builds local binary of docker-scan, for the running platform
func (Build) Local(ctx context.Context) error {
	return golang.Local(ctx, types.LocalBuildOpts{
		BuildOpts: types.BuildOpts{
			Dir:        "dist",
			In:         "./cmd/docker-scan",
			EnvVars:    envVars,
			BuildFlags: buildFlags,
		},
		Out: "docker-scan",
	})
}

// Cross builds docker-scan binaries for all the supported platforms
func (Build) Cross(ctx context.Context) error {
	return golang.Cross(ctx, types.CrossBuildOpts{
		BuildOpts: types.BuildOpts{
			Dir:        "dist",
			In:         "./cmd/docker-scan",
			EnvVars:    envVars,
			BuildFlags: buildFlags,
		},
		OutFileFormat: "docker-scan_%s_%s",
		Platforms: []types.Platform{
			{"linux", "amd64"},
			{"linux", "arm64"},
			{"darwin", "amd64"},
			{"darwin", "arm64"},
			{"windows", "amd64"},
		},
	})
}
