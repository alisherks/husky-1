package main

import (
	"context"
	"flag"
	"os"
	"path"

	"github.com/go-courier/husky/pkg/husky"
	"github.com/go-courier/husky/pkg/log"
	"github.com/go-courier/husky/version"
	"github.com/go-logr/glogr"
	"github.com/spf13/cobra"
)

var (
	logger      = glogr.New().WithName("husky")
	projectRoot = husky.ResolveGitRoot()
	theHusky    = husky.HuskyFrom(log.WithLogger(logger)(context.Background()), path.Join(projectRoot, ".husky.yaml"))
)

var CmdRoot = &cobra.Command{
	Use:     "husky",
	Short:   "husky " + version.Version,
	Version: version.Version,
}

func init() {
	flag.Parse()
	CmdRoot.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	Init(projectRoot)
}

func Init(root string) {
	githooks, _ := husky.ListGithookName(root)

	for _, githook := range githooks {
		_ = os.WriteFile(path.Join(root, ".git/hooks", githook), []byte(`#!/bin/sh

husky hook $(basename "$0") $*
`), os.ModePerm)
	}
}
