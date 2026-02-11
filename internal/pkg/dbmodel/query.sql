-- name: ListFiles :many
SELECT * FROM files
ORDER BY name
LIMIT 2;

-- name: AddFile :execresult
INSERT INTO files(name, content)
VALUES(@name, @content);