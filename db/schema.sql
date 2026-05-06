CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    price BIGINT NOT NULL CHECK (price>=0),
    stock INTEGER NOT NULL CHECK (stock>=0),
    category VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
)