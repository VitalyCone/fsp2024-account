DROP INDEX IF EXISTS idx_review_services_id;
DROP TRIGGER IF EXISTS set_updated_at_review_services ON review_services;
DROP FUNCTION IF EXISTS update_updated_at_column_review_services();
DROP TABLE IF EXISTS review_services;