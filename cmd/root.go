package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)


var rootCmd = &cobra.Command{
    Use:   "nudl",
    Short: "nudl is a cli tool for transform json data into rest api",
    Long:  "nudl is a cli tool for transform json data into rest api",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing nudl '%s'\n", err)
        os.Exit(1)
    }
}
