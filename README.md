# gsusuarios

Este programa crea un listado CSV para importar usuarios a GSuite a partir de un fichero CSV generado por Séneca. Para ello se utiliza el módulo *jramireziesgb/gsuite*.

El programa espera que el fichero de entrada esté codificado en UTF-8. Para convertir el fichero CSV generado por Séneca a UTF-8 se puede utilizar el siguiente comando Linux:

```bash
$ iconv -f "windows-1252" -t "UTF-8" RegAlum.csv > RegAlum.utf8.csv
```


## Instalación:
Para poder utilizar el programa habrá que compilarlo con [GO](https://golang.org/) versión 12 o posterior. A continuación se supone que [GO](https://golang.org/) está instalado en el sistema:

```bash
# Primero clonamos el repositorio de GitHub
$ git clone https://github.com/jramireziesgb/gsusuarios.git

#Ahora lo compilamos
$ go build

# Para instalarlo en el sistema simplemente lo copiamos a /usr/local/bin
$ sudo cp gsusuarios /usr/local/bin
```


## Ejemplo de uso:

* RegAlum.csv: es el fichero que genera Séneca con los alumnos que vamos a importar
* listado-gsuite.csv: es el fichero que importaremos en GSuite

```bash
$ gsusuarios -f RegAlum.utf8.csv -d "midominio.com" -x eso -u "/Alumnos/ESO" > listado-gsuite.csv
```
