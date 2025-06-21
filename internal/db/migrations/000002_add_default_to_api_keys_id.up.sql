ALTER TABLE api_keys
ALTER COLUMN id SET DEFAULT gen_random_uuid();