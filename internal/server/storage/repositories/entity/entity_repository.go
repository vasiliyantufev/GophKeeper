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

func (e *Entity) GetList(userID int64) ([]model.Entity, error) {
	entity := []model.Entity{}
	return entity, nil
}

func (e *Entity) Exists(entityRequest *model.CreateEntityRequest) (bool, error) {
	var exists bool
	row := e.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM entity "+
		"where entity.user_id = $1 and entity.metadata->>'name' = $2 and entity.metadata->>'type' = $3 and entity.deleted_at IS NULL)",
		entityRequest.UserID, entityRequest.Metadata.Name, entityRequest.Metadata.Type)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (e *Entity) Delete(userID int64, name string, typeEntity string) error {
	var id int64
	if err := e.db.Pool.QueryRow("UPDATE entity SET deleted_at = $1 "+
		"where entity.user_id = $2 and entity.metadata->>'name' = $3 and entity.metadata->>'type' = $4 and entity.deleted_at IS NULL",
		time.Now(),
		userID,
		name,
		typeEntity,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (e *Entity) Update(entityID int64, data []byte) error {
	//var id int64
	return nil
}
