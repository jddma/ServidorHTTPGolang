# ServidorHTTPGolang
  Código fuente un servidor web escrito en Golang que se puede usar en cualquier proyecto en el que se desea usar golang como lenguaje de backend. El servidor está pensado para implementar el patrón de diseño MVC y ser usado es una aplicación REST. Permite el uso de handlers(rutas), de middewares y de designar el tipo de solicitud que se realizará a cada una de las rutas; Incluye también un sistema de sesiones para autenticar usuarios. El proyecto funciona como una plantilla sobre la cual se estructura el saplicativo web a diseñar. deseado.
  
## Dependencias.
- gorilla/securecookie

## Descarga e implementación
``` bash
$ git clone https://github.com/jddma/ServidorHTTPGolang.git
$ go get github.com/gorilla/securecookie
```
## Licencia
ServidorHTTPGolang esta bajo la licencia GNU General Public License v3.0, Leer el archivo [LICENSE](https://github.com/jddma/ServidorHTTPGolang/blob/master/LICENSE) para mas información.