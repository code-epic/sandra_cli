package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
	"golang.org/x/term"
)

// Cargan
type Cargador struct {
	Sp *spinner.Spinner
}

type Descargar struct{}

type Color string

var ()

// Declarando parametros de la instalacion
const (
	TB                      = 1000000000000
	GB                      = 1000000000
	MB                      = 1000000
	KB                      = 1000
	layout           string = "2006-01-02"
	CBlack           Color  = "\u001b[30m"
	CRed                    = "\u001b[31m"
	CGreen                  = "\u001b[2m"
	CYellow                 = "\u001b[33m"
	CBlue                   = "\u001b[34m"
	CReset                  = "\u001b[0m"
	BASE_PATH        string = "/usr/local"
	SERVICE          string = "/etc/systemd/system"
	LOGS             string = "/var/log/sandra"
	DIR              string = "sandra"
	CMD              string = "cmd"
	DEAMON           string = "sandrad"
	BASE_REPO        string = "https://github.com/code-epic/sandra-enterprise/raw/main/"
	DW_MKCERT        string = "https://dl.filippo.io/mkcert/latest?for=linux/amd64"
	DW_SANDRA        string = BASE_REPO + "pkg/linux/x86_64/sandra.zip"
	DW_SANDRA_DAEMON string = BASE_REPO + "pkg/linux/x86_64/sandra_daemon.zip"
	DW_SANDRA_TOOLS  string = BASE_REPO + "pkg/linux/x86_64/sandra_tools.zip"
	DW_SANDRA_CLI    string = BASE_REPO + "pkg/linux/x86_64/sandra_cli.zip"
	DW_MONGODB       string = BASE_REPO + "db/mongo.sse.x86_64.zip"
	DW_MYSQLDB       string = BASE_REPO + "db/security.sse.x86_64.zip"
)

func PrintColor(color Color, message string) {
	fmt.Println(string(color), message, string(CReset))
}

func (D *Descargar) App(url string) (fileName string) {
	tokens := strings.Split(url, "/")
	fileName = tokens[len(tokens)-1]
	Spinner.Start("Descargando: " + fileName)
	if Verbose {
		fmt.Println(url)
	}

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	peso := PesoHumano(int(n), 2)
	Spinner.Stop("[+] " + peso + " descargados de " + fileName)
	fmt.Println("")
	return
}

func (D *Descargar) DB() {

}

// Start Iniciar cargador
func (C *Cargador) Start(msj string) {
	some := []string{"[ ◢ ] ", "[ ◣ ] ", "[ ◤ ] ", "[ ◥ ] "}
	C.Sp = spinner.New(some, 100*time.Millisecond)
	C.Sp.Suffix = msj
	err := C.Sp.Color("yellow", "bold")
	if err != nil {
		fmt.Println("Error : ", err)
		//return
	}
	C.Sp.Start()
}

// Stop finalizar cargando
func (C *Cargador) Stop(msj string) {
	C.Sp.FinalMSG = msj
	err := C.Sp.Color("green", "bold")
	if err != nil {
		fmt.Println("Error :", err)
		//return
	}
	C.Sp.Stop()
}

func credentials() (string, string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Introduzca el IP o Hostname default (localhost): ")
	Ip, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Introduzca el Usuario: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Introduzca el Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", "", err
	}
	password := string(bytePassword)

	sIP := strings.TrimSpace(Ip)

	return sIP, strings.TrimSpace(username), strings.TrimSpace(password), nil
}

// PesoHumano Valor del peso
func PesoHumano(length int, decimals int) (out string) {
	var unit string
	var i int
	var remainder int

	// Get whole number, and the remainder for decimals
	if length > TB {
		unit = "TB"
		i = length / TB
		remainder = length - (i * TB)
	} else if length > GB {
		unit = "GB"
		i = length / GB
		remainder = length - (i * GB)
	} else if length > MB {
		unit = "MB"
		i = length / MB
		remainder = length - (i * MB)
	} else if length > KB {
		unit = "KB"
		i = length / KB
		remainder = length - (i * KB)
	} else {
		return strconv.Itoa(length) + " B"
	}

	if decimals == 0 {
		return strconv.Itoa(i) + " " + unit
	}

	// This is to calculate missing leading zeroes
	width := 0
	if remainder > GB {
		width = 12
	} else if remainder > MB {
		width = 9
	} else if remainder > KB {
		width = 6
	} else {
		width = 3
	}

	// Insert missing leading zeroes
	remainderString := strconv.Itoa(remainder)
	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}
	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}

	return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
}

// ExecCmd Ejecutar comandos en Linux
func ExecCmd(sCmd string) {
	cmd := exec.Command("bash", "-c", sCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error ejecutando los comando de verificacion")
		return
	}
	fmt.Print(string(out))
}
