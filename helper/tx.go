package helper

import (
	"gorm.io/gorm"
)

//func CommitOrRollback(tx *sql.Tx) {
//	err := recover()
//	if err != nil {
//		errorRollback := tx.Rollback()
//		PanicIfError(errorRollback)
//		panic(err)
//	} else {
//		errorCommit := tx.Commit()
//		PanicIfError(errorCommit)
//	}
//}

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		if errorRollback != nil && errorRollback != gorm.ErrRecordNotFound {
			// Rollback failed, log or handle the error accordingly
			PanicIfError(errorRollback)
		}
		panic(err)
	} else {
		errorCommit := tx.Commit().Error
		if errorCommit != nil && errorCommit != gorm.ErrRecordNotFound {
			// Commit failed, log or handle the error accordingly
			PanicIfError(errorCommit)
		}
	}
}
