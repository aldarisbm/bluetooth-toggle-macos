package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	appleClamshellState = "AppleClamshellState"
	on                  = "1"
	off                 = "0"
)

func main() {
	f, err := os.Create("app.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		_ = runJob()
		l, err := f.WriteString("boop")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		fmt.Println(l, "bytes written successfully")
		time.Sleep(3 * time.Second)
	}
}

func runJob() error {
	connectedToPower, err := isConnectedToACPower()
	if err != nil {
		return err
	}
	lidClosed, err := isLidCLosed()
	if err != nil {
		return err
	}
	if !connectedToPower && lidClosed {
		state, err := isBluetoothOff()
		if err != nil {
			return err
		}
		if state {
			return nil
		}
		turnOffBluetooth()
	}
	if connectedToPower {
		turnOnBluetooth()
	}
	return nil
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

func isLidCLosed() (bool, error) {
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

func turnOffBluetooth() error {
	if err := exec.Command("/usr/local/bin/blueutil", "-p", off).Run(); err != nil {
		return err
	}
	return nil
}

func turnOnBluetooth() error {
	if err := exec.Command("/usr/local/bin/blueutil", "-p", on).Run(); err != nil {
		return err
	}
	return nil
}

func isBluetoothOff() (bool, error) {
	out, err := exec.Command("/usr/local/bin/blueutil", "-p").Output()
	if err != nil {
		return false, err
	}
	outString := string(out)
	fmt.Printf("output: %s\n", outString)
	return outString == off, nil
}
