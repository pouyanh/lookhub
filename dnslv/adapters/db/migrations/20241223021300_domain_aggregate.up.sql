-- Domains table
CREATE TABLE domains
(
	id         SERIAL PRIMARY KEY,
	name       TEXT                    NOT NULL,

	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL,

	CONSTRAINT uq_name UNIQUE (name)
);

CREATE INDEX idx_domains_name_hash ON domains USING HASH (name);

CREATE INDEX idx_domains_updated_at ON domains (updated_at);

-- Resource records table
CREATE TABLE resource_records
(
	id         SERIAL PRIMARY KEY,
	domain_id  INT REFERENCES domains (id) ON DELETE CASCADE ON UPDATE CASCADE,
	type       VARCHAR(255)            NOT NULL,
	value      TEXT                    NOT NULL,
	ttl        INT                     NOT NULL,

	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_resource_records_domain_id ON resource_records (domain_id);

CREATE INDEX idx_resource_records_ttl ON resource_records (ttl);

CREATE INDEX idx_resource_records_updated_at ON resource_records (updated_at);
