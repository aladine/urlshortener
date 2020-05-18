CREATE TABLE url_hashes (
    id serial PRIMARY KEY,
    user_id bigint, 
    url text, 
    title text, 
    hash char(4) UNIQUE, 
    created timestamp
    )