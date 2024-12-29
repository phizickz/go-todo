  CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL
);

INSERT INTO tasks (title, body)
VALUES ('pgabc', 'pgdef'),
('pghello', 'pgworld'),
('post', 'gresql'),
('abcdef', 'postgres');

