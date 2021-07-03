# Reviews

Microservicio que permite agregar reseñas a articulos del [ecommerce](https://github.com/nmarsollier/ecommerce). Diseñado para la materia Arquitectura de Microservicios de UTN FRM.

Se comunica con los microservicios de [Auth](https://github.com/nmarsollier/ecommerce_auth_node) y [Catalog](https://github.com/nmarsollier/ecommerce_catalog_java)

## Casos de Uso:

- Agregar reseña.
- Obtener las reseñas de un articulo.
- Obtener el puntaje de un articulo.
- Eliminar reseña.

## Dependencias

- Golang (Chi)
- MongoDB
- RabbitMQ

## Documentación

En http://localhost:8020/swagger/ se puede ver la documentación del microservicio.

## Ejecución

`go run main.go`
