package version

// Version is provided by ldflags at compile time
var Version = "(devel)"

// Get returns the version
func Get() string {
	return Version
}
