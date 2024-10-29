CREATE TABLE IF NOT EXISTS log (
  id TEXT PRIMARY_KEY,
  created_at TEXT NOT NULL,
  updated_at TEXT,
  content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS post (
  id TEXT PRIMARY_KEY,
  created_at TEXT NOT NULL,
  updated_at TEXT,
  title TEXT NOT NULL,
  subtitle TEXT,
  content TEXT NOT NULL
);

/*
CREATE TABLE IF NOT EXISTS tag (
  post_id TEXT NOT NULL,
  tag TEXT
);
*/

CREATE TABLE IF NOT EXISTS metadata (
  key TEXT PRIMARY KEY,
  value TEXT
);
