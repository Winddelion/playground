package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SearchConfig struct {
	IgnoreCase bool
	Invert     bool
	LineNumber bool
	Pattern    string
	FileName   string
}

type Match struct {
	LineNumber  int32
	MatchedWord string
}

func Search(cfg SearchConfig) ([]Match, error) {
	var matches []Match
	counter := 0

	file, err := os.Open(cfg.FileName)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", cfg.FileName, err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for scan.Scan() {
		counter++
		line := scan.Text()
		if strings.Contains(line, cfg.Pattern) {
			matches = append(matches, Match{
				LineNumber:  int32(counter),
				MatchedWord: line,
			})
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}
