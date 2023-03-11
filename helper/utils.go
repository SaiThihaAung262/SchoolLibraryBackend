package helper

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GenerateUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	splitID := strings.Split(string(newUUID), "-")
	fmt.Println("Generated UUID:")
	fmt.Printf("%s", splitID[0])
	return splitID[0]
}
