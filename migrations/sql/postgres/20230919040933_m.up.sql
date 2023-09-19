-- Create "settings" table
CREATE TABLE "public"."settings" (
  "id" bigserial NOT NULL,
  "name" character varying(255) NULL,
  "value" jsonb NULL DEFAULT '{}',
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_settings_deleted_at" to table: "settings"
CREATE INDEX "idx_settings_deleted_at" ON "public"."settings" ("deleted_at");
-- Create index "where:deleted_at IS NULL" to table: "settings"
CREATE UNIQUE INDEX "where:deleted_at IS NULL" ON "public"."settings" ("name");
