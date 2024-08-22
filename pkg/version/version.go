package version

import "fmt"

const (
	major      = 0
	minor      = 1
	reversion  = 0
	subversion = 0
)

var TagVersion string // ldflags传递

func Version() string {
	if TagVersion != "" {
		return TagVersion
	} else {
		return fmt.Sprintf("V%v.%v.%v.%v", major, minor, reversion, subversion)
	}
}

func PrintVersion() {
	fmt.Printf("Version: %s\nBuild:   %s\n", Version(), Date)
}
