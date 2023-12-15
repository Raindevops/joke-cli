/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"tools-cli/internal/app/dadjoke"

	"github.com/spf13/cobra"
)

var dadjokeCmd = &cobra.Command{
	Use:   "dadjoke",
	Short: "Get a random dad joke",
	Long:  `This command fetches a random dadjoke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")
		if jokeTerm != "" {
			dadjoke.GetRandomJokeWithTerm(jokeTerm)
		} else {
			dadjoke.GetRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(dadjokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dadjokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dadjokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
