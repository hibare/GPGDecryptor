package version

import (
	"fmt"

	"github.com/hibare/GPGDecryptor/internal/constants"
	"github.com/hibare/GoCommon/pkg/version"
)

var (
	CurrentVersion = "0.0.0"
)

func GetNewVersionInfo() version.Version {
	v := version.Version{
		CurrentVersion: fmt.Sprintf("v%s", CurrentVersion),
		GithubOwner:    constants.GithubOwner,
		GithubRepo:     constants.GithubRepo,
	}
	v.CheckUpdate()

	return v
}
