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

func (e *Entity) CreateEntity(entityRequest *model.CreateEntityRequest) (*model.Entity, error) {
	entity := &model.Entity{}
	return entity, nil
}

func (e *Entity) GetListEntity(userId int64) ([]model.Entity, error) {
	entity := []model.Entity{}
	return entity, nil
}

func (e *Entity) EntityExists(entityRequest *model.CreateEntityRequest) (bool, error) {
	var exists bool
	return exists, nil
}

func (e *Entity) DeleteEntity(EntityID int64) error {
	//var id int64
	return nil
}

func (e *Entity) UpdateEntity(entityID int64, data []byte) error {
	//var id int64
	return nil
}
