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
	domain_id  INT REFERENCES domains (id) ON DELETE CASCADE ON UPDATE CASCADE,
	id         SERIAL                  NOT NULL,
	type       VARCHAR(255)            NOT NULL,
	value      TEXT                    NOT NULL,
	ttl        INT                     NOT NULL,

	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
	PRIMARY KEY (domain_id, id)
);

CREATE INDEX idx_resource_records_domain_id ON resource_records (domain_id);

CREATE INDEX idx_resource_records_ttl ON resource_records (ttl);

CREATE INDEX idx_resource_records_updated_at ON resource_records (updated_at);

-- Resource record domain-scoped id

CREATE OR REPLACE FUNCTION reset_record_id()
	RETURNS TRIGGER AS
$$
BEGIN
	-- Reset the sequence if all records for the domain are deleted
	IF NOT EXISTS (SELECT 1 FROM resource_records WHERE domain_id = OLD.domain_id) THEN
		PERFORM setval(pg_get_serial_sequence('resource_records', 'id'), 1, false);
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_delete_reset_record_id
	AFTER DELETE
	ON resource_records
	FOR EACH ROW
EXECUTE FUNCTION reset_record_id();
