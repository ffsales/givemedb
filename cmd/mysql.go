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

type MysqlParam struct {
	ContainerName string
	DatabaseName  string
	RootPass      string
	UserName      string
	UserPass      string
}

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Set o MySql how DB to be created",
	RunE: func(cmd *cobra.Command, args []string) error {

		params, error := getParameters(cmd.Flags())
		if error != nil {
			log.Fatalf("Erro ao iniciar container: %v", error)
		}

		docker.CreateMysql(params.ContainerName, params.DatabaseName, params.RootPass, params.UserPass, params.UserName)

		fmt.Println("ContainerName:" + params.ContainerName)
		return nil
	},
}

func init() {
	givemeCmd.AddCommand(mysqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mysqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getParameters(flags *pflag.FlagSet) (MysqlParam, error) {
	pass, err := flags.GetString("user-pass")
	mysqlUser, err := flags.GetString("user-name")
	rootPass, err := flags.GetString("root-pass")
	containerName, err := flags.GetString("container-name")
	dbName, err := flags.GetString("db-name")

	if err != nil {
		fmt.Printf("Error retrieving flag: %s\n", err.Error())
		return MysqlParam{}, err
	}

	return MysqlParam{
		ContainerName: containerName,
		DatabaseName:  dbName,
		RootPass:      rootPass,
		UserPass:      pass,
		UserName:      mysqlUser,
	}, err
}
