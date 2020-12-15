package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-chi/docgen"
	"github.com/spf13/cobra"
	"github.com/thomasxnguy/bitcoinaddress/api"
)

var (
	routes bool
)

// gendocCmd represents the gendoc command
var gendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "Generate project documentation",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if routes {
			genRoutesDoc()
		}
	},
}

func init() {
	RootCmd.AddCommand(gendocCmd)
	gendocCmd.Flags().BoolVarP(&routes, "routes", "r", false, "create api routes markdown file")
}

func genRoutesDoc() {
	api, _ := api.New(false)
	fmt.Print("generating routes markdown file: ")
	md := docgen.MarkdownRoutesDoc(api, docgen.MarkdownOpts{
		ProjectPath: "github.com/thomasxnguy/bitcoinaddress",
		Intro:       "HTTP server to generate bitcoin addresses",
	})
	if err := ioutil.WriteFile("routes.md", []byte(md), 0644); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("OK")
}
