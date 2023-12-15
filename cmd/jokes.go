/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"joke-cli/internal/app/jokes"

	"github.com/spf13/cobra"
)

var jokesCmd = &cobra.Command{
	Use:   "jokes",
	Short: "Get a random joke",
	Long:  `This command fetch a random joke from the official joke API. See the github project for more informations : https://github.com/15Dkatz/official_joke_api `,
	Run: func(cmd *cobra.Command, args []string) {
		jokes.GetJoke()
	},
}

func init() {
	rootCmd.AddCommand(jokesCmd)
}
