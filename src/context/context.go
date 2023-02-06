// Package context includes the context of Lip.
package context

import (
	"os"
	"strings"

	"github.com/liteldev/lip/utils/version"
)

//------------------------------------------------------------------------------
// Constants

// Version is the version of Lip.
var VersionString = "v0.0.0"

const DefaultGoproxy = "https://goproxy.io"

const DefaultRegistryURL = "https://registry.litebds.com"

//------------------------------------------------------------------------------
// Variables

// Version is the version of Lip.
var Version version.Version

// Goproxy is the goproxy address.
var Goproxy string

// RegistryURL is the registry address.
var RegistryURL string

//------------------------------------------------------------------------------
// Functions

// Init initializes the
func Init() {
	var err error

	// Set Version.
	Version, err = version.NewFromString(strings.TrimPrefix(VersionString, "v"))
	if err != nil {
		Version, _ = version.NewFromString("0.0.0")
	}

	// Set Goproxy.
	if goproxy := os.Getenv("GOPROXY"); goproxy != "" {
		Goproxy = goproxy
	} else {
		Goproxy = DefaultGoproxy
	}

	// Set RegistryURL.
	if registryURL := os.Getenv("LIP_REGISTRY"); registryURL != "" {
		RegistryURL = registryURL
	} else {
		RegistryURL = DefaultRegistryURL
	}
}
