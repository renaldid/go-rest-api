package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		erorrCommit := tx.Commit()
		PanicIfError(erorrCommit)
	}
}
