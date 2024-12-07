DROP TRIGGER IF EXISTS set_updated_at_orders ON orders;

DROP FUNCTION IF EXISTS update_updated_at_column_orders();

DROP TABLE IF EXISTS orders;