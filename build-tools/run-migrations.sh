if [ -f .env ]; then
  source .env
fi

export SQL_FILE="$(dirname $0)/../migrations/add_student_table.sql"
PSQL_COMMAND="psql -h postgres-service -d $POSTGRES_DB -U $POSTGRES_USER -w -f $SQL_FILE"

export PGPASSWORD=$POSTGRES_PASSWORD
$PSQL_COMMAND

if ! $PSQL_COMMAND; then
    echo "Error: Failed to execute SQL file."
    exit 1
fi