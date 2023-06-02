-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: CountRoleByID :one
SELECT COUNT(*)::integer FROM roles
WHERE id = $1;

-- name: CountRoleAll :one
SELECT COUNT(*)::integer FROM roles;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY name;

-- name: ListRolesPagination :many
SELECT * FROM roles
ORDER BY name LIMIT $1 OFFSET $2;

-- name: CreateRole :one
INSERT INTO roles (
    name
) VALUES (
             $1
         )
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;