package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/yourusername/subenum/internal"
)

var rootCmd = &cobra.Command{
    Use:   "subenum",
    Short: "Subdomain enumeration CLI tool",
    Run: func(cmd *cobra.Command, args []string) {
        err := internal.RunWorkflow()
        if err != nil {
            fmt.Println("❌ Error:", err)
            os.Exit(1)
        }
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}
