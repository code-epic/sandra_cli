/**
* Copyright © 2023 Carlos Enrique Pena gesaodin
**/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var (
	Seleccion   string
	Opciones    string
	Host        string
	Port        string
	User        string
	Nombre      string
	Password    string
	APPS               = []string{"mysql", "mongosh", "openssl", "curl", "gcc", "git"}
	SANDRA_HOME string = BASE_PATH + "/" + DIR
	SANDRA_BIN  string = SANDRA_HOME + "/bin"
	SANDRA_WWW  string = SANDRA_HOME + "/public_web"
	Spinner     Cargador
	installCmd  = &cobra.Command{
		Use:   "install ",
		Short: "Instalar herramientas, comandos y servcios",
		Long: `Una herramienta para verificar e instalar Sandra Server

Ayuda a evaluar los paquetes que posee instalado y sistema operativo.
e instala el servidor por primera vez`,
		Run: func(cmd *cobra.Command, args []string) {
			opt, _ := cmd.Flags().GetString("option")

			if opt != "" {
				ValidarArgumentos(opt)
			} else {
				PrintColor(CRed, Help("install"))
			}
			Verbose, _ = cmd.Flags().GetBool("verbose")
		},
		TraverseChildren: true,
	}
	Dw Descargar
)

func Help(cmd string) string {
	return `Intente escribir sandra ` + cmd + ` -h ver en pantalla las opciones.`
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("verbose", "v", false, "Mostrar detalles")

	installCmd.PersistentFlags().StringVarP(&Host, "host", "H", "", `Host del servidor`)
	installCmd.PersistentFlags().StringVarP(&Port, "port", "P", "", `Puerto`)
	installCmd.PersistentFlags().StringVarP(&User, "user", "U", "", `Usuario Base de Datos`)
	installCmd.PersistentFlags().StringVarP(&Nombre, "name", "N", "", `Nombdre de la Base de Datos`)
	installCmd.PersistentFlags().StringVarP(&Password, "passwd", "W", "", `Clave Base de Datos`)

	installCmd.Flags().BoolP("config", "c", false, "Configurar archivo de reglas")
	installCmd.PersistentFlags().StringVarP(&Opciones, "option", "o", "", `Escribir la opcion que se desea para Sandra Server ejemplo:

service: Incluye consola, demonio y base de datos internas
tools: Herramientas para el escaneo y analisis
data-base: Permite instalar base de datos de terceros
data-base-wkf: Permite instalar base de datos WorkFlow
mkcert: Crear certificados ssl https://mkcert.org
`)

}

func ValidarArgumentos(valor string) {
	switch valor {
	case "service":
		VerificarPaquetes()
	case "tools":
	case "data-base-wkf":
		CrearMysqlDBWKF(Host, User, Password, Nombre)
	case "data-base":
		InstalarBaseDatos()
	case "mkcert":
		MkCert()
	default:
		fmt.Println("Intente selecionar --help para mas informacion")
		return
	}

}

func VerificarPaquetes() {
	cant := len(APPS)
	var sCommand string
	fmt.Println("")
	Spinner.Start("Verificando paquetes")
	for i := 0; i < cant; i++ {
		v := APPS[i]
		sCommand += `which ` + v + ` 1>/dev/null 2>/dev/null || 
		echo -e "[-] Verifique que tenga instalado ` + v + ` para continuar ";
		# sleep 1;
		`
	}
	cmd := exec.Command("bash", "-c", sCommand)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando los comando de verificacion")
		return
	}
	Spinner.Stop("[+] Los paquetes han sido verificados correctamente")
	fmt.Println(string(out))

	if string(out) != "" {
		PrintColor(CRed, "- El Proceso ha terminado. Debe instalar todos los paquetes para continuar... ")
		return
	}
	InstalarServicio()
}

