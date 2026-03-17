package version

import (
	"fmt"

	"github.com/hibare/GPGDecryptor/internal/constants"
	"github.com/hibare/GoCommon/v2/pkg/version"
)

var (
	CurrentVersion = "0.0.0"
)

func GetNewVersionInfo() version.VersionIface {
	v := version.NewVersion(constants.GithubOwner, constants.GithubRepo, fmt.Sprintf("v%s", CurrentVersion), version.Options{})

	v.CheckUpdate()

	return v
}
