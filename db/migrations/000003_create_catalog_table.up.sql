CREATE TABLE catalog (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name        VARCHAR(255)      NOT NULL,
  products    UUID[]            NOT NULL,
  created_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  version     INTEGER           NOT NULL DEFAULT 1
); 