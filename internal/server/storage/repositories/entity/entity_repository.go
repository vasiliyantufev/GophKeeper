package entity

import (
	"encoding/json"
	"time"

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

func (e *Entity) Create(entityRequest *model.CreateEntityRequest) (int64, error) {
	var id int64

	metadata := model.MetadataEntity{Name: entityRequest.Metadata.Name, Description: entityRequest.Metadata.Description, Type: entityRequest.Metadata.Type}
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		return 0, err
	}
	if err = e.db.Pool.QueryRow(
		"INSERT INTO entity (user_id, data, metadata, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING entity_id",
		entityRequest.UserID,
		entityRequest.Data,
		jsonMetadata,
		time.Now(),
		time.Now(),
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
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
