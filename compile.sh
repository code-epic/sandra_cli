#!/bin/bash
#


clear
echo -e "\n"
echo -e "[+] Iniciando Compilacion"

REPO="github.com/code-epic/sandra_cli"
VERSION="$REPO/cmd.Version=1.0.1-"$(git log --pretty=format:"%h" -n 1)
BUILD_DATE="$REPO/cmd.Fecha=$( date '+%Y-%m-%d %T')" 
COMPILACION="$REPO/cmd.Compilacion=$1"




#GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -X '$VERSION' -X '$BUILD_DATE' -X '$COMPILACION'" -o sandra_cli main.go
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w -X '$VERSION' -X '$BUILD_DATE' -X '$COMPILACION'" -o sandra_cli main.go

echo -e "    - Compilacion terminada\n"
