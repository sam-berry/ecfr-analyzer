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
curl -X POST -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/import-agencies'
```

### Import Titles

```
curl -X POST -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/import-titles'
```

To only import specific titles:

```
curl -X POST -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/import-titles?titles=16,17,18'
```

## Computed Value APIs

### Title Metrics

```
curl -X POST -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/compute/title-metrics'
```

### Title Metrics

```
curl -X POST -H 'Authorization: Bearer E0E91B8C-60B6-439A-8C48-6D66D5A1BE55' 'localhost:8090/ecfr-service/compute/agency-metrics'
```

## Areas for Improvement

* Precompute text values for title XML. Write a job to parse the entire CFR and save
  chapters/parts/sections/etc as structured data that can be easily queried. Do not need XLST for
  this, can be done by iterating through the tree and following the DIV# guidelines.
* Create common Goroutine runner that encapsulates the channel and wait group processing, as it is
  similar throughout the project
* Subagencies metric import - this was an experiment that ended up working, but it could be handled cleaner
  instead of passing the `onlySubAgencies` variable and forking the top-level logic
* Import historical CFR records and compute metrics based on changes over time