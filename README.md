I can't over engineer stuff at work. So, I'm going to over engineer the shit out of this codebase.

### How to run server code

## Generate smtp certificate for email

- Run below commands to generate self-signed smtp certificate. Without this, sending mail won't work.
  - `mkdir -p server/certs/smtp` && `cd server/certs/smtp`
  - `openssl req -x509 -newkey rsa:4096 -nodes -keyout key.pem -out cert.pem -sha256 -days 36500 -addext "subjectAltName = DNS:localhost"`, this will issue a self-signed certificate for 100 years!

## First approach (suitable for local development)

### Before running the code

- You should have **go >= 1.22, postgres and mailhog/mailpit (smtp server)** installed in your system
- Install `air` with this command `go install github.com/air-verse/air@latest`. This will help with hot reloading
- Configure your env value in .env.dev
- Create appropriate database, db user, db password. You can follow below command to create it in postgres
  - Login to postgres: `psql -U postgres`. This is for default login without password.
  - Create database: `CREATE DATABASE $DB_NAME;`. $DB_NAME value should match with .env.dev $DB_NAME value
  - Create user: `CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';`, value should match with .env.dev
  - Grant permissions: `GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;`
  - Grant permissions on public schema: `GRANT USAGE, CREATE ON SCHEMA public TO $DB_USER;`

Now copy paste below commands one at a time in terminal

- Run `cd server`
- Run `go mod download`
- Run `air`
- Now api server should be available at `http:localhost:$HTTP_PORT`

## Second approach - docker version(suitable for trying the code)

- Install docker
- Run `cd server`
- Run `docker compose up --build`
- Now api server should be available at `http:localhost:$HTTP_PORT`
