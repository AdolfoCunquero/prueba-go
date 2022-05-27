## Prueba Go

El API cuenta con un unico endpoint `/test` que devuelve un array de 15 objetos recuperados del API [api.chucknorris.io](https://api.chucknorris.io/jokes/random)

Para mejorar el tiempo de respuesta, se utilizan `Goroutines` realizando multiples solicitudes simultaneamente, con la capacidad de solicitar los datos por lotes hasta llegar a la cantidad de registros unicos necesarios, para ello se puede modificar el valor por defecto de la variable `totalRecords` para el total de registros y `batchZize` para el tamaño de solicitudes concurrentes, ambas variables se encuentran en la ruta `routers/api.go`

#### Ejecutar el proyecto en docker

Se creará una imagen llamada **prueba-go**

- `docker build --tag prueba-go .`
- `docker run --rm -it -p 3500:3500 prueba-go`

Ver los resultados obtenidos desde el navegador o alguna herramienta para consumir servicios REST
[http://localhost:3500/test](http://localhost:3500/test)