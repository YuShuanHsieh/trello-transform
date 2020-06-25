package main

import (
	"log"
	"os"
	"time"

	transform "github.com/YuShuanHsieh/trello-transform"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

const (
	cmdFlagKey     = "key"
	cmdFlagToken   = "token"
	cmdFlagBoardID = "id"
	cmdFlagList    = "list"
	cmdFlagTarget  = "target"
	cmdFlagDebug   = "debug"
)

func CmdAction() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		debug := ctx.Bool(cmdFlagDebug)
		key := ctx.String(cmdFlagKey)
		token := ctx.String(cmdFlagToken)
		id := ctx.String(cmdFlagBoardID)
		lists := ctx.StringSlice(cmdFlagList)
		targets := ctx.StringSlice(cmdFlagTarget)

		var config zap.Config

		if debug {
			config = zap.NewDevelopmentConfig()
		} else {
			config = zap.NewProductionConfig()
		}
		config.DisableCaller = true
		config.DisableStacktrace = true

		logger, _ := config.Build()

		logger.Debug("setup a client for trello API",
			zap.String("board-id", id),
		)

		trans := transform.New(logger, key, token, id)

		trans.AddSelect(transform.SelectByListNames(lists...))

		logger.Debug("setup targets",
			zap.Strings("targets", targets),
		)

		for _, target := range targets {
			switch target {
			case "titles":
				trans.AddTransformFunc(transform.ToTitles)
			case "links":
				trans.AddTransformFunc(transform.ToLinks)
			}
		}
		return trans.Exec()
	}
}

func main() {
	app := cli.App{
		Name:        "trellotrans",
		Usage:       "transform card's info of trello board to plain text",
		Description: "transform card's info of trello board to plain text",
		Commands:    nil,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  cmdFlagDebug,
				Usage: "open debug mode",
			},
			&cli.StringFlag{
				Required: true,
				Name:     cmdFlagKey,
				Value:    "",
				Usage:    "the api key of trello",
			},
			&cli.StringFlag{
				Required: true,
				Name:     cmdFlagToken,
				Value:    "",
				Usage:    "the develop token of trello",
			},
			&cli.StringFlag{
				Required: true,
				Name:     cmdFlagBoardID,
				Value:    "",
				Usage:    "the target trello board id",
			},
			&cli.StringSliceFlag{
				Required: true,
				Name:     cmdFlagList,
				Value:    nil,
				Usage:    "the name of a list which would be parsed",
			},
			&cli.StringSliceFlag{
				Required: true,
				Name:     cmdFlagTarget,
				Value:    nil,
				Usage:    "the target types of output",
			},
		},
		Action:               CmdAction(),
		EnableBashCompletion: true,
		Compiled:             time.Time{},
		Copyright:            "",
		Author:               "Cherie Hsieh",
		Email:                "cjamhe01385@gmail.com",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
