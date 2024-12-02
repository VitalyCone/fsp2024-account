CREATE TABLE review_services(
    id SERIAL PRIMARY KEY,
    object_id INT NOT NULL,
    rating INT NOT NULL,
    creator_username VARCHAR(32) NOT NULL,  -- Assuming you have a user table
    header VARCHAR(255),
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_services_object_id FOREIGN KEY (object_id) REFERENCES services(id) ON DELETE CASCADE,
    CONSTRAINT fk_creator_username FOREIGN KEY (creator_username) REFERENCES users(username) ON DELETE CASCADE
);

CREATE INDEX idx_review_services_id ON review_services(id);

CREATE OR REPLACE FUNCTION update_updated_at_column_review_services()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
 LANGUAGE plpgsql;

-- Создание триггера, который будет вызываться перед обновлением записи в таблице reviews
CREATE TRIGGER set_updated_at_review_services
BEFORE UPDATE ON review_services
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_review_services();