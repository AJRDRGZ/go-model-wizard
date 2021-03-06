package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func showMainMenu() {
	color.Cyan("*************************************")
	color.Cyan("* Sistema de generación de paquetes *")
	color.Cyan("*************************************")
	fmt.Println()

	scanName()
	scanTable()
	scanFields()
	scanSourceConfig()
}

func scanName() {
	color.Cyan("1. Digite el nombre del paquete en singular y minuscula:")
	fmt.Scan(&n)
	if n == "" {
		color.Red("el nombre del paquete es obligatorio")
		os.Exit(1)
	}
}

func scanTable() {
	color.Cyan("2. Digite el nombre de la tabla en plural y minuscula:")
	fmt.Scan(&t)
	if t == "" {
		color.Red("el nombre de la tabla es obligatorio")
		os.Exit(1)
	}
}

func scanFields() {
	color.Cyan("3. Digite los campos del modelo.")
	color.Cyan("El formato es: nombre:tipo:nonulo:tamaño.")
	color.Cyan("* cada campo debe estar separada por un espacio. ej:")
	color.Cyan("name:string:f:50 age:int birth:time.Time:t other:bool")
	color.Cyan("* nombre: nombre del campo, minúsculas.")
	color.Cyan("* tipo: string, int, float32, float64, time.Time, bool.")
	color.Cyan("* nonulo: t si permite nulos, f no permite nulos. (por defecto es f)")
	color.Cyan("* tamaño: número entero. Sólo aplica para string.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := scanner.Text()
	err := scanner.Err()
	if err != nil {
		color.Red("error al leer los campos:", err)
		os.Exit(1)
	}

	fs = getFields(fields)
	if len(fs) == 0 {
		color.Red("no se han recibido campos del modelo")
		os.Exit(1)
	}
}

func getFields(value string) []Field {
	var err error
	rs := make([]Field, 0)
	fields := strings.Split(value, " ")
	for _, v := range fields {
		field := strings.Split(v, ":")
		nn := "NOT NULL"
		i := 0
		if len(field) >= 3 {
			if strings.ToLower(field[2]) == "t" {
				nn = ""
			}
		}
		if len(field) == 4 {
			i, err = strconv.Atoi(field[3])
			if err != nil {
				log.Fatalf("%s no es un número válido: %v", field[3], err)
			}

		}
		f := Field{field[0], field[1], nn, i}
		rs = append(rs, f)
	}

	return rs
}

func scanSourceConfig() {
	var v string
	ps = make(map[string]string, 0)

	color.Cyan("1. Ubicación del archivo de configuracion (json)")
	color.Cyan("* ruta absoluta o relativa al archivo de configuracion")
	fmt.Scan(&v)

	file, err := ioutil.ReadFile(v)
	if err != nil {
		e := fmt.Sprintf("no se pudo abrir el archivo de configuración: %v", err)
		color.Red(e)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &ps)
	if err != nil {
		e := fmt.Sprintf("no se pudo convertir la configuración en mapa: %v", err)
		color.Red(e)
		os.Exit(1)
	}
}
