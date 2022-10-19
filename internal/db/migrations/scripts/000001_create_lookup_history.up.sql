CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS lookup_history
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    domain varchar NOT NULL,
    response jsonb NULL,
    created_at int NOT NULL
);

CREATE INDEX IF NOT EXISTS lookup_history_created_at_idx ON lookup_history USING btree (created_at);