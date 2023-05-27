package metadata

import (
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/database"
	"github.com/vasiliyantufev/gophkeeper/internal/model"
)

type Metadata struct {
	db *database.DB
}

func New(db *database.DB) *Metadata {
	return &Metadata{
		db: db,
	}
}

func (m *Metadata) CreateMetadata(metadataRequest *model.CreateMetadataRequest) (*model.Metadata, error) {
	metadata := &model.Metadata{}
	if err := m.db.Pool.QueryRow(
		"INSERT INTO metadata (name, description, created_at, updated_at) "+
			"VALUES ($1, $2, $3, $4) RETURNING metadata_id, name, description",
		metadataRequest.Name,
		metadataRequest.Description,
		time.Now(),
		time.Now(),
	).Scan(&metadata.ID, &metadata.Name, &metadata.Description); err != nil {
		return nil, err
	}
	return metadata, nil
}

//func (m *Metadata) GetMetadata(metadataRequest *model.GetMetadataRequest) (*model.Metadata, error) {
//	metadata := &model.Metadata{}
//	err := m.db.Pool.QueryRow("SELECT name, description FROM metadata WHERE metadata_id=$1",
//		metadataRequest.MetadataId).Scan(
//		&metadata.Name,
//		&metadata.Description,
//	)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, errors.ErrRecordNotFound
//		} else {
//			return nil, err
//		}
//	}
//	return metadata, nil
//}
