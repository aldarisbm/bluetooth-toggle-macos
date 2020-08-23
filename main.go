package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const appleClamshellState = "AppleClamshellState"

func main() {
	f, err := os.Create("app.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		clamshellMode, _ := isInClamshellMode()
		clamstring := fmt.Sprintf("Clamshell mode: %t\n", clamshellMode)
		isConnectedToACPower()
		l, err := f.WriteString(clamstring)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		time.Sleep(3 * time.Second)
	}
}

func isConnectedToACPower() (bool, error) {
	out, err := exec.Command("/usr/bin/pmset", "-g", "batt").Output()
	if err != nil {
		return false, fmt.Errorf("Error while running pmset command: %s", err)
	}
	outString := strings.Replace(string(out), "'", "", -1)
	outSplit := strings.Split(strings.Split(string(outString), "\n")[0], " ")
	fmt.Printf("%v\n", outSplit[len(outSplit)-2:])

	return false, nil
}

func isInClamshellMode() (bool, error) {
	out, err := exec.Command("/usr/sbin/ioreg", "-r", "-k", appleClamshellState).Output()
	if err != nil {
		return false, fmt.Errorf("Error while running ioreg command: %s", err)
	}
	outSplit := strings.Split(string(out), "\n")
	for _, line := range outSplit {
		if strings.Contains(line, appleClamshellState) {
			state := strings.Split(line, " ")
			return state[len(state)-1] == "Yes", nil
		}
	}
	return false, fmt.Errorf("Was not able to find \"%s\" running command", appleClamshellState)
}

// func toggleBluetooth() {

// }
