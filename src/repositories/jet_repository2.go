package repositories

import (
	"context"

	"github.com/go-gorp/gorp"
	"github.com/labstack/gommon/log"

	"gorp-tips/db"
	"gorp-tips/models"
)

type jetRepository2 struct {
	exec gorp.SqlExecutor
}

func NewJetRepository2(exec gorp.SqlExecutor) JetRepository {
	return &jetRepository2{
		exec: exec,
	}
}

func (r *jetRepository2) GetJets(ctx context.Context, req models.Request) ([]models.Result, error) {
	query := db.GetSQL2("query.sql", req)
	log.Debug(query)

	var results []models.Result
	if _, err := r.exec.Select(&results, query, map[string]interface{}{
		"age":        req.Age,
		"pilot_name": "%" + req.PilotName + "%",
		"jet_name":   "%" + req.JetName + "%",
		"language":   req.Language,
	}); err != nil {
		log.Error(err)
		return nil, err
	}
	return results, nil
}
