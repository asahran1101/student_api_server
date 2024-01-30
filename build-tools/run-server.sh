cd $(dirname $0)/..

if [ -f .env ]; then
  source .env
fi

export SQL_FILE="$(dirname $0)/migrations/create_table.sql"

export PGPASSWORD=$PASSWORD

PSQL_COMMAND="psql -h localhost -p $PORT -d $DB_NAME -U $USER -w -f $SQL_FILE"
$PSQL_COMMAND
echo "This is DB_NAME  $PORT"
if ! $PSQL_COMMAND; then
    echo "Error: Failed to execute SQL file."
    exit 1
fi

go run ./services/main/main.go