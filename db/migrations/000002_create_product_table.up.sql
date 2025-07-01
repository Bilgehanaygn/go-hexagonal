CREATE TABLE product (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name        VARCHAR(255)      NOT NULL,
  price       NUMERIC(12,2)     NOT NULL,
  created_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  version     INTEGER           NOT NULL DEFAULT 1
); 