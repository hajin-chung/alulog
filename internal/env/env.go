package env

import "os"

var AwsAccessKey string
var AwsSecret string
var PostPath string
var Secret string

func LoadEnv() {
	AwsAccessKey, _ = os.LookupEnv("AWS_ACCESS_KEY_ID")
	AwsSecret, _ = os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	Secret, _ = os.LookupEnv("DEPS_SECRET")
	PostPath, _ = os.LookupEnv("DEPS_POST_PATH")
}
