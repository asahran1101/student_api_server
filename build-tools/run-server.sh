cd $(dirname $0)/..

go get -u github.com/gin-gonic/gin
go get -u github.com/golang/mock
go get -u github.com/joho/godotenv
go get -u github.com/lib/pq
go get -u github.com/stretchr/testify/assert


if [ -f .env ]; then
  source .env
fi

export SQL_FILE="$(dirname $0)/../migrations/add_student_table.sql"

export PGPASSWORD=$PASSWORD

PSQL_COMMAND="psql -h localhost -p $PORT -d $DB_NAME -U $USER -w -f $SQL_FILE"
$PSQL_COMMAND
if ! $PSQL_COMMAND; then
    echo "Error: Failed to execute SQL file."
    exit 1
fi

go run ./services/main/main.go