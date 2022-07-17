package batch_common

import (
	"fmt"
	"os"
)

type ProjectEnv struct {
	Project     string
	Environment string
	Region      string
}

var Env ProjectEnv

func InitEnv() error {
	project := os.Getenv("PROJECT")
	environment := os.Getenv("ENV")
	region := os.Getenv("REGION")
	Env.Project = project
	Env.Environment = environment
	Env.Region = region
	fmt.Println(Env)
	fmt.Println("Env 초기화 완료")
	return nil
}
