CREATE TABLE api_keys (
    id UUID PRIMARY KEY,
    project_id UUID NOT NULL,
    prefix TEXT NOT NULL,
    key_hash TEXT NOT NULL,
    scopes TEXT[] NOT NULL,
    environment TEXT NOT NULL,
    revoked BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),
    expires_at TIMESTAMP
);