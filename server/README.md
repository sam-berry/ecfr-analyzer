# eCFR Analyzer

## Requirements
* Go 1.21+
* Postgres 17

## Set Up Database Locally
* `createuser ecfr-app`
* `createdb ecfr`
* `psql ecfr`
* `grant all privileges on database ecfr to "ecfr-app";`
* `grant all on schema public TO "ecfr-app";`

## Run locally
`go run server.go`

## Environment Variables

### Development
```
export ECFR_TOKEN_SECRET="E0E91B8C-60B6-439A-8C48-6D66D5A1BE55"
export ECFR_DB_USER="ecfr-app"
export ECFR_DB_PASS=""
export ECFR_DB_HOST="localhost"
export ECFR_DB_PORT="5432"
export ECFR_DB_NAME="ecfr"
export ECFR_DB_INSTANCE_CONNECTION_NAME=""
export ECFR_DEVELOPMENT="true"
```