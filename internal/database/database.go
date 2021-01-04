package database

import (
	"fmt"

	"gorm.io/gorm"
)

// TablePrefix db table prefix
const TablePrefix = "mc_"

// TableInterface database table functions
type TableInterface interface {
	TableName() string
	PrimaryKey() string
}

// WithTablePrefix return a string with table prefix append model
func WithTablePrefix(model string) string {
	return fmt.Sprintf("%s%s", TablePrefix, model)
}

// Has does it exist
func Has(db *gorm.DB, model TableInterface, scopes ...func(*gorm.DB) *gorm.DB) (bool, error) {
	var result int
	dbSession := db.Session(&gorm.Session{NewDB: true, DryRun: true})
	stmt := dbSession.Model(model).Select(model.PrimaryKey()).Scopes(scopes...).Find(model).Statement
	sql := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	row := db.Raw("SELECT EXISTS(?)", gorm.Expr(sql)).Row()
	if row == nil {
		return false, nil
	}

	if err := row.Scan(&result); err != nil {
		return false, err
	}

	return result > 0, nil
}
