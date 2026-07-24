### New proyect Opengine (Suite Automation)



## Arquitectura del backend 

```
opengine/
├── cmd/
│   └── api/
│       └── main.go                 # Punto de entrada de la aplicación
├── internal/
│   ├── domain/                     # Modelos puros y structs de la base de datos
│   │   ├── environment.go
│   │   ├── endpoint.go
│   │   ├── test_case.go
│   │   └── account.go
│   ├── repository/                 # Capa de acceso a PostgreSQL (SQL puro/pgx)
│   │   ├── test_case_repository.go
│   │   └── account_repository.go
│   ├── usecase/                    # Motor de construcción, sustitución y ejecutor
│   │   ├── template_engine.go      # Inyección de variables {{placeholders}}
│   │   ├── assertion_evaluator.go  # Lógica de comparación con gjson
│   │   └── suite_runner.go         # Coordinador concurrente (Goroutines + WaitGroup)
│   └── handler/                    # Handlers HTTP (Gin) para exponer las APIs a React
│       ├── test_case_handler.go
│       └── suite_handler.go
├── pkg/                            # Módulos reutilizables o helpers generales
│   └── httpclient/                 # Cliente HTTP optimizado con timeouts
├── config/                         # Lectura de variables de entorno (ENV)
├── docker-compose.yml              # PostgreSQL + Nginx + Go App
├── Dockerfile                      # Build multi-stage optimizado
└── go.mod
│
├── psql                            # Repo BD
│   └── docker-compose.yml          # Orquestador para desplegar la BD
│   └── backup.sql                  # backup

```


## Motor de Payloads Dinámicos

```
[endpoints_manager.request_body] (Plantilla Base con {{placeholders}})
                 +
[test_cases.request_body]        (Mapa Key-Value de Variables)
                 +
[Variables de Sistema]           (Fechas dinámicas, TraceIDs, etc.)
                 │
                 ▼ (Inyección en Memoria con strings.Replacer en Go)
                 │
      [Payload JSON Consolidado] ──► (HTTP Request a la API Externa)
```