# gsusuarios

Usando el paquete jramireziesgb/gsuite crea, a partir de un fichero CSV generado por Séneca, un listado CSV para importar los usuarios a GSuite.

Ejemplos de uso:
RegAlum.csv: es el fichero que genera Séneca con los alumnos que vamos a importar

listado.csv: es el fichero que importaremos en GSuite

\$ gsusuarios -f RegAlum.csv -d "midominio.com" -x eso -u "/Alumnos/ESO" > listado.csv
