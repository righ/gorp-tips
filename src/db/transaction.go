package db

import (
	"strconv"

	"github.com/go-gorp/gorp"
)

type NestableTx struct {
	*gorp.Transaction

	savePoint int
	next      *NestableTx
	resolved  bool
}

func (tx *NestableTx) Begin() (*NestableTx, error) {
	tx.next = &NestableTx{
		Transaction: tx.Transaction,
		savePoint:   tx.savePoint + 1,
	}
	if err := tx.Savepoint("SP" + strconv.Itoa(tx.next.savePoint)); err != nil {
		return nil, err
	}
	return tx.next, nil
}

func (tx *NestableTx) Rollback() error {
	tx.resolved = true
	if tx.savePoint > 0 {
		return tx.RollbackToSavepoint("SP" + strconv.Itoa(tx.savePoint))
	}
	return tx.Transaction.Rollback()
}

func (tx *NestableTx) Commit() error {
	if tx.next != nil && !tx.next.resolved {
		if err := tx.next.Commit(); err != nil {
			return err
		}
	}
	tx.resolved = true

	if tx.savePoint > 0 {
		return tx.ReleaseSavepoint("SP" + strconv.Itoa(tx.savePoint))
	}
	return tx.Transaction.Commit()
}
