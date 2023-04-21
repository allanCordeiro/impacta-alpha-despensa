CREATE TABLE IF NOT EXISTS stock_products (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    creation_date date NOT NULL,
    quantity int NOT NULL,
    expiration_date date NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );