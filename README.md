# Postgres-Golang-Connection

- ### Connection String
  ```Golang
  postgres://<username>:<password>@localhost/<database_name>?sslmode=disable
  ```
- ### Run Migration
  - #### Up
          go run main.go --migration up
  - #### Down
          go run main.go --migration down
- ### Error Handling

  If you're getting the bellow error

  ```
  Error parsing migration (table_name.sql): ERROR: no Up/Down annotations found, so no statements were executed.
  ```

  In the migration file Set tags as `Up` and `Down` instead of `up` `down`

  - -- +migrate UP
  - -- +migrate Down
