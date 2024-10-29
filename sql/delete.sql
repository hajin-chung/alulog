-- delete log of id
DELETE FROM log WHERE id = $id;

-- delete log of post
DELETE FROM post WHERE id = $id;
