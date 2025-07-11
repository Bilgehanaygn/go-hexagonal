ALTER TABLE product
  ADD COLUMN status VARCHAR(10);

UPDATE product
  SET status = 'Active'
  WHERE status IS NULL;

ALTER TABLE product
  ALTER COLUMN status SET NOT NULL;