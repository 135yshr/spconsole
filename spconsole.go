package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"spherogo"
	"strconv"
	"strings"
)

var (
	devideId string
)

const (
	lf      = '\n'
	cr      = '\r'
	newline = "\r\n"
)

func main() {

	flag.Parse()

	Stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("sphero > ")
		line, err := Stdin.ReadString(lf)
		printError(err)
		line = strings.TrimRight(line, newline)

		commands := strings.Split(line, " ")
		switch strings.ToLower(commands[0]) {
		case "rgb":
			fmt.Println(commands[1:])
			err := commandRgb(commands[1:])
			printError(err)
		case "roll":
			err := commandRoll(commands[1:])
			printError(err)
		case "version":
		case "exit", "quit":
			fmt.Println("bye...")
			os.Exit(0)
		default:
		}
	}
}

func init() {
	flag.StringVar(&devideId, "d", "", "sphero device file")
}

func commandRgb(commands []string) error {
	sphero := spherogo.NewSphero(devideId)
	switch strings.ToLower(commands[0]) {
	case "get":
		sphero.GetRgb()
	case "set":
		red, green, blue, err := string2color(commands[1:])
		if err != nil {
			return err
		}
		persistent, err := strconv.ParseBool(commands[len(commands)-1])
		if err != nil {
			return err
		}
		sphero.SetRgb(red, green, blue, persistent)
	default:
		return fmt.Errorf("not support parameter")
	}
	return nil
}

func commandRoll(commands []string) error {
	if len(commands) < 3 {
		return fmt.Errorf("roll <speed> <heading> <state>")
	}
	speed, err := parse2Byte(commands[0])
	if err != nil {
		return err
	}
	heading, err := parse2Uint16(commands[1])
	if err != nil {
		return err
	}
	state, err := strconv.ParseBool(commands[2])
	if err != nil {
		return err
	}
	sphero := spherogo.NewSphero(devideId)
	sphero.Roll(speed, heading, state)
	return nil
}

func string2color(param []string) (byte, byte, byte, error) {
	switch param[0] {
	case "red":
		return 0xFF, 0x00, 0x00, nil
	case "green":
		return 0x00, 0xFF, 0x00, nil
	case "blue":
		return 0x00, 0x00, 0xFF, nil
	default:
		if len(param) < 3 {
			return 0xFF, 0xFF, 0xFF, fmt.Errorf("not support param count")
		}
		red, err := strconv.ParseUint(param[0], 10, 8)
		if err != nil {
			return 0xFF, 0xFF, 0xFF, err
		}
		green, err := strconv.ParseUint(param[1], 10, 8)
		if err != nil {
			return 0xFF, 0xFF, 0xFF, err
		}
		blue, err := strconv.ParseUint(param[2], 10, 8)
		if err != nil {
			return 0xFF, 0xFF, 0xFF, err
		}
		return uint8(red), uint8(green), uint8(blue), nil
	}
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func parse2Byte(s string) (byte, error) {
	ret, err := strconv.ParseUint(s, 10, 8)
	return byte(ret), err
}

func parse2Uint16(s string) (uint16, error) {
	ret, err := strconv.ParseUint(s, 10, 16)
	return uint16(ret), err
}
