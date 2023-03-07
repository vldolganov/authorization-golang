CREATE DATABASE users;
    CREATE USER dolganoffadmin WITH ENCRYPTED PASSWORD 'dolganoffadmin';
    GRANT ALL PRIVILEGES ON DATABASE users TO dolganoffadmin;