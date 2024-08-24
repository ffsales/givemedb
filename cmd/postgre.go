/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// postgreCmd represents the postgre command
var postgreCmd = &cobra.Command{
	Use:   "postgre",
	Short: "Set o Postgre how DB to be created",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Postgres Created")

		return nil
	},
}

func init() {
	givemeCmd.AddCommand(postgreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postgreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postgreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
