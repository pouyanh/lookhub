-- DOMAINS QUERIES

-- name: GetUnexpiredDomain :one
SELECT *
FROM domains
WHERE NOW() - (@ttl::INT * INTERVAL '1 second') > updated_at;
