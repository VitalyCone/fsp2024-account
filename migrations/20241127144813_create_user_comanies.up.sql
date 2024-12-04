CREATE TABLE user_company_members (
    username VARCHAR(32) NOT NULL,
    company_id INT NOT NULL,
    PRIMARY KEY (username, company_id),
    FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);

-- Добавление индексов для улучшения производительности
CREATE INDEX idx_user_members_username ON user_company_members(username);
CREATE INDEX idx_company_members_id ON user_company_members(company_id);