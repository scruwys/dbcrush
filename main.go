package main

import (
	"github.com/spf13/cobra"
)

var flagHostname string
var flagUsername string
var flagPassword string
var flagPort int
var flagDatabase string
var flagSeed int

var cmd = &cobra.Command{
	Use:   "dbcrush <path-to-file>",
	Short: "dbcrush is a CLI for simulating database traffic.",
	Args:  cobra.ExactArgs(1),
	Run:   Run,
}

func main() {
	cmd.PersistentFlags().StringVar(
		&flagHostname,
		"host",
		"localhost",
		"database server host or socket directory")

	cmd.PersistentFlags().IntVarP(
		&flagPort,
		"port",
		"p",
		3306,
		"database server port")

	cmd.PersistentFlags().StringVarP(
		&flagUsername,
		"username",
		"U",
		"admin",
		"database user name")

	cmd.PersistentFlags().StringVar(
		&flagPassword,
		"password",
		"",
		"force password prompt (should happen automatically)")

	cmd.PersistentFlags().StringVarP(
		&flagDatabase,
		"dbname",
		"d",
		"default",
		"database user name")

	cmd.PersistentFlags().IntVarP(
		&flagSeed,
		"seed",
		"s",
		0,
		"value needed to generate a random number")

	cmd.Execute()
}
