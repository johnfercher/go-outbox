package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/rand"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddNew() error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		anyTable := &AnyTable{
			ID:      id.String(),
			AnyData: fmt.Sprintf("%d", rand.Int()),
		}

		err = tx.Create(anyTable).Error
		if err != nil {
			return err
		}

		outboxID, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		outbox := &Outbox{
			ID:        outboxID.String(),
			TableID:   id.String(),
			TableName: "any_tables",
			Status:    "created",
		}

		return tx.Create(outbox).Error
	})
}
