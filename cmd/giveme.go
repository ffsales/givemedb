/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// givemeCmd represents the giveme command
var givemeCmd = &cobra.Command{
	Use:   "giveme",
	Short: "Create a container of DB of type MySQL or Postgres",
	Long: `This command allow to create a contaner Docker of image MySQSL or Postgres, for example:

	givemedb-cli giveme mysql
	givemedb-cli giveme postgre

This application was build using Cobra framework.`,
	SuggestFor: []string{"mysql", "postgre"},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return []string{"mysql", "postgre"}, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(givemeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// givemeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// givemeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
