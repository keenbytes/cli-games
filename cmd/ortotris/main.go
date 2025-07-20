// Package main is the entry point for the ortotris command-line game.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"

	"github.com/keenbytes/broccli/v3"
	"github.com/keenbytes/cli-games/pkg/ortotris"
)

func main() {
	cli := broccli.NewBroccli(
		"ortotris",
		"Clone of a classic Ortotris game",
		"Mikołaj Gąsior <m@gasior.dev>",
	)

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

func versionHandler(_ context.Context, _ *broccli.Broccli) int {
	_, _ = fmt.Fprintf(os.Stdout, VERSION+"\n")

	return 0
}

//nolint:contextcheck,mnd
func startHandler(ctx context.Context, cli *broccli.Broccli) int {
	game := ortotris.NewGame()

	wordsFile, err := os.Open(filepath.Clean(cli.Flag("words")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading words from file: %s", err.Error())

		return 1
	}

	defer func() {
		err := wordsFile.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error closing words file")
		}
	}()

	game.ReadWords(wordsFile)
	game.RandomizeWords()

	speed := cli.Flag("speed")
	if speed == "" {
		speed = "400"
	}

	speedInt, err := strconv.Atoi(speed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid speed value")

		return 1
	}

	gui := newGameInterface(game, speedInt)

	ctxGui, cancelGui := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	quit := make(chan struct{})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		gui.run(ctxGui, cancelGui)

		quit <- struct{}{}

		waitGroup.Done()
	}()
	go func() {
		for {
			select {
			case <-quit:
				waitGroup.Done()
			case <-sigs:
				cancelGui()
			case <-ctx.Done():
				cancelGui()
			}
		}
	}()

	waitGroup.Wait()

	return 0
}
