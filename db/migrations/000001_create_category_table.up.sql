CREATE TABLE category (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name        VARCHAR(255)      NOT NULL,
  kind        VARCHAR(50)       NOT NULL,
  status      VARCHAR(20)       NOT NULL,
  parent_category_id  UUID      REFERENCES category(id),
  created_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  version     INTEGER           NOT NULL DEFAULT 1
);