// Instalar Servicios Generales
func InstalarServicio() {

	fileName := Dw.App(DW_SANDRA)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("   Configurando la plataforma Sandra Server              ")
	fmt.Println("   " + SANDRA_HOME)
	fmt.Println("---------------------------------------------------------")
	sCmd := `#!/bin/bash
# Creando enlace para los servicios al sistema
DIR=$(pwd)
BASE_URL="` + SANDRA_HOME + `";
RUTA=$HOME/.bashrc;

mkdir -p ` + SANDRA_HOME + `;

# Creando la base para los paquete /usr/local/bin/sandra
echo -e "[+] Creando espacio de trabajo en $BASE_URL"
unzip ` + fileName + ` 1>/dev/null && echo -e "    - El paquete ha sido descomprimido"
rm -rf ` + fileName + `
mv sandra/* $BASE_URL 2>/dev/null;
echo -e "    - Moviendo al area de trabajo"`

	ExecCmd(sCmd)
	InstalarBaseDatos()

	CrearVariables()
	PrintColor(CGreen, "[*] Proceso finalizado con exito")
}

// InstalarBaseDatos Instalar esquema de base de datos
func InstalarBaseDatos() {
	fmt.Println("[+] Esquema de datos Mysql... ")
	mysql := Dw.App(DW_MYSQLDB)
	fmt.Println("    ---------------------------------------------------------")
	fmt.Println("     Configurando base de datos de mysql para            ")
	fmt.Println("     " + mysql)
	fmt.Println("    ---------------------------------------------------------")
	ip_mysql, us_mysql, pw_mysql, _ := credentials()
	fmt.Println("")
	fmt.Println("[+] Preparando la instalacion mysql, por favor espere...")
	CrearMysqlDB(ip_mysql, us_mysql, pw_mysql, mysql)
	fmt.Println("")
	fmt.Println("[+] Esquema de datos MongoDB... ")
	mgo := Dw.App(DW_MONGODB)
	fmt.Println("    ---------------------------------------------------------")
	fmt.Println("     Configurando base de datos de mongo para            ")
	fmt.Println("     " + mgo)
	fmt.Println("    ---------------------------------------------------------")
	ip_mgo, us_mgo, pw_mgo, _ := credentials()
	fmt.Println("")
	fmt.Println("[+] Preparando la instalacion mongodb, por favor espere...")
	CrearMongoDB(ip_mgo, us_mgo, pw_mgo, mgo)
	fmt.Println("")
	EliminarTemporales(mysql, mgo)
	fmt.Println("")
}

// CrearMysqlDump Crear Base de datos code-epic
func CrearMysqlDB(ip string, user string, pass string, my_dbname string) {

	if ip != "" {
		ip = "-h " + ip
	}

	sCmd := `
echo -e "[+] Creando Base de datos "
unzip ` + my_dbname + ` 1>/dev/null && echo -e "    - El paquete ` + my_dbname + ` se ha descomprimido ";
echo -e "    - Archivo descomprimido: ";
mysql -u` + user + ` -p` + pass + ` ` + ip + ` -e "CREATE DATABASE ` +
		DB_NAME + ` CHARACTER SET utf8mb4 COLLATE utf8mb4_spanish_ci;";
echo -e "    - Base de datos (` + DB_NAME + `) creada ";`
	ExecCmd(sCmd)

	sCmd = `mysql -u` + user + ` -p` + pass + ` ` + ip + ` ` + DB_NAME + ` < mysqldb/code_epic.sql 2>/dev/null ;
echo -e "    - Estructura de seguridad creada";`
	ExecCmd(sCmd)

	sCmd = `
echo -e "[+] Creando Base de datos Workflow"
mysql -u` + user + ` -p` + pass + ` ` + ip + ` -e "CREATE DATABASE ` +
		DB_WKF + ` CHARACTER SET utf8mb4 COLLATE utf8mb4_spanish_ci;";
echo -e "    - Base de datos (` + DB_WKF + `) creada ";`
	ExecCmd(sCmd)

	sCmd = `mysql -u` + user + ` -p` + pass + ` ` + ip + ` ` + DB_WKF + ` < mysqldb/wkf.sql 2>/dev/null;
echo -e "    - Estructura de workflow creada";`
	ExecCmd(sCmd)

}

