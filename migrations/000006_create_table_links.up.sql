CREATE TABLE links (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  description TEXT,
  tags VARCHAR[],
  category_ids INTEGER[],
  user_id INTEGER NOT NULL,
  status SMALLINT NOT NULL REFERENCES statuses(id),
  is_public BOOLEAN DEFAULT TRUE
);

CREATE INDEX idx_links_tags ON links USING GIN (tags);
CREATE INDEX idx_links_category_ids ON links USING GIN (category_ids);