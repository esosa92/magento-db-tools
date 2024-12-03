package main

import "github.com/spf13/cobra"

var dump_core_config_data_cmd = &cobra.Command{
	Use:   "dump-core-config-data",
	Short: "When run provide path to env file and it will dump the core_config_data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
