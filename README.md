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
```txt
Technical Description:
Our project is build only with one package (main) but each function is in a different file. This method make hte debugging part easier. Here is the content of each file:

main.go:
    - start menu
    - start connection to database
    - notification
    - 

newEvent.go:
    - creation new events
    - upload event in database

showEvent.go:
    - show an event (or all events)
    - filter your research

modifyEvent.go:
    - modify an existing event
    - update the event in the database

deleteEvent.go:
    - delete an existing event
    - update event in database

export.go:
    - export/convert the file in JSON or CSV

This code work like a tree structure. The user will chose an option, and the switch in the main will follow the way. When the user finished, he will come back to the origin of the tree (main.go)

Obviously, we have incorporate errors management into our code in forecasting possible user errors.
```
