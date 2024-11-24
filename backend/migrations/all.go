// This is auto-generated file using 'gofr migrate' tool. DO NOT EDIT.
package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate {
	
		20241123111558: create_table_jobs(),	
		20241123182920: create_table_recruiters(),	
		20241123200819: create_table_students(),
	}
}
