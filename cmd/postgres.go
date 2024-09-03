/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ffsales/givemedb/service/docker"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	PostgrePass string
)

type PostgresParams struct {
	ContainerName string
	DatabaseName  string
	RootPass      string
}

var postgreCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Set o Postgre how DB to be created",
	RunE: func(cmd *cobra.Command, args []string) error {

		params, error := getParametersPostgres(cmd.Flags())
		if error != nil {
			log.Fatalf("Erro ao iniciar container: %v", error)
		}

		docker.CreatePostgres(params.ContainerName, params.DatabaseName, params.RootPass)

		fmt.Println("Postgres Created:", params.RootPass)

		return nil
	},
}

func init() {
	postgreCmd.PersistentFlags().StringVarP(&PostgrePass, "postgres-pass", "g", "", "Root pass Postgress")
	postgreCmd.MarkPersistentFlagRequired("postgres-pass")

	givemeCmd.AddCommand(postgreCmd)
}

func getParametersPostgres(flags *pflag.FlagSet) (PostgresParams, error) {
	pass, err := flags.GetString("postgres-pass")
	containerName, err := flags.GetString("container-name")
	dbName, err := flags.GetString("db-name")

	if err != nil {
		fmt.Printf("Error retrieving flag: %s\n", err.Error())
		return PostgresParams{}, err
	}

	return PostgresParams{
		ContainerName: containerName,
		RootPass:      pass,
		DatabaseName:  dbName,
	}, err
}
