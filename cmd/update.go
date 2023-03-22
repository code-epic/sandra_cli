/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var (
	Update    Descargar
	updateCmd = &cobra.Command{
		Use:   "update ",
		Short: "Actualizar servicios Sandra Server",
		Long: `Permite actualizar servicio tales como: 
el demonio de sandrad, las variables de entorno, paquetes, 
API internas de Sandra Server`,
		Run: func(cmd *cobra.Command, args []string) {
			opt, _ := cmd.Flags().GetString("option")

			if opt != "" {
				ValidarParametros(opt)
			} else {
				PrintColor(CRed, Help("update"))
			}
			Verbose, _ = cmd.Flags().GetBool("verbose")
		},
		TraverseChildren: true,
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolP("verbose", "v", false, "Mostrar detalles")
	updateCmd.Flags().BoolP("config", "c", false, "Configurar archivo de reglas")
	updateCmd.PersistentFlags().StringVarP(&Opciones, "option", "o", "", `Escribir la opcion que se desea para Sandra Server ejemplo:

service: Actualiza la consola y el demonio de servicio
tools: Herramientas para el escaneo y analisis
data-base: Permite actualizar base de datos de terceros
`)

}

func ValidarParametros(valor string) {
	switch valor {
	case "service":
		UpdateService()
	case "tools":
		UpdateTools()
	case "consola":
		UpdateConsola()
	case "cli":
		UpdateCli()
	case "data-base":
		UpdateDataBase()
	default:
		fmt.Println("Intente selecionar --help para mas informacion")
		return
	}

}

// UpdateService Actualizar servicio sandrad
func UpdateService() {

	sCmd := `
clear
echo -e "[+] Deteniendo el servicio de sandrad.service"
systemctl stop sandrad`
	fileName := Update.App(DW_SANDRA_DAEMON)

	ExecCmd(sCmd)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("     Actualizando la plataforma Sandra Server            ")
	fmt.Println("     " + SANDRA_HOME)
	fmt.Println("---------------------------------------------------------")
	sCmd = `#!/bin/bash
BASE_URL="` + SANDRA_HOME + `";
echo -e "[+] Descomprimiendo archivo"
unzip ` + fileName + ` 1>/dev/null && echo -e "    - El paquete ha sido descomprimido"
cp sandra/* $BASE_URL 2>/dev/null;
echo -e "    - Actualizando sandrad en $BASE_URL"
echo -e "    - Eliminando temporales"
rm -rf ` + fileName + `
rm -rf sandra
echo -e "[+] Iniciando el servicio sandrad"
systemctl start sandrad
`
	ExecCmd(sCmd)

}

// UpdateTools Herramienta de actualizacion de comandos del Sandra Server
func UpdateTools() {

	sCmd := `clear;echo -e "[+] Actualizando herramientas de comandos ` + SANDRA_BIN + `"`
	ExecCmd(sCmd)
	fileName := Update.App(DW_SANDRA_TOOLS)

	fmt.Println("---------------------------------------------------------")
	fmt.Println("   Actualizando las herramientas Sandra Server           ")
	fmt.Println("   " + SANDRA_BIN)
	fmt.Println("---------------------------------------------------------")
	sCmd = `#!/bin/bash
BASE_URL="` + SANDRA_BIN + `";
echo -e "[+] Descomprimiendo archivo"
unzip ` + fileName + ` 1>/dev/null && echo -e "    - El paquete ha sido descomprimido"
cp tools/* $BASE_URL 2>/dev/null;
echo -e "    - Actualizando sandrad een $BASE_URL"
echo -e "    - Eliminando temporales"
rm -rf tools
rm -rf ` + fileName + `
echo -e "[+] Proceso finalizado con exito..."`

	ExecCmd(sCmd)
}

// UpdateDataBase Actualizando las Base de Datos
func UpdateDataBase() {

}

// UpdateCli Actualizar servicio sandrad
func UpdateCli() {

	fileName := Update.App(DW_SANDRA_CLI)
	sCmd := `clear`
	ExecCmd(sCmd)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("     Actualizando command line interface                 ")
	fmt.Println("     " + SANDRA_BIN)
	fmt.Println("---------------------------------------------------------")
	sCmd = `#!/bin/bash
BASE_URL="` + SANDRA_BIN + `";
echo -e "[+] Descomprimiendo archivo"
unzip ` + fileName + ` 1>/dev/null && echo -e "    - El paquete ha sido descomprimido"
rm -rf ` + fileName + `
cp sandra_cli $BASE_URL 2>/dev/null;
echo -e "    - Actualizando sandra_cli en $BASE_URL"
echo -e "    - Eliminando temporales"
rm -rf sandra_cli
`
	ExecCmd(sCmd)
}

func UpdateConsola() {
	fileName := Update.App(DW_SANDRA_CONSOLA)
	sCmd := `clear`
	ExecCmd(sCmd)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("     Actualizando Interfaz grafica (Consola)             ")
	fmt.Println("     " + SANDRA_WWW + "/consola")
	fmt.Println("---------------------------------------------------------")
	sCmd = `#!/bin/bash
BASE_URL="` + SANDRA_WWW + `/consola";
echo -e "[+] Descomprimiendo archivo"
unzip ` + fileName + ` 1>/dev/null && echo -e "    - El paquete ha sido descomprimido"
rm -rf ` + fileName + `
cp -r sandra/public_web/consola/* $BASE_URL 2>/dev/null;
rm -rf sandra
echo -e "[+] Eliminando archivos temporales"
echo -e "[+] Version $(sed -n '2,1 p' $BASE_URL/assets/version.sdr) Instalada"
`
	ExecCmd(sCmd)

}
