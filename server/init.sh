#!/bin/sh
# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."

until pg_isready -h database -p 5432 -U "$DB_USER"; do
  >&2 echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

echo "PostgreSQL is up and running"


echo "Checking if database exists..."
DB_EXISTS=$(PGPASSWORD="$POSTGRES_PASSWORD" psql -h database -p 5432 -U "$POSTGRES_USER" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'")

if [ "$DB_EXISTS" = "1" ]; then
  echo "Database $DB_NAME already exists"
else
  echo "Creating database $DB_NAME..."
  PGPASSWORD="$POSTGRES_PASSWORD" psql -h database -p 5432 -U "$POSTGRES_USER" -c "CREATE DATABASE $DB_NAME;"
fi

# Create user and set permissions
echo "Creating user and setting permissions..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h database -p 5432 -U "$POSTGRES_USER" -c "DO \$\$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname='$DB_USER') THEN
    CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
  END IF;
END \$\$;"
PGPASSWORD="$POSTGRES_PASSWORD" psql -h database -p 5432 -U "$POSTGRES_USER" -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"

# Grant permissions on the public schema
echo "Granting permissions on the public schema..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h database -p 5432 -U "$POSTGRES_USER" -d "$DB_NAME" -c "GRANT USAGE, CREATE ON SCHEMA public TO $DB_USER;"

# Run migrations
echo "Running migrations..."
migrate -path /app/db/migrations -database "postgres://$DB_USER:$DB_PASSWORD@db:5432/$DB_NAME?sslmode=disable" up

echo "Database setup complete"

# Start the application
echo "Starting the application..."
exec /app/an-overengineered-app

echo "Database setup and application start complete"