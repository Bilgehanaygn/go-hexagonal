CREATE TABLE catalog (
  id          UUID PRIMARY KEY,
  name        VARCHAR(255)      NOT NULL,
  created_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  version     INTEGER           NOT NULL DEFAULT 1
); 