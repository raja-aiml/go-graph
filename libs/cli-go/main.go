package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "langgraph",
		Short: "LangGraph CLI placeholder",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("LangGraph CLI")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
