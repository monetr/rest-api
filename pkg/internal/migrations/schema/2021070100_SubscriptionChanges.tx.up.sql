ALTER TABLE "accounts" ADD COLUMN "stripe_customer_id" TEXT NULL;
ALTER TABLE "accounts" ADD CONSTRAINT "uq_accounts_stripe_customer_id" UNIQUE ("stripe_customer_id");

ALTER TABLE "accounts" ADD COLUMN "stripe_subscription_id" TEXT NULL;
ALTER TABLE "accounts" ADD CONSTRAINT "uq_accounts_stripe_subscription_id" UNIQUE ("stripe_subscription_id");

ALTER TABLE "accounts" ADD COLUMN "subscription_active_until" TIMESTAMPTZ NULL;

DROP TABLE "subscriptions";

-- Just clean up from an old idea.
DELETE FROM "users" WHERE user_id = -1;
DELETE FROM "logins" WHERE login_id = -1;
DELETE FROM "accounts" WHERE account_id = -1;