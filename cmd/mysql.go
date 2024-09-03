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
	UserPass string
	DbName   string
	RootPass string
	UserName string
)

type MysqlParam struct {
	ContainerName string
	DatabaseName  string
	RootPass      string
	UserName      string
	UserPass      string
}

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

	mysqlCmd.PersistentFlags().StringVarP(&RootPass, "root-pass", "r", "", "User Root's password")
	mysqlCmd.MarkPersistentFlagRequired("root-pass")

	mysqlCmd.PersistentFlags().StringVarP(&UserPass, "user-pass", "p", "", "User's password")
	mysqlCmd.MarkPersistentFlagRequired("user-pass")

	mysqlCmd.PersistentFlags().StringVarP(&UserName, "user-name", "u", "", "UserName")
	mysqlCmd.MarkPersistentFlagRequired("user-name")

	givemeCmd.AddCommand(mysqlCmd)
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
