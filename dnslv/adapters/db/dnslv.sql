-- DOMAIN AGGREGATE

-- name: GetUnexpiredDomain :many
SELECT d.id     AS domain_id,
			 d.name   AS domain_name,
			 rr.id    AS resource_record_id,
			 rr.type  AS resource_record_type,
			 rr.value AS resource_record_value,
			 rr.ttl   AS resource_record_ttl
FROM domains d
			 LEFT JOIN resource_records rr ON d.id = rr.domain_id
WHERE NOW() - (INTERVAL '1 second' * @ttl::INT) <= d.updated_at;

-- name: DomainExists :one
SELECT id
FROM domains
WHERE name = @name;

-- name: CreateDomain :one
INSERT INTO domains (name)
VALUES (@name)
RETURNING *;

-- name: AddResourceRecord :copyfrom
INSERT INTO resource_records (domain_id, type, value, ttl)
VALUES (@domain_id, @type, @value, @ttl);
