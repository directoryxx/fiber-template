# Fiber Template

## Installation

1. Install Docker
2. Copy .env.example to .env
3. Run Docker Compose

## Generator CRUD
````
go run cmd/cli/main.go generatorcrud <feature>
````

Example
````
go run cmd/cli/main.go generatorcrud user
````

## Migration
1. Create Migration
````
go run cmd/cli/main.go migration create <table name>
````
Example
````
go run cmd/cli/main.go migration create create_table_users
````

2. Migrating Table
````
go run cmd/cli/main.go migration up
````

3. Rollback Table
````
go run cmd/cli/main.go migration down
````

## Test
````
BYPASS_ENV_FILE=true TEST_MODE=true go test ./test/ -v
````