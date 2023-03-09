set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE users;
    CREATE USER dolganoffadmin WITH ENCRYPTED PASSWORD 'dolganoffadmin';
    GRANT ALL PRIVILEGES ON DATABASE users TO dolganoffadmin;
EOSQL