package helper

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
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

func AddSevenDay(myTime time.Time) time.Time {
	return myTime.AddDate(0, 0, 7)
}

func CalculatExpireDate(myTime time.Time, addDay int) time.Time {
	return myTime.AddDate(0, 0, addDay)
}
