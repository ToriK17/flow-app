CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT NOT NULL,
    price NUMERIC NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sales (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    sale_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE inventory_audit (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    expected_quantity INT,
    actual_quantity INT,
    noted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
