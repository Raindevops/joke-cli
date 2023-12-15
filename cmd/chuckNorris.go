/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	chucknorris "tools-cli/internal/app/chuckNorris"

	"github.com/spf13/cobra"
)

// chuckNorrisCmd represents the chuckNorris command
var chuckNorrisCmd = &cobra.Command{
	Use:   "chuckNorris",
	Short: "Print a Chuck Norris phrase",
	Long:  `Print a Chuck Norris phrase using the api https://api.chucknorris.io/`,
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		list, _ := cmd.Flags().GetBool("list-categories")
		if category != "" {
			chucknorris.GetPhraseByCategory()
		} else if list {
			chucknorris.ListAllCategories()
		} else {
			chucknorris.RandomPhrase()
		}
	},
}

func init() {
	rootCmd.AddCommand(chuckNorrisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chuckNorrisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	chuckNorrisCmd.PersistentFlags().String("category", "", "Find a Chuck Norris phrase by category")
	chuckNorrisCmd.Flags().BoolP("list-categories", "", false, "List all categories availabled for Chuck Norris phrases")
}
