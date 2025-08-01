package cmd

import (
	"fmt"

	"github.com/imzadev/nudl/internal/api"
	"github.com/imzadev/nudl/internal/jsonloader"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "new",
	Short: "generate new rest api",
	Long: `generate new rest api from json file with your own data.
	For example :
	nudl new -p 3000 -j home/.../data.json`,

	Run: generateApi,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("port", "p", "3000", "api server port")
	generateCmd.Flags().StringP("path", "j", "", "json file path")
}

func generateApi(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetString("port")
	path, _ := cmd.Flags().GetString("path")

	logo := `
 _   _ _   _ ____  _    
| \ | | | | |  _ \| |   
|  \| | | | | | | | |   
| |\  | |_| | |_| | |___
|_| \_|\___/|____/|_____|
                        
    JSON → API on TUI
`

	Data, err := jsonloader.New(path).LoadAndParse()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(logo)
		fmt.Println("================================================")
		fmt.Printf("Starting Api on :\nhttp://127.0.0.1:%s\n", port)
		fmt.Printf("================================================\n\n")
		api.New(port, Data).Start()
	}

}
