package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shenikar/geo_broadcasting_system/internal/models"
)

// IncidentRepository определяет контракт для работы с бд инцидентов
type IncidentRepository interface {
	Create(ctx context.Context, incident *models.Incident) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Incident, error)
	Update(ctx context.Context, incident *models.Incident) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page, pageSize int) ([]*models.Incident, error)
	FindActiveByLocation(ctx context.Context, lat, lon float64) ([]*models.Incident, error)
}

// IncidentService определяет контрак для бизнес-логики управления инцидентами
type IncidentService interface {
	CreateIncident(ctx context.Context, incident *models.Incident) error
	GetIncident(ctx context.Context, id uuid.UUID) (*models.Incident, error)
	UpdateIncident(ctx context.Context, incident *models.Incident) error
	DeactivateIncident(ctx context.Context, id uuid.UUID) error
	ListIncident(ctx context.Context, page, pageSize int) ([]*models.Incident, error)
	CheckLocation(ctx context.Context, userID string, lat, lon float64) ([]*models.Incident, error)
}

type incidentService struct {
	repo IncidentRepository
}

func NewIncidentService(repo IncidentRepository) IncidentService {
	return &incidentService{
		repo: repo,
	}
}
