package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/keenbytes/broccli/v3"
	"github.com/keenbytes/cli-games/pkg/ortotris"
)

func main() {
	cli := broccli.NewBroccli("ortotris", "Clone of a classic Ortotris game", "Mikolaj Gasior")

	cmd := cli.Command("start", "Starts the game", startHandler)
	cmd.Flag(
		"words",
		"f",
		"",
		"Text file with wordlist",
		broccli.TypePathFile,
		broccli.IsExistent|broccli.IsRequired,
	)
	cmd.Flag("speed", "s", "", "Snake speed", broccli.TypeInt, 0)

	_ = cli.Command("version", "Shows version", versionHandler)

	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		os.Args = []string{"App", "version"}
	}

	cli.Run(context.Background())
}

func versionHandler(ctx context.Context, c *broccli.Broccli) int {
	fmt.Fprintf(os.Stdout, VERSION+"\n")

	return 0
}

func startHandler(ctx context.Context, c *broccli.Broccli) int {
	g := ortotris.NewGame()

	f, err := os.Open(c.Flag("words"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading words from file: %s", err.Error())

		return 1
	}
	defer f.Close()

	g.ReadWords(f)
	g.RandomizeWords()

	// TODO: speed
	speed := c.Flag("speed")
	if speed == "" {
		speed = "400"
	}

	speedInt, _ := strconv.Atoi(speed)

	gui := newGameInterface(g, speedInt)

	ctxGui, cancelGui := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	quit := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		gui.run(ctxGui, cancelGui)

		quit <- struct{}{}

		wg.Done()
	}()
	go func() {
		for {
			select {
			case <-quit:
				wg.Done()
			case <-sigs:
				cancelGui()
			case <-ctx.Done():
				cancelGui()
			}
		}
	}()

	wg.Wait()

	return 0
}
