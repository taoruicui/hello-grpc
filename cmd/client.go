package cmd

import (
	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "serve",
	Run: func(cmd *cobra.Command, args []string) {
		//ctx := context.Background()
		//s, _ := NewClient()
		//Run(ctx, s, "8888")
	},
}