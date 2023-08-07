package config

import (
	"runtime"

	"github.com/carlmjohnson/versioninfo"
)

// Automatically set by the compiler or build script
var (
	Version   string
	Build     string
	BuildTime string
	OS        string
)

func GetVersion() string {
	if Version == "" {
		return versioninfo.Version
	}
	return Version
}

func GetBuild() string {
	if Build == "" {
		return versioninfo.Revision
	}
	return Build
}

func GetBuildTime() string {
	if BuildTime == "" {
		return versioninfo.LastCommit.Local().Format("2006-01-02T15:04+00:00")
	}
	return BuildTime
}

func GetOS() string {
	if OS == "" {
		return runtime.GOOS
	}
	return OS
}
