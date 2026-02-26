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
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path> - path to file or directory",
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
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.StringArg("path")

			if path == "" {
				fmt.Printf("you need to write path to file or directory")
				return nil
			}

			result, err := code.GetPathSize(path, cmd.Bool("human"))

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
