package service

import "passos/internal/database"

type HealthStatusProvider interface {
	Status() map[string]string
}

type HealthService struct {
	db database.Service
}

func NewHealthService(db database.Service) *HealthService {
	return &HealthService{db: db}
}

func (s *HealthService) Status() map[string]string {
	if s == nil || s.db == nil {
		return map[string]string{"status": "unknown"}
	}
	return s.db.Health()
}
