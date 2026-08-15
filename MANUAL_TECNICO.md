# Manual Técnico - GraphQL Go CRUD

## Tabla de Contenidos
1. [Arquitectura](#arquitectura)
2. [Stack Tecnológico](#stack-tecnológico)
3. [Estructura del Proyecto](#estructura-del-proyecto)
4. [Instalación y Configuración](#instalación-y-configuración)
5. [Desarrollo Local](#desarrollo-local)
6. [Base de Datos](#base-de-datos)
7. [GraphQL](#graphql)
8. [Docker y Deployment](#docker-y-deployment)
9. [CI/CD](#cicd)
10. [Troubleshooting](#troubleshooting)

---

## Arquitectura

### Flujo de Datos
```
Usuario (Browser)
    ↓
Templates (Go)
    ↓
HTTP Handler (/graphql, /)
    ↓
GraphQL Server (graph-gophers/graphql-go)
    ↓
Resolver Layer
    ↓
Database Layer (DB Package)
    ↓
PostgreSQL
```

### Capas

| Capa | Responsabilidad | Archivos |
|------|-----------------|----------|
| **Web** | Servir templates HTML | `internal/web/web.go` |
| **GraphQL** | Definir schema y resolver | `internal/graphql/graphql.go` |
| **Database** | Operaciones CRUD y SQL | `internal/db/db.go` |
| **Main** | Inicialización y setup | `main.go` |

---

## Stack Tecnológico

- **Lenguaje**: Go 1.25
- **GraphQL**: `github.com/graph-gophers/graphql-go`
- **Base de Datos**: PostgreSQL 14+
- **Driver PostgreSQL**: `github.com/lib/pq`
- **Contenerización**: Docker + Docker Compose
- **CI/CD**: GitHub Actions
- **Templating**: `html/template` (Go nativo)

---

## Estructura del Proyecto

```
graphqlexample/
├── main.go                          # Punto de entrada
├── go.mod                           # Dependencias
├── go.sum                           # Checksum de dependencias
├── Dockerfile                       # Imagen Docker
├── docker-compose.yml               # Docker Compose local
├── docker-compose-prod.yml          # Docker Compose producción
├── .env                             # Variables de entorno local
├── .github/
│   └── workflows/
│       ├── go-ci.yml                # CI Pipeline
│       └── cd.yml                   # CD Pipeline
├── internal/
│   ├── db/
│   │   ├── db.go                    # Conexión y CRUD
│   │   └── db_test.go               # Tests unitarios
│   ├── graphql/
│   │   └── graphql.go               # Schema y Resolver
│   └── web/
│       └── web.go                   # Handler de templates
└── templates/
    └── index.html                   # UI HTML/CSS/JS
```

---

## Instalación y Configuración

### Requisitos Previos
- Go 1.25+
- PostgreSQL 14+
- Docker y Docker Compose (opcional, para deployment)
- Git

### Pasos de Instalación

#### 1. Clonar el repositorio
```bash
git clone https://github.com/rianeiromiron/graphqlgo.git
cd graphqlgo
```

#### 2. Descargar dependencias
```bash
go mod download
```

#### 3. Configurar variables de entorno
```bash
# Crear .env en raíz del proyecto
cat > .env << EOF
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_contraseña
DB_NAME=gestion_documentos
DB_SSLMODE=disable
APP_PORT=8080
EOF
```

#### 4. Crear base de datos
```bash
# Conectar a PostgreSQL
psql -U postgres -h localhost

# En la consola psql
CREATE DATABASE gestion_documentos;
\q
```

#### 5. Ejecutar la aplicación
```bash
go run .
```

La app estará disponible en `http://localhost:8080`

---

## Desarrollo Local

### Estructura de Paquetes

#### `main.go`
- Lee variables de entorno
- Inicializa conexión a BD
- Corre migraciones
- Levanta servidor HTTP en puerto `APP_PORT`

**Funciones principales**:
```go
func getEnv(key, fallback string) string        // Lee variable de entorno
func getEnvInt(key string, fallback int) int    // Lee variable numérica
```

#### `internal/web/web.go`
- Define handler para ruta `/`
- Carga template `templates/index.html`
- Retorna HTML renderizado

```go
func NewHandler() http.Handler  // Retorna handler para ruta /
```

#### `internal/graphql/graphql.go`
- Define schema GraphQL como string
- Implementa resolver con métodos que mapean a campos GraphQL
- Handler HTTP para `/graphql`

**Funciones principales**:
```go
func NewHTTPHandler(db *sql.DB) http.HandlerFunc  // Handler para /graphql
```

#### `internal/db/db.go`
- Gestiona conexión a PostgreSQL
- Ejecuta migraciones
- Implementa operaciones CRUD

**Funciones principales**:
```go
func New(cfg Config) (*sql.DB, error)           // Conecta a BD
func Migrate(db *sql.DB) error                  // Crea tablas
func ListTasks(db *sql.DB) ([]Task, error)      // Lee todas las tareas
func GetTaskByID(db *sql.DB, id string) (*Task, error)
func CreateTask(db *sql.DB, input CreateTaskInput) (*Task, error)
func UpdateTask(db *sql.DB, input UpdateTaskInput) (*Task, error)
func DeleteTask(db *sql.DB, id string) (bool, error)
```

### Ejecutar Tests
```bash
go test ./...
```

### Compilar
```bash
go build .
```

### Variables de Entorno

| Variable | Defecto | Descripción |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Host de PostgreSQL |
| `DB_PORT` | `5432` | Puerto de PostgreSQL |
| `DB_USER` | `postgres` | Usuario de BD |
| `DB_PASSWORD` | - | Contraseña de BD |
| `DB_NAME` | `gestion_documentos` | Nombre de BD |
| `DB_SSLMODE` | `disable` | SSL mode: disable/require/verify-ca/verify-full |
| `APP_PORT` | `8080` | Puerto de app |

---

## Base de Datos

### Schema
```sql
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### CRUD Operaciones

#### Crear Tarea
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createTask(input: {name: \"Mi Tarea\"}) { id name } }"
  }'
```

#### Listar Tareas
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ tasks { id name completed } }"}'
```

#### Obtener Tarea por ID
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ task(id: \"<id>\") { id name description } }"
  }'
```

#### Actualizar Tarea
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { updateTask(input: {id: \"<id>\", completed: true}) { id completed } }"
  }'
```

#### Eliminar Tarea
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { deleteTask(id: \"<id>\") }"
  }'
```

### Migraciones
Las migraciones se ejecutan automáticamente al iniciar la app con `db.Migrate(conn)` en `main.go`.

Para resetear la BD:
```bash
psql -U postgres -h localhost -d gestion_documentos -c "DROP TABLE tasks;"
# Reiniciar app para recrear tabla
```

---

## GraphQL

### Schema
```graphql
type Query {
    tasks: [Task!]!
    task(id: ID!): Task
}

type Mutation {
    createTask(input: CreateTaskInput!): Task
    updateTask(input: UpdateTaskInput!): Task
    deleteTask(id: ID!): Boolean!
}

type Task {
    id: ID!
    name: String!
    description: String
    completed: Boolean!
}

input CreateTaskInput {
    name: String!
    description: String
}

input UpdateTaskInput {
    id: ID!
    name: String
    description: String
    completed: Boolean
}
```

### Resolver Pattern

Cada campo en GraphQL se mapea a un método Go:

```go
// Query.tasks → func (r *Resolver) Tasks(ctx context.Context) ([]*Task, error)
// Query.task(id) → func (r *Resolver) Task(ctx context.Context, args struct{ID string}) (*Task, error)
// Mutation.createTask → func (r *Resolver) CreateTask(ctx context.Context, args struct{Input CreateTaskInput}) (*Task, error)
```

### Endpoint GraphQL
- **URL**: `http://localhost:8080/graphql`
- **Método**: POST
- **Content-Type**: `application/json`
- **Body**: `{"query": "...", "variables": {...}}`

---

## Docker y Deployment

### Dockerfile (Multi-stage)

**Stage 1 (Builder)**:
- Base: `golang:1.25-alpine`
- Compila el binario
- Copia templates

**Stage 2 (Runtime)**:
- Base: `alpine:3.20`
- Solo el binario compilado
- Solo los templates
- Variables de entorno predeterminadas

### Compilar Imagen
```bash
docker build -t rianeiromiron/graphqlgo:latest .
```

### Ejecutar Localmente
```bash
# Con docker-compose
docker compose -f docker-compose-prod.yml up -d

# Verificar
docker ps
docker logs graphqlgo-prod
```

### Variables de Entorno en Docker

Se inyectan desde el archivo `.env` o desde `docker-compose-prod.yml`:

```yaml
environment:
  APP_PORT: 8080
  DB_HOST: ${DB_HOST}
  DB_PORT: ${DB_PORT:-5432}
  DB_USER: ${DB_USER}
  DB_PASSWORD: ${DB_PASSWORD}
  DB_NAME: ${DB_NAME}
  DB_SSLMODE: ${DB_SSLMODE:-disable}
```

### Red Docker
La app corre en la red `pruebas-graphql` definida en `docker-compose-prod.yml`.

Esto permite agregar más servicios (API, workers, etc.) en la misma red.

### Acceder a PostgreSQL desde Contenedor
Si PostgreSQL está en `host.docker.internal`:
```yaml
DB_HOST: host.docker.internal  # Windows/Mac
DB_HOST: 172.17.0.1            # Linux
```

---

## CI/CD

### GitHub Actions CI (`.github/workflows/go-ci.yml`)

Se ejecuta en cada **push** a `master` o `main`.

**Pasos**:
1. Checkout código
2. Setup Go 1.25
3. Corre tests
4. Valida formato (`gofmt`)
5. Compila binario

**Triggers**:
```yaml
on:
  push:
    branches:
      - master
      - main
```

### GitHub Actions CD (`.github/workflows/cd.yml`)

Se ejecuta en cada **push** a `master` o `main` (después de CI).

**Pasos**:
1. Checkout código
2. Setup Go 1.25
3. Corre tests
4. Compila app
5. Login a Docker Hub
6. Build y push imagen
7. Genera `.env` con secrets
8. Muestra `.env`

**Secrets requeridos** (en GitHub Settings → Secrets):
```
DOCKER_USERNAME
DOCKER_PASSWORD
DB_HOST
DB_PORT
DB_USER
DB_PASSWORD
DB_NAME
DB_SSLMODE
```

**Resultado**:
- Imagen publicada en Docker Hub: `docker.io/rianeiromiron/graphqlgo:latest`

### Configurar Secrets en GitHub

1. Ve a tu repositorio
2. Settings → Secrets and variables → Actions
3. Haz clic en "New repository secret"
4. Agrega cada secret (ver tabla arriba)

---

## Troubleshooting

### Error: `failed to connect to database`

**Causa**: PostgreSQL no está corriendo o credenciales son incorrectas.

**Solución**:
```bash
# Verificar que PostgreSQL está corriendo
psql -U postgres -h localhost

# Verificar variables de entorno
echo $DB_HOST $DB_PORT $DB_USER $DB_PASSWORD
```

### Error: `panic: open templates/index.html: no such file or directory`

**Causa**: App no encuentra el archivo de template.

**Solución**:
```bash
# Verificar que el archivo existe
ls templates/

# Asegurarse de ejecutar desde raíz del proyecto
cd graphqlexample
go run .
```

### Error en Docker: `connection refused`

**Causa**: El contenedor no puede alcanzar PostgreSQL.

**Solución**:
```bash
# Verificar que PostgreSQL está corriendo
docker ps | grep db-1

# Usar IP real en lugar de host.docker.internal
ipconfig | findstr "IPv4"
# Actualizar .env con la IP real
```

### Error: `port already in use`

**Causa**: Puerto 8080 (o 8081) está ocupado.

**Solución**:
```bash
# Detener contenedores
docker compose -f docker-compose-prod.yml down

# O usar otro puerto en .env
echo "APP_PORT=8082" >> .env
```

### GraphQL Query devuelve error

**Causa**: Syntax error o campo no existe.

**Solución**:
```bash
# Verificar schema en internal/graphql/graphql.go
# Probar query en herramienta GraphQL

# Ejemplo válido:
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ tasks { id name } }"}'
```

### Tests fallan

**Causa**: Base de datos no está configurada o migraciones no corrieron.

**Solución**:
```bash
# Asegurarse de tener PostgreSQL corriendo
psql -U postgres -h localhost

# Ejecutar tests con verbose
go test -v ./...
```

---

## Desarrollo Futuro

### Posibles Mejoras
- [ ] Autenticación/Autorización
- [ ] Validación de input mejorada
- [ ] Paginación en listados
- [ ] Filtros avanzados
- [ ] Caching (Redis)
- [ ] Logs estructurados (logrus/zap)
- [ ] Métricas (Prometheus)
- [ ] Health checks
- [ ] Graceful shutdown

### Agregar Nuevas Entidades
1. Crear tabla en `internal/db/db.go` (en `Migrate()`)
2. Crear struct en `internal/db/db.go`
3. Agregar CRUD functions en `internal/db/db.go`
4. Actualizar schema GraphQL en `internal/graphql/graphql.go`
5. Implementar resolver en `internal/graphql/graphql.go`

---

## Referencias

- [Go Documentation](https://golang.org/doc)
- [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

---

**Última actualización**: 2026-08-15
**Versión**: 1.0.0
