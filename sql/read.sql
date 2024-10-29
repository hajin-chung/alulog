-- read logs
SELECT * FROM log;

-- read post list
SELECT id, created_at, updated_at, title, subtitle FROM post;

-- read post of id
SELECT id, created_at, updated_at, title, subtitle, text FROM post WHERE id = $id;
