CREATE TABLE categories (
  id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name              VARCHAR(255) NOT NULL,
  kind              VARCHAR(50)  NOT NULL,
  status            VARCHAR(20)  NOT NULL,
  created_date      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  last_modified_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);