package dao

import (
	"log"
)

// DAO wraps all data access objects.
type DAO struct {
	Account *Account
	Prices  *Prices
	Kv      *KV
}

// New returns a new instance of DAO.
func New() *DAO {
	return &DAO{
		Account: &Account{},
		Prices:  &Prices{},
		Kv:      &KV{},
	}
}

// BatchInsert inserts a batch of records into mysql in a transaction.
func (*DAO) BatchInsert(obj []interface{}) error {
	if len(obj) == 0 {
		return nil
	}
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Println("Batch insert panic: ", r)
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	for _, o := range obj {
		if err := tx.Create(o).Error; err != nil {
			tx.Rollback()
			return err
		}

	}

	return tx.Commit().Error
}
