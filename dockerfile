# Usa una imagen base oficial de Go
FROM golang:1.18-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el go.mod y go.sum para la gestión de dependencias
COPY go.mod go.sum ./

# Descarga las dependencias de Go
RUN go mod tidy

# Copia el código fuente del servicio de productos
COPY . .

# Expone el puerto que utilizará la aplicación
EXPOSE 3010

# Construye la aplicación Go
RUN go build -o list-products-service .

# Ejecuta la aplicación
CMD ["./list-products-service"]
