  CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    body TEXT NOT NULL,
    title TEXT NOT NULL
);

INSERT INTO tasks (title, body)
VALUES ('pgabc', 'pgdef'),
('pghello', 'pgworld'),
('post', 'gresql'),
('abcdef', 'postgres');

