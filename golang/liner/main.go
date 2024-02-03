package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/peterh/liner"
)

var (
	history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	f, err := os.Open(history_fn)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if _, err := line.ReadHistory(f); err != nil {
			return err
		}
	}
	defer f.Close()

	var promptErr error
	for {
		l, err := line.Prompt("deepl> ")
		if err != nil {
			promptErr = err
			break
		}
		if l == "exit" {
			break
		}
		fmt.Printf("got: %s\n", l)
		line.AppendHistory(l)
	}

	ff, err := os.Create(history_fn)
	if err != nil {
		return err
	}
	defer ff.Close()
	if _, err := line.WriteHistory(ff); err != nil {
		return err
	}

	return promptErr
}
