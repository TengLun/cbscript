package main

import (
	"cbs"
	"errors"
	"fmt"
	"os"

	"github.com/tenglun/logger"
	"gopkg.in/urfave/cli.v2"
)

func main() {

	// Create file logger
	logger := logger.CreateLogger("log")

	var api string
	var filename string
	var debug bool
	var autoupdate bool

	app := &cli.App{
		Name:  "cbscript",
		Usage: "Add, Update, or Remove references from your accounts blacklist",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "api_key",
				Aliases:     []string{"ak"},
				Usage:       "security token for api calls",
				Destination: &api,
			},
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "file that contains blacklist entries",
				Destination: &filename,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "write to file instead of api",
				Destination: &debug,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add file references to blacklist",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "autoupdate",
						Aliases:     []string{"ap"},
						Usage:       "Force update if reference already present",
						Destination: &autoupdate,
					},
				},
				Action: func(c *cli.Context) error {

					if filename == "" {
						err := errors.New("filename must be specified. Run `cbscript help` to learn usage")
						fmt.Println(err)
						logger.Println(err)
						return err
					}

					if api == "" {
						err := errors.New("api key must be specified. Run `cbscript help` to learn usage")
						fmt.Println(err)
						logger.Println(err)
						return err
					}

					if autoupdate {
						err := cbs.Action(logger, filename, api, debug, "addupdate")
						if err != nil {
							logger.Println(err)
							return err
						}
						return nil
					}
					if !autoupdate {
						err := cbs.Action(logger, filename, api, debug, "add")
						if err != nil {
							logger.Println(err)
							return err
						}
						return nil
					}
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"up"},
				Usage:   "update file references on blacklist",
				Action: func(c *cli.Context) error {

					if filename == "" {
						err := errors.New("filename must be specified. Run `cbscript help` to learn usage")
						fmt.Println(err)
						logger.Println(err)
						return err
					}

					if api == "" {
						err := errors.New("api key must be specified. Run `cbscript help` to learn usage")
						fmt.Println(err)
						logger.Println(err)
						return err
					}

					err := cbs.Action(logger, filename, api, debug, "update")
					if err != nil {
						logger.Println(err)
						return err
					}
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "remove file references from the blacklist",
				Action: func(c *cli.Context) error {

					if filename == "" {
						err := errors.New("filename must be specified. Run `cbscript help` to learn usage")
						fmt.Println(err)
						logger.Println(err)
						return err
					}

					if api == "" {
						err := errors.New("api key must be specified. Run `cbscript help` to learn usage")
						fmt.Println(err)
						logger.Println(err)
						return err
					}

					err := cbs.Action(logger, filename, api, debug, "remove")
					if err != nil {
						logger.Println(err)
						return err
					}
					return nil
				},
			},
		},
	}

	app.Version = "0.1.1"
	app.Name = "Custom Blacklist Builder"
	app.Description = "A simple program to add, update, and remove entries from your accounts custom blacklist"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Teng Lun",
			Email: "fivemilesgone@protonmail.com",
		},
	}
	app.Run(os.Args)
}
