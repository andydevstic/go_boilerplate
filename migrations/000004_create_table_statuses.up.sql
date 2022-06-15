CREATE TABLE statuses (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

INSERT INTO statuses (name) VALUES ('Inactive'), ('Active'), ('Deleted');