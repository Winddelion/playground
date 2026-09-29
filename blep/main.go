package main

//TODO:

import (
	"fmt"
	"os"
	"strings"
)

type Options struct {
	IgnoreCase bool
	Invert     bool
	LineNumber bool
	Patterns   []string
	Files      []string
}

func help() {
	fmt.Println("lorem ipsum")
}

func ParseArgs(argv []string) (*Options, error) {
	opts := &Options{}
	var positionals []string

	i := 0
	for i < len(argv) {
		token := argv[i]

		if token == "--" {
			positionals = append(positionals, argv[i+1:]...)
			break
		}

		if len(token) > 1 && token[0] == '-' {
			if err := parseFlags(token, opts); err != nil {
				return nil, err
			}
		} else {
			positionals = append(positionals, token)
		}
		i++
	}

	if len(positionals) == 0 {
		return nil, fmt.Errorf("no pattern given")
	}
	opts.Patterns = []string{positionals[0]}
	opts.Files = positionals[1:]

	return opts, nil
}

func parseFlags(token string, opts *Options) error {
	// -- format
	if !strings.HasPrefix(token, "-") {
		return fmt.Errorf("not a flag: %s", token)
	}
	if strings.HasPrefix(token, "--") && len(token) > 2 {
		name := token[2:]

		switch name {
		case "ignore-case":
			opts.IgnoreCase = true
		case "line-number":
			opts.LineNumber = true
		case "invert-match":
			opts.Invert = true
		case "help":
			help()
		default:
			return fmt.Errorf("unknown flag: --%s", name)
		}
		return nil
	}
	// - format
	body := token[1:]
	if body == "" {
		return fmt.Errorf("unknown flag: -")
	}
	for flagIdx := 0; flagIdx < len(body); flagIdx++ {
		switch body[flagIdx] {
		case 'i':
			opts.IgnoreCase = true
		case 'n':
			opts.LineNumber = true
		case 'v':
			opts.Invert = true
		default:
			return fmt.Errorf("unknown flag: -%c", body[flagIdx])
		}
	}
	return nil
}

func main() {
	opts, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", *opts)

}
