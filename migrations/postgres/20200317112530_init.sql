-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "moddatetime";

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "moddatetime";