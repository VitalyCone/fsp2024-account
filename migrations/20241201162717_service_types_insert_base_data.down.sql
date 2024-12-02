DELETE FROM service_types WHERE name IN ('service', 'company')
AND EXISTS (SELECT 1 FROM service_types WHERE name IN ('service', 'company'));