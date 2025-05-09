/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rediseb/db-comp/config"
	"github.com/rediseb/db-comp/db"
	"github.com/spf13/cobra"
	"log"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbConfigs, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		for dbKey, dbConfig := range dbConfigs {
			fmt.Printf("Validating connection to %s...\n", dbKey)
			_, err := db.Connect(dbConfig)
			if err != nil {
				log.Printf("Error connecting to database %s: %v", dbKey, err)
			} else {
				fmt.Printf("Database connection to %s is valid!\n", dbKey)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
