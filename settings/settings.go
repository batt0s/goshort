package settings

import (
	"log"
	"os"
	"strings"
)

const HOST string = "localhost"
const PORT string = "8080"

var BASEPATH string = getBasePath()
var SECRETKEY string = getSecretFromEnv()

func getBasePath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("ERROR while getting BASEPATH\n%s", err.Error())
	}
	return path
}

func getSecretFromEnv() string {
	secret := os.Getenv("SECRET")
	if strings.TrimSpace(secret) == "" {
		log.Println("No SECRET in env.")
		secret = "ultimateUs31TAsecretkeyolc0IHuxdn9h8FyIe6GpzQkP3ZbBznJ3"
	}
	return secret
}
