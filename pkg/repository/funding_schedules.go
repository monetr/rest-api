package repository

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/monetr/rest-api/pkg/models"
	"github.com/pkg/errors"
	"time"
)

func (r *repositoryBase) GetFundingSchedules(ctx context.Context, bankAccountId uint64) ([]models.FundingSchedule, error) {
	span := sentry.StartSpan(ctx, "GetFundingSchedules")
	defer span.Finish()

	span.Data = map[string]interface{}{
		"accountId":     r.AccountId(),
		"bankAccountId": bankAccountId,
	}

	result := make([]models.FundingSchedule, 0)
	err := r.txn.ModelContext(span.Context(), &result).
		Where(`"funding_schedule"."account_id" = ?`, r.AccountId()).
		Where(`"funding_schedule"."bank_account_id" = ?`, bankAccountId).
		Select(&result)
	if err != nil {
		span.Status = sentry.SpanStatusInternalError
		return nil, errors.Wrap(err, "failed to retrieve funding schedules")
	}

	span.Status = sentry.SpanStatusOK

	return result, nil
}

func (r *repositoryBase) GetFundingSchedule(ctx context.Context, bankAccountId, fundingScheduleId uint64) (*models.FundingSchedule, error) {
	span := sentry.StartSpan(ctx, "GetFundingSchedule")
	defer span.Finish()

	span.Data = map[string]interface{}{
		"accountId":         r.AccountId(),
		"bankAccountId":     bankAccountId,
		"fundingScheduleId": fundingScheduleId,
	}

	var result models.FundingSchedule
	err := r.txn.ModelContext(span.Context(), &result).
		Where(`"funding_schedule"."account_id" = ?`, r.AccountId()).
		Where(`"funding_schedule"."bank_account_id" = ?`, bankAccountId).
		Where(`"funding_schedule"."funding_schedule_id" = ?`, fundingScheduleId).
		Limit(1).
		Select(&result)
	if err != nil {
		span.Status = sentry.SpanStatusInternalError
		return nil, errors.Wrap(err, "could not retrieve funding schedule")
	}

	span.Status = sentry.SpanStatusOK

	return &result, nil
}

func (r *repositoryBase) CreateFundingSchedule(ctx context.Context, fundingSchedule *models.FundingSchedule) error {
	span := sentry.StartSpan(ctx, "CreateFundingSchedule")
	defer span.Finish()

	span.Data = map[string]interface{}{
		"accountId":     r.AccountId(),
		"bankAccountId": fundingSchedule.BankAccountId,
	}

	fundingSchedule.AccountId = r.AccountId()
	if _, err := r.txn.ModelContext(span.Context(), fundingSchedule).Insert(fundingSchedule); err != nil {
		span.Status = sentry.SpanStatusInternalError
		return errors.Wrap(err, "failed to create funding schedule")
	}

	span.Status = sentry.SpanStatusOK

	return nil
}

func (r *repositoryBase) UpdateNextFundingScheduleDate(ctx context.Context, fundingScheduleId uint64, nextOccurrence time.Time) error {
	span := sentry.StartSpan(ctx, "UpdateNextFundingScheduleDate")
	defer span.Finish()

	span.Data = map[string]interface{}{
		"accountId":         r.AccountId(),
		"fundingScheduleId": fundingScheduleId,
	}

	_, err := r.txn.ModelContext(span.Context(), &models.FundingSchedule{}).
		Set(`"last_occurrence" = "next_occurrence"`).
		Set(`"next_occurrence" = ?`, nextOccurrence).
		Where(`"funding_schedule"."account_id" = ?`, r.AccountId()).
		Where(`"funding_schedule"."funding_schedule_id" = ?`, fundingScheduleId).
		Update()
	if err != nil {
		span.Status = sentry.SpanStatusInternalError
		return errors.Wrap(err, "failed to set next occurrence")
	}

	span.Status = sentry.SpanStatusOK

	return nil
}
