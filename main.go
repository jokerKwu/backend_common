package main

import (
	"fmt"
	common "main/batch_common"
	"os"
)

func main() {
	fmt.Println("main start")
	os.Setenv("PROJECT", "medical_web")
	os.Setenv("ENV", "dev")
	os.Setenv("REGION", "ap-north-2")
	if err := common.InitEnv(); err != nil {
		return
	}

	if err := common.InitAws("ap-north-2"); err != nil {
		return
	}
	if err := common.InitMongoDB(); err != nil {
		return
	}
	fmt.Println("main test success")
	return
}
