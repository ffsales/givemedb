/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	ContainerName string
	DatabaseName  string
)

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

	givemeCmd.PersistentFlags().StringVarP(&ContainerName, "container-name", "c", "", "Container's name")
	givemeCmd.MarkPersistentFlagRequired("container-name")

	givemeCmd.PersistentFlags().StringVarP(&DbName, "db-name", "d", "", "Database's name")
	givemeCmd.MarkPersistentFlagRequired("db-name")

	rootCmd.AddCommand(givemeCmd)
}
