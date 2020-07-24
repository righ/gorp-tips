package repositories

import (
	"context"
	"strings"

	"github.com/go-gorp/gorp"
	"github.com/labstack/gommon/log"

	"gorp-tips/models"
)

type JetRepository interface {
	GetJets(ctx context.Context, req models.Request) ([]models.Result, error)
}

type jetRepository struct {
	exec gorp.SqlExecutor
}

func NewJetRepository(exec gorp.SqlExecutor) JetRepository {
	return &jetRepository{
		exec: exec,
	}
}

func (r *jetRepository) GetJets(ctx context.Context, req models.Request) ([]models.Result, error) {
	query := "SELECT jets.name AS jetName, jets.age AS jetAge, jets.color AS jetColor, pilots.name AS pilotName, languages.language "
	query += "FROM jets "
	query += "JOIN pilots ON pilots.id = jets.pilot_id "
	query += "LEFT JOIN pilot_languages ON pilot_languages.pilot_id = jets.pilot_id "
	query += "LEFT JOIN languages ON languages.id = pilot_languages.language_id "
	conds, variables := makeCondition(req)
	if conds != "" {
		query += "WHERE " + conds
	}
	query += " ORDER BY jets.age, jets.id"
	log.Debug(query)

	var results []models.Result
	if _, err := r.exec.Select(&results, query, variables); err != nil {
		log.Error(err)
		return nil, err
	}
	return results, nil
}

func makeCondition(req models.Request) (string, map[string]interface{}) {
	conds := []string{}
	context := map[string]interface{}{}

	if req.Age > 0 {
		conds = append(conds, "jets.age = :age")
		context["age"] = req.Age
	}
	if req.PilotName != "" {
		conds = append(conds, "pilots.name LIKE :pilot_name")
		context["pilot_name"] = "%" + req.PilotName + "%"
	}
	if req.JetName != "" {
		conds = append(conds, "jets.name LIKE :jet_name")
		context["jet_name"] = "%" + req.JetName + "%"
	}
	if req.Language != "" {
		conds = append(conds, "languages.language = :language")
		context["language"] = req.Language
	}

	return strings.Join(conds, " AND "), context
}
