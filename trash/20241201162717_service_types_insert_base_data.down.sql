DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'service_types') THEN
        DELETE FROM service_types WHERE name IN ('service', 'company');
    END IF;
END $$;