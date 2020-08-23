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
		l, err := f.WriteString(clamstring)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		connectedToACPower, _ := isConnectedToACPower()
		connectedString := fmt.Sprintf("Connected to AC Power: %t\n", connectedToACPower)
		l, err = f.WriteString(connectedString)
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
	cleansedString, err := cleansePmsetString(string(out))
	if err != nil {
		return false, err
	}
	return cleansedString == "AC Power", nil
}

func cleansePmsetString(out string) (string, error) {
	equalityStringCheck := "Now drawing from "
	outSplit := strings.Split(string(out), "\n")
	for _, line := range outSplit {
		if strings.Contains(line, equalityStringCheck) {
			outString := strings.Replace(line, "'", "", -1)
			outString = strings.Replace(outString, equalityStringCheck, "", -1)
			return outString, nil
		}
	}
	return "", fmt.Errorf("Was not able to find \"%s\" running command", equalityStringCheck)
}

func isInClamshellMode() (bool, error) {
	out, err := exec.Command("/usr/sbin/ioreg", "-r", "-k", appleClamshellState).Output()
	if err != nil {
		return false, fmt.Errorf("Error while running ioreg command: %s", err)
	}
	cleansedString, err := cleanseIoregString(string(out))
	if err != nil {
		return false, err
	}
	return cleansedString == "Yes", nil
}

func cleanseIoregString(out string) (string, error) {
	outSplit := strings.Split(out, "\n")
	for _, line := range outSplit {
		if strings.Contains(line, appleClamshellState) {
			state := strings.Split(line, " ")
			return state[len(state)-1], nil
		}
	}
	return "", fmt.Errorf("Was not able to find \"%s\" running command", appleClamshellState)
}

// func toggleBluetooth() {

// }
