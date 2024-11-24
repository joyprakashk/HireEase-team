package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createTableJobs = `CREATE TABLE IF NOT EXISTS jobs
(
    id          UUID        	not null primary key,
	title 		varchar(20) 	not null,
    description varchar(2000)	not null,
    company     varchar(20) 	not null,
	location 	varchar(20) 	not null,
	domain 		varchar(20) 	not null,
	skills		varchar(100) 	not null,
    created_at  TIMESTAMP   	not null,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
);`

func create_table_jobs() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableJobs)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