// CrearMysqlDump Crear Base de datos code-epic
// cargar dump de Workflow
func CrearMysqlDBWKF(ip string, user string, pass string, my_dbname string) {

	fmt.Println("[+] Esquema de datos Mysql... ")
	mysql := Dw.App(DW_MYSQLDB)
	fmt.Println("    ---------------------------------------------------------")
	fmt.Println("     Configurando base de datos de mysql para            ")
	fmt.Println("     " + mysql)
	fmt.Println("    ---------------------------------------------------------")
	if ip != "" {
		ip = "-h " + ip
	}
	sCmd := `
echo -e "[+] Creando Base de datos "
unzip ` + mysql + ` 1>/dev/null && echo -e "    - El paquete ` + mysql + ` se ha descomprimido ";
echo -e "    - Archivo descomprimido: ";
mysql -u` + user + ` -p` + pass + ` ` + ip + ` -e "CREATE DATABASE ` + my_dbname + ` CHARACTER SET utf8mb4 COLLATE utf8mb4_spanish_ci;";
echo -e "    - Base de datos (` + my_dbname + `) creada ";`
	ExecCmd(sCmd)

	sCmd = `mysql -u` + user + ` -p` + pass + ` ` + ip + ` ` + my_dbname + ` < mysqldb/wkf.sql 2>/dev/null ;
echo -e "    - Estructura de seguridad creada";`
	ExecCmd(sCmd)

}

func CrearMongoDB(ip string, user string, pass string, my_dbname string) {

	sCmd := `
echo -e "[+] Creando Base de datos "
unzip ` + my_dbname + ` 1>/dev/null && echo -e "    - El paquete ` + my_dbname + ` se ha descomprimido ";
echo -e "    - Archivo descomprimido: ";
mongorestore --db ` + DB_MONGO + ` dump/code-epic 2>/dev/null;
echo -e "    - Base de datos (` + DB_MONGO + `) creada ";`
	ExecCmd(sCmd)
}

func EliminarTemporales(mysql string, mgo string) {
	sCmd := `echo "[+] Proceeso para eliminar temporales"
rm -rf sandra
echo -e "    - sandra eliminado"
rm -rf ` + mysql + `;
rm -rf mysqldb;
echo -e "    - mysqldb eliminado"
rm -rf ` + mgo + `;
rm -rf dump;
echo -e "    - mongodb eliminados";`
	ExecCmd(sCmd)

}

func CrearVariables() {
	sCmd := `#!/bin/bash
# Creando enlace para los servicios al sistema
DIR=$(pwd)
BASE_URL="` + SANDRA_HOME + `";
RUTA=$HOME/.bashrc;

cp sandra_cli $BASE_URL/bin

echo -e "[+] Copiando archivo del servicio \n"
cp $BASE_URL/cmd/sandrad.service ` + SERVICE + `/sandrad.service

ln -s $BASE_URL/bin/sandra_cli /usr/bin/sandra
ln -s $BASE_URL/bin/sandra_dwn /usr/bin/sandra_dwn
ln -s $BASE_URL/bin/sandra_scanf /usr/bin/sandra_scanf

echo -e "[+] Cambiando el contexto SELinux del demonio $BASE_URL/sandrad"
semanage fcontext -a -t bin_t $BASE_URL/sandrad
echo -e "    - Finalizo el proceso de cambio "
restorecon -vF $BASE_URL/sandrad
echo -e "    - Restaurando el contexto de seguridad "

echo -e "[+] Declarando variables de entorno al PATH"
echo -e "export SANDRA_HOME=$BASE_URL" >> $RUTA
echo -e "export SANDRA_BIN=$BASE_URL/bin" >> $RUTA
echo "export PATH=$PATH:$BASE_URL:$BASE_URL/bin" >> $RUTA

source $HOME/.bashrc;
systemctl daemon-reload
systemctl reset-failed
systemctl enable sandrad
systemctl start sandrad`
	ExecCmd(sCmd)
}

// MkCert Descargar e Instalar los certificados
func MkCert() {

	mkname := Dw.App(DW_MKCERT)

	sCmd := `#!/bin/bash
# Creando enlace para los servicios al sistema
BASE_URL="` + SANDRA_HOME + `";

chmod +x ` + mkname + `
cp mkcert-v*-linux-amd64 /usr/local/bin/mkcert
mkcert -install 2>/dev/null
mkcert -key-file $BASE_URL/signure/sandra.app.key -cert-file $BASE_URL/signure/sandra.app.crt localhost
`
	ExecCmd(sCmd)

}
