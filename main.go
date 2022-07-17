package main

import (
	"fmt"
	"main/batch_common/aws/ssm"
	"main/batch_common/db"
)

func main() {
	fmt.Println("main startㄴ")
	if err := ssm.InitAws("ap-north-2"); err != nil {
		return
	}
	if err := db.InitMongoDB(); err != nil {
		return
	}
	fmt.Println("main test success")
	return
}
