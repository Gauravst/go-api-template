package middleware

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CheckAuth() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		token := os.Getenv("AUTH_TOKEN")
		if token == "" {
			fmt.Println("Error: AUTH_TOKEN not set. Please login first.")
			os.Exit(1)
		}
		fmt.Println("Authorization check passed")
	}
}

func CheckSomething() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Your middleware logic here
	}
}
