package system

// AppVersion ...
var AppVersion string

func init() {
	if AppVersion == "" {
		AppVersion = "draft"
	}
}

type Version struct {
	App       string
	GoVersion string
	Commit    string
}
