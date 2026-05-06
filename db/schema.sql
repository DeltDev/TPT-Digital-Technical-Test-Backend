-- tabel data produk
CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    price BIGINT NOT NULL CHECK (price>=0),
    stock INTEGER NOT NULL CHECK (stock>=0),
    category VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now()
);

-- seeding data dummy
INSERT INTO products (name, description, price, stock, category, is_active)
VALUES
('Indomie Goreng', 'Instant noodle classic', 3500, 100, 'Food', true),
('Mechanical Keyboard', 'RGB, blue switch', 750000, 25, 'Electronics', true),
('Gaming Mouse', '16000 DPI sensor', 250000, 40, 'Electronics', true),
('Office Chair', 'Ergonomic mesh chair', 1200000, 10, 'Furniture', true),
('Notebook', 'A5 lined notebook', 15000, 200, 'Stationery', true);