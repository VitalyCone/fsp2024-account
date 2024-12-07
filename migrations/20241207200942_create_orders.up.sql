CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    company_id INT NOT NULL,
    service_id INT NOT NULL,
    order_status VARCHAR(50) NOT NULL,
    price FLOAT NOT NULL,
    will_be_finished_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (username) REFERENCES users(username),
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (service_id) REFERENCES services(id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column_orders()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание триггера, который будет вызываться перед обновлением записи
CREATE TRIGGER set_updated_at_orders
BEFORE UPDATE ON orders
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_orders();