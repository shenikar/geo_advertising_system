package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shenikar/geo_broadcasting_system/internal/models"
	"github.com/shenikar/geo_broadcasting_system/internal/service"
)

type IncidentRepository struct {
	db *pgxpool.Pool
}

func NewIncidentRepository(db *pgxpool.Pool) service.IncidentRepository {
	return &IncidentRepository{
		db: db,
	}
}

// Create создает новую запись об инциденте в бд
func (r *IncidentRepository) Create(ctx context.Context, incident *models.Incident) error {
	query := `
		INSERT INTO incidents (name, description, location, radius_meters, status)
		VALUES ($1, $2, ST_SetSRID(ST_MakePoint($3, $4), 4326), $5, $6) RETURING id, created_at, updated_at;	
	`
	err := r.db.QueryRow(ctx, query,
		incident.Name,
		incident.Description,
		incident.Longitude,
		incident.Latitude,
		incident.RadiusMeters,
		incident.Status,
	).Scan(&incident.ID, &incident.CreatedAt, &incident.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create incident: %w", err)
	}
	return nil
}

// GetByID возвращает инцидент по его UUID
func (r *IncidentRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Incident, error) {
	incident := &models.Incident{}
	query := `
		SELECT 
			id,
			name,
			description,
			ST_Y(location::geometry) as latitude,
			ST_X(location::geometry) as longitude,
			radius_meters,
			status,
			created_at,
			updated_at
		FROM incidents
		WHERE id = $1;
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&incident.ID,
		&incident.Name,
		&incident.Description,
		&incident.Latitude,
		&incident.Longitude,
		&incident.RadiusMeters,
		&incident.Status,
		&incident.CreatedAt,
		&incident.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("incident with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get incident by id: %w", err)
	}
	return incident, nil
}
