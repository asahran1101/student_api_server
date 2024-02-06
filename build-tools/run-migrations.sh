if [ -f .env ]; then
  source .env
fi

export SQL_FILE="$(dirname $0)/../migrations/add_student_table.sql"
PSQL_COMMAND="psql -h postgres-service -d $DB_NAME -U $USER -w -f $SQL_FILE"

export PGPASSWORD=$PASSWORD
$PSQL_COMMAND

if ! $PSQL_COMMAND; then
    echo "Error: Failed to execute SQL file."
    exit 1
fi