DROP INDEX IF EXISTS idx_review_companies_id;
DROP INDEX IF EXISTS idx_review_companies_id;
DROP TRIGGER IF EXISTS set_updated_at_review_companies ON review_companies;
DROP FUNCTION IF EXISTS update_updated_at_column_review_companies();
DROP TABLE IF EXISTS review_companies;