# gsusuarios

Este programa crea un listado CSV para importar usuarios a GSuite a partir de un fichero CSV generado por Séneca. Para ello utiliza el módulo jramireziesgb/gsuite.

El programa espera que el fichero de entrada esté codificado en UTF-8. Para convertir el fichero CSV generado por Séneca a UTF-8 se puede utilizar el siguiente comando Linux:

```bash
$ iconv -f "windows-1252" -t "UTF-8" RegAlum.csv > RegAlum.utf8.csv
```

Ejemplos de uso:

* RegAlum.csv: es el fichero que genera Séneca con los alumnos que vamos a importar
* listado-gsuite.csv: es el fichero que importaremos en GSuite

```bash
$ gsusuarios -f RegAlum.utf8.csv -d "midominio.com" -x eso -u "/Alumnos/ESO" > listado-gsuite.csv
```
