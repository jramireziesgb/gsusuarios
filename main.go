package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/jramireziesgb/gsuite"
	"io"
	"log"
	"os"
)

const AppVersion = "1.0"

func main() {

	filePtr := flag.String("f", "", "Nombre del fichero .csv")
	sufixPtr := flag.String("x", "", "Sufijo para los nombres de usurio")
	dominioPtr := flag.String("d", "iesgbrenan.com", "Dominio de las cuentas")
	unidadPtr := flag.String("u", "/Alumnos", "Unidad Organizativa de las cuentas")
	versionPtr := flag.Bool("v", false, "Muestra la versión del programa")

	flag.Parse()

	if *versionPtr {
		fmt.Println("Versión", AppVersion)
		os.Exit(0)
	}

	if *filePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	csvFile, error := os.Open(*filePtr)
	if error != nil {
		panic(error)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var alumnos []gsuite.GesUser

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		alumno := gsuite.GesUser{
			Alumno:                      line[0],
			EstadoMatricula:             line[1],
			NIdEscolar:                  line[2],
			DNI:                         line[3],
			Direccion:                   line[4],
			Codigopostal:                line[5],
			Localidadderesidencia:       line[6],
			Fechadenacimiento:           line[7],
			Provinciaderesidencia:       line[8],
			Telefono:                    line[9],
			Telefonodeurgencia:          line[10],
			CorreoElectronico:           line[11],
			Curso:                       line[12],
			Ndelexpedientedelcentro:     line[13],
			Unidad:                      line[14],
			Primerapellido:              line[15],
			Segundoapellido:             line[16],
			Nombre:                      line[17],
			DNIPrimerturor:              line[18],
			PrimerapellidoPrimertutor:   line[19],
			SegundoapellidoPrimertutor:  line[20],
			NombrePrimertutor:           line[21],
			SexoPrimertutor:             line[22],
			DNISegundotutor:             line[23],
			PrimerapellidoSegundotutor:  line[24],
			SegundoapellidoSegundotutor: line[25],
			NombreSegundotutor:          line[26],
			SexoSegundotutor:            line[27],
			Localidaddenacimiento:       line[28],
			Añodelamatricula:            line[29],
			Ndematriculasenestecurso:    line[30],
			Observacionesdelamatricula:  line[31],
			Provincianacimiento:         line[32],
			Paisdenacimiento:            line[33],
			Edad:                        line[34],
			Nacionalidad:                line[35],
			Sexo:                        line[36],
			Fechadematricula:            line[37],
			NSegSocial:                  line[38],
		}
		alumnos = append(alumnos, alumno)
	}

	var usuarios []gsuite.User
	var usuario gsuite.User

	for _, v := range alumnos[1:] {
		usuario.NewUser(&v, *sufixPtr, *dominioPtr, *unidadPtr)
		usuarios = append(usuarios, usuario)
	}

	// Cabecera del fichero CSV
	fmt.Println("First Name [Required],Last Name [Required],Email Address [Required],Password [Required],Org Unit Path [Required]")

	for _, v := range usuarios {
		fmt.Println(v.String())
	}

}
