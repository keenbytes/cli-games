// Package term contains some helper functions.
package term

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	errTerminalSize = errors.New("error getting terminal size")
	errTerminalSet  = errors.New("error setting terminal")
)

func errRunningSttySize(err error) error {
	return fmt.Errorf("%w: %s", errTerminalSize, err.Error())
}

func errRunningSttyCbreak(err error) error {
	return fmt.Errorf("%w: %s", errTerminalSet, err.Error())
}

func errRunningSttyEcho(err error) error {
	return fmt.Errorf("%w: %s", errTerminalSet, err.Error())
}

// InitTTY initialises terminal window.
func InitTTY() {
	cmd1 := exec.Command("stty", "cbreak", "min", "1")
	cmd1.Stdin = os.Stdin

	err := cmd1.Run()
	if err != nil {
		log.Fatal(errRunningSttyCbreak(err))
	}

	cmd2 := exec.Command("stty", "-echo")
	cmd2.Stdin = os.Stdin

	err = cmd2.Run()
	if err != nil {
		log.Fatal(errRunningSttyEcho(err))
	}
}

// Clear clears terminal window.
func Clear(stdout *os.File) {
	_, _ = fmt.Fprintf(stdout, "\u001b[2J\u001b[1000A\u001b[1000D")
}

// GetSize gets terminal size by calling stty command.
func GetSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		return 0, 0, errRunningSttySize(err)
	}

	nums := strings.Split(string(out), " ")

	height, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, errRunningSttySize(err)
	}

	width, err := strconv.Atoi(strings.Replace(nums[1], "\n", "", 1))
	if err != nil {
		return 0, 0, errRunningSttySize(err)
	}

	return width, height, nil
}
