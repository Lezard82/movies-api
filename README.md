# movies-api
Technical assesment for backend developer possition at KenTech, S.L.U.


# project structure

/movies-api
│── /cmd                 # Punto de entrada de la aplicación
│── /config              # Configuración del proyecto (variables de entorno, etc.)
│── /internal
│   ├── /app             # Casos de uso (Aplicación - Lógica de negocio)
│   ├── /domain          # Entidades y modelos del dominio
│   ├── /infrastructure  # Implementaciones externas (DB, HTTP, APIs externas)
│   ├── /interface       # Interfaces de los adaptadores
│── /pkg                 # Paquetes reutilizables
│── main.go              # Punto de entrada de la aplicación
│── go.mod               # Módulo Go
│── go.sum               # Dependencias
