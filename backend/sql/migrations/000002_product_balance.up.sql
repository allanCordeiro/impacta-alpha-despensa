CREATE TABLE IF NOT EXISTS product_balance (
    id int PRIMARY KEY NOT NULL,
    product_id varchar(255) NOT NULL,
    deducted_amount int NOT NULL,
    deducted_date timestamp NOT NULL
);