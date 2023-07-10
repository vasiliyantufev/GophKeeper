package entity

import (
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/model"
)

type Entity struct {
	db *database.DB
}

func New(db *database.DB) *Entity {
	return &Entity{
		db: db,
	}
}

func (e *Entity) Create(entityRequest *model.CreateEntityRequest) (*model.Entity, error) {
	entity := &model.Entity{}
	return entity, nil
}

func (e *Entity) GetList(userId int64) ([]model.Entity, error) {
	entity := []model.Entity{}
	return entity, nil
}

func (e *Entity) Exists(entityRequest *model.CreateEntityRequest) (bool, error) {
	var exists bool
	return exists, nil
}

func (e *Entity) Delete(EntityID int64) error {
	//var id int64
	return nil
}

func (e *Entity) Update(entityID int64, data []byte) error {
	//var id int64
	return nil
}
