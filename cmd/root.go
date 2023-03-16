/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	Verbose     bool   = false
	Limit       int    = 64
	Version     string = ""
	Compilacion string = ""
	Fecha       string = ""
	Autor       string = "Carlos Peña"
	Pagina      string = "https://code-epic.com"
	rootCmd            = &cobra.Command{
		Use:   "sandra_cli",
		Short: "Sandra Server Command Line Interface",
		Long: `Sandra Server Enterprise es una plataforma para arquitecturas empresariales 
que estará operativa para su organización ampliando la relación entre sistemas 
de diferentes naturaleza, así como la colaboración entre sus tecnologías adyacentes
`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetBool("version")

			if version {
				Ver()
			}

		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sandra_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Version del sistema")

}

func Ver() {
	fmt.Println("")
	myF := figure.NewColorFigure("Sandra CLI", "", "blue", true)
	myF.Print()

	fmt.Println("")
	fmt.Println("Command Line Interface")
	fmt.Println("Version: ", Version)
	fmt.Println("Fecha: ", Fecha)
	fmt.Println("Compilacion: ", Compilacion)
	fmt.Println("")
	fmt.Println("Autor: ", Autor)
	fmt.Println("Pagina: ", Pagina)
	fmt.Println("")
	fmt.Println("")

}
