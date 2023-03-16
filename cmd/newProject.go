/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newProjectCmd represents the newProject command
var (
	Type          string
	Name          string
	Language      string
	Author        string
	newProjectCmd = &cobra.Command{
		Use:   "newProject [create|update] ",
		Short: "Crear un nuevo proyecto web, movil, desktop",
		Long: `Desarrolla aplicaciones en el contexto de un espacio de trabajo de Angular y Dart. 
Un espacio de trabajo contiene los archivos de uno o más proyectos son el conjunto de 
archivos que componen una aplicación o una biblioteca.`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 0 {
				CrearProyecto()
			} else {
				PrintColor(CRed, Help("newProject"))
			}
		},
		TraverseChildren: true,
	}
)

func init() {
	rootCmd.AddCommand(newProjectCmd)

	newProjectCmd.PersistentFlags().StringVarP(&Type, "type", "t", "web", `Tipo: web, movil, desktop`)
	newProjectCmd.PersistentFlags().StringVarP(&Name, "name", "n", "sandra_dev", `Nombre descriptivo`)
	newProjectCmd.PersistentFlags().StringVarP(&Author, "author", "a", "code.epic", `Autor del desarrollo`)
	newProjectCmd.PersistentFlags().StringVarP(&Language, "language", "l", "angular", `Lenguaje: angular, flutter `)

}

func CrearProyecto() {

}
