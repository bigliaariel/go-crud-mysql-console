English

To get a free Mysql service, go to freesqldatabase.com and
with a few simple steps you will get a free mysql service of 5MB enough to do
tests.

There you will get the following data:

Host: sql10.freesqldatabase.com
Database name: sql105xxxxx
Database user: sql105xxxxx
Database password: xxxxxxxx
Port number: 3306

With these data you can build the connection string of the main.go file:

db, err := sql.Open("mysql", "databaseUser:youpassword@tcp(sql10.freesqldatabase.com:3306)/databaseName")



-------------------------------------------------------------------------------
Español

Para obtener un servivio gratuito de Mysql, diríjase a la direccion freesqldatabase.com y
con unos simpleas pasos obtendrá un servicio a mysql gratiuto de 5MB suficientes para hacer
las pruebas.

Allí obtendrá los siguientes datos:

Host: sql10.freesqldatabase.com
Database name: sql105xxxxx
Database user: sql105xxxxx
Database password: xxxxxxxx
Port number: 3306

con esos datos usted podrá armar la cadena de conección del archivo main.go:

db, err := sql.Open("mysql", "databaseUser:youpassword@tcp(sql10.freesqldatabase.com:3306)/databaseName")


