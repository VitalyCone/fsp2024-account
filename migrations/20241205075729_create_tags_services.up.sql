CREATE TABLE tags_services (
    tag_id INT NOT NULL,
    object_id INT NOT NULL,
    PRIMARY KEY (tag_id, object_id),
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    FOREIGN KEY (object_id) REFERENCES services(id) ON DELETE CASCADE
);