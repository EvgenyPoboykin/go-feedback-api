package utils

var version = "v1"

func VersionApiUrl(url string) string {
	return "/" + "api" + "/" + version + url
}
