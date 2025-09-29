package version

import (
	"fmt"

	"github.com/hibare/GPGDecryptor/internal/constants"
	"github.com/hibare/GoCommon/v2/pkg/version"
)

var (
	CurrentVersion = "0.0.0"
)

func GetNewVersionInfo() version.Versions {
	v, err := version.NewVersion(version.Options{
		CurrentVersion: fmt.Sprintf("v%s", CurrentVersion),
		GithubOwner:    constants.GithubOwner,
		GithubRepo:     constants.GithubRepo,
	})
	if err != nil {
		panic(err)
	}

	v.CheckUpdate()

	return v
}
