ALTER TABLE "card" DROP CONSTRAINT "card_account_id_fkey";

ALTER TABLE "address" DROP CONSTRAINT "address_users_id_fkey";

ALTER TABLE "transfer" DROP CONSTRAINT "transfer_destination_id_fkey";

ALTER TABLE "transfer" DROP CONSTRAINT "transfer_origin_id_fkey";

ALTER TABLE "deposit" DROP CONSTRAINT "deposit_account_id_fkey";

ALTER TABLE "account" DROP CONSTRAINT "account_users_id_fkey";

DROP INDEX "card_account_id_idx";

DROP INDEX "card_id_idx";

DROP INDEX "address_created_at_idx";

DROP INDEX "address_street_idx";

DROP INDEX "address_state_idx";

DROP INDEX "address_city_idx";

DROP INDEX "address_neighbor_idx";

DROP INDEX "address_users_id_idx";

DROP INDEX "address_id_idx";

DROP INDEX "transfer_created_at_idx";

DROP INDEX "transfer_amount_idx";

DROP INDEX "transfer_type_idx";

DROP INDEX "transfer_destination_id_origin_id_idx";

DROP INDEX "transfer_destination_id_idx";

DROP INDEX "transfer_origin_id_idx";

DROP INDEX "transfer_id_idx";

DROP INDEX "deposit_created_at_idx";

DROP INDEX "deposit_amount_idx";

DROP INDEX "deposit_deposit_uuid_idx";

DROP INDEX "deposit_id_idx";

DROP INDEX "users_created_at_idx";

DROP INDEX "users_email_idx";

DROP INDEX "users_cpf_idx";

DROP INDEX "users_id_idx";

DROP INDEX "account_created_at_idx";

DROP INDEX "account_users_id_idx";

DROP INDEX "account_id_idx";

DROP TABLE "card";

DROP TABLE "address";

DROP TABLE "transfer";

DROP TABLE "deposit";

DROP TABLE "account";

DROP TABLE "users";

DROP TYPE "transfer_type";
