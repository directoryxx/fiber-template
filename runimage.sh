docker build -t test .
docker run --rm -p 8000:8000 -e BYPASS_ENV_FILE="file" -e APP_PORT=8000 -e DB_HOST="db-standalone-postgresql-1" -e DB_USERNAME="postgres" -e DB_PASSWORD="passwordq" -e DB_NAME="template" -e DB_PORT="5432" --network dev test
