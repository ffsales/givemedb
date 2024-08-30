/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	ContainerName string
	UserPass      string
	DbName        string
	RootPass      string
	UserName      string
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

	givemeCmd.PersistentFlags().StringVarP(&ContainerName, "container-name", "c", "", "Container's name")
	givemeCmd.MarkPersistentFlagRequired("container-name")

	givemeCmd.PersistentFlags().StringVarP(&DbName, "db-name", "d", "", "Database's name")
	givemeCmd.MarkPersistentFlagRequired("db-name")

	givemeCmd.PersistentFlags().StringVarP(&RootPass, "root-pass", "r", "", "User Root's password")
	givemeCmd.MarkPersistentFlagRequired("root-pass")

	givemeCmd.PersistentFlags().StringVarP(&UserPass, "user-pass", "p", "", "User's password")
	givemeCmd.MarkPersistentFlagRequired("user-pass")

	givemeCmd.PersistentFlags().StringVarP(&UserName, "user-name", "u", "", "UserName")
	givemeCmd.MarkPersistentFlagRequired("user-name")

	rootCmd.AddCommand(givemeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// givemeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// givemeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
