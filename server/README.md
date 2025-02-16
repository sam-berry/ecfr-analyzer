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
export ECFR_ADMIN_TOKEN="E0E91B8C-60B6-439A-8C48-6D66D5A1BE55"
export ECFR_DB_USER="ecfr-app"
export ECFR_DB_PASS=""
export ECFR_DB_HOST="localhost"
export ECFR_DB_PORT="5432"
export ECFR_DB_NAME="ecfr"
export ECFR_DB_INSTANCE_CONNECTION_NAME=""
export ECFR_DEVELOPMENT="true"
```

## API

### Import Agencies

```
curl -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/import-agencies'
```

### Import Titles

```
curl -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/import-titles'
```

To only import specific titles:
```
curl -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/import-titles?titles=16,17,18'
```
