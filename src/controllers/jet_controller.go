package controllers

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	"gorp-tips/models"
	"gorp-tips/repositories"
)

func GetJets(ctx context.Context, repo repositories.JetRepository, req models.Request) ([]models.Result, error) {
	return repo.GetJets(ctx, req)
}
