CREATE TABLE catalog_product (
  id          UUID PRIMARY KEY,
  catalog_id   UUID  NOT NULL,
  product_id   UUID  NOT NULL,
  price  DECIMAL(10,2)  NOT NULL,
  created_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ       NOT NULL DEFAULT NOW(),
  version     INTEGER           NOT NULL DEFAULT 1,
  FOREIGN KEY (catalog_id) REFERENCES catalog(id),
  FOREIGN KEY (product_id) REFERENCES product(id)
);