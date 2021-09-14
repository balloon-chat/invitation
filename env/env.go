package env

import (
	"fmt"
	goEnv "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DEBUG                        bool
	VERSION                      string
	ClientEntryPoint             string
	GoogleApplicationCredentials string
)

type EnvironmentBase struct {
	VERSION                      string `env:"VERSION"`
	GoogleApplicationCredentials string `env:"GOOGLE_APPLICATION_CREDENTIALS"`
}

type Environment struct {
	ClientEntryPoint string `env:"CLIENT_ENTRY_POINT"`
}

func init() {
	loadEnv()
}

func loadEnv() {
	// Cloud Functionsの場合は、envファイルを読み込まない
	if VERSION = os.Getenv("VERSION"); VERSION == "CF" {
		DEBUG = false
		if ClientEntryPoint = os.Getenv("CLIENT_ENTRY_POINT"); ClientEntryPoint == "" {
			log.Fatalln("Environment variable CLIENT_ENTRY_POINT is empty")
		}
		return
	}

	// それ以外の場合は、envファイルから環境変数を読み込む
	loadBaseEnvFile()
	loadEnvFile()
}

func loadBaseEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error while loading .env file:", err)
	}

	var environment EnvironmentBase
	_, err = goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables", err)
	}

	// バージョン(テスト、プロダクション)の指定
	if VERSION = environment.VERSION; VERSION == "" {
		log.Fatalln("Environment variable VERSION is empty")
	}
	DEBUG = VERSION == "development"

	// Firebaseの秘密鍵のパス
	if DEBUG && environment.GoogleApplicationCredentials == "" {
		log.Fatalln("Environment variable GOOGLE_APPLICATION_CREDENTIALS is empty")
	}
	GoogleApplicationCredentials = environment.GoogleApplicationCredentials
}

func loadEnvFile() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", VERSION))
	if err != nil {
		log.Fatalln("error while loading .env file:", err)
	}

	var environment Environment
	_, err = goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables", err)
	}

	if ClientEntryPoint = environment.ClientEntryPoint; ClientEntryPoint == "" {
		log.Fatalln("Environment variable CLIENT_ENTRY_POINT is empty")
	}
}
