package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createTableRecruiters = `CREATE TABLE IF NOT EXISTS recruiters
(
    id          varchar(20)     not null primary key,
	name 		varchar(20) 	not null,
    company 	varchar(20)		not null,
    jobs 		UUID[] 			not null,
	password 	varchar(100)	not null,
    created_at  TIMESTAMP   	not null,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
);`

func create_table_recruiters() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableRecruiters)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
