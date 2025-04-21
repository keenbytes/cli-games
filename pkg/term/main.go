package term

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// InitTTY initialises terminal window
func InitTTY() {
	cmd1 := exec.Command("stty", "cbreak", "min", "1")
	cmd1.Stdin = os.Stdin
	err := cmd1.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd2 := exec.Command("stty", "-echo")
	cmd2.Stdin = os.Stdin
	err = cmd2.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Clear clears terminal window
func Clear(stdout *os.File) {
	fmt.Fprintf(stdout, "\u001b[2J\u001b[1000A\u001b[1000D")
}

// GetSize gets terminal size by calling stty command
func GetSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	nums := strings.Split(string(out), " ")
	h, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, err
	}
	w, err := strconv.Atoi(strings.Replace(nums[1], "\n", "", 1))
	if err != nil {
		return 0, 0, err
	}
	return w, h, nil
}
