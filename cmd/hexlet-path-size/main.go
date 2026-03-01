package main

import (
	"context"
	"fmt"
	"log"
	"os"

	code "github.com/bkoshelev/go-project-242"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "path",
				Value:     "",
				UsageText: "path to file or directory",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "recursive size of directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.StringArg("path")

			if path == "" {
				fmt.Printf("you need to write path to file or directory")
				return nil
			}

			result, err := code.GetPathSize(path, cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\t%s", result, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
