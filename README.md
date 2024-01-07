# Go-calendar
A go calendar that can connect to a postgres database

## Prérequis

Golang (1.20.4), Postgresql

## Dependencies


| Dependency        | Version  | 	Description           | 
|-------------------|----------|------------------------|
| github.com/lib/pq | v1.10.9	 | Postgres driver for Go |

## Installation

1 - Clone the repo
```shell
git clone github.com/BySajed/Go-Calendar
```
2 - Rename "example.config.env" to "config.env"
3 - Edit config.env file

| Variable Name | Value Type | Description                                          |
|---------------|------------|------------------------------------------------------|
| DB_HOST       | string     | The hostname of the database server.                 |
| DB_USER       | string     | The username to use when connecting to the database. |
| DB_PASSWORD   | string     | The password to use when connecting to the database. |
| DB_NAME       | string     | The name of the database to connect to.              |
| DB_PORT       | string     | Database port                                        |

4 - Create a database on your postgresql, assign user rights and enter the database information in the "config.env" file. Then execute this SQL file in your database.
```sql
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title TEXT,
    date TEXT,
    hour TEXT,
    place TEXT,
    category TEXT,
    description TEXT
);
```
## Usage 

Build and run the calendar
```shell
go run .
```
