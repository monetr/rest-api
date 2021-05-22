package repository

import (
	"context"
	"github.com/monetrapp/rest-api/pkg/models"
	"github.com/pkg/errors"
)

func (r *repositoryBase) GetAccount() (*models.Account, error) {
	if r.account != nil {
		return r.account, nil
	}

	var account models.Account
	err := r.database.NewSelect().
		Model(&account).
		Where(`"account"."account_id" = ?`, r.AccountId()).
		Limit(1).
		Scan(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve account")
	}

	r.account = &account

	return r.account, nil
}
