package helper

import "database/sql"

func CommitorRollback(tx *sql.Tx) {
	error := recover()
	if error != nil {
		errorWhenRollback := tx.Rollback()
		PanicIfError(errorWhenRollback)
		panic(error)
	} else {
		errorWhenCommitted := tx.Commit()
		PanicIfError(errorWhenCommitted)
	}
}
