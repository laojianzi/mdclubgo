package database

import (
	"fmt"

	"gorm.io/gorm"
)

// TablePrefix db table prefix
const TablePrefix = "mc_"

// WithTablePrefix return a string with table prefix append model
func WithTablePrefix(model string) string {
	return fmt.Sprintf("%s%s", TablePrefix, model)
}

// Has does it exist
func Has(db *gorm.DB, model interface{}, scopes ...func(*gorm.DB) *gorm.DB) (bool, error) {
	var result int
	sql := db.Session(&gorm.Session{DryRun: true}).Model(model).Select("id").Scopes(scopes...).Statement.SQL.String()
	err := db.Raw("SELECT EXISTS(?)", gorm.Expr(sql)).Row().Scan(&result)
	if err != nil {
		return false, err
	}

	return result == 1, nil
}
