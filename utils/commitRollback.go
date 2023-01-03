package utils

import (
	"gorm.io/gorm"
)

func CommitRollback(transaction *gorm.DB) {
	err := recover()
	if err != nil {
		errRollback := transaction.Rollback().Error
		PanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := transaction.Commit().Error
		PanicIfError(errCommit)
	}
}
