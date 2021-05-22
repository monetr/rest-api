//+build !vault

package repository

import (
	"github.com/monetrapp/rest-api/pkg/models"
	"github.com/uptrace/bun"
)

type repositoryBase struct {
	userId, accountId uint64
	database          *bun.DB
	account           *models.Account
}
