package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createTableStudents = `CREATE TABLE IF NOT EXISTS students
(
    id          varchar(20)     not null primary key,
	name 		varchar(20) 	not null,
    gender 		varchar(6) 		not null,
    branch 		varchar(20) 	not null,
    cgpa 		float4 			not null,
    courses 	varchar(1000) 	not null,
    github 		varchar(100) 	not null,
    linkedin 	varchar(100) 	not null,
    interests 	varchar(100) 	not null,
    skills 		varchar(100) 	not null,
    location 	varchar(20) 	not null,
	password 	varchar(100)	not null,
    created_at  TIMESTAMP   	not null,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
);`

func create_table_students() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableStudents)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
