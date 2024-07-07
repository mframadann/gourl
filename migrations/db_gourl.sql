CREATE TABLE tb_users (
    user_id SERIAL PRIMARY KEY,
    pub_id VARCHAR(40) UNIQUE,
    first_name VARCHAR(20),
    last_name VARCHAR(20) DEFAULT NULL,
    email_address VARCHAR(30) UNIQUE,
    password VARCHAR,
    registered_at DATE
);


CREATE TABLE tb_groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR
);


CREATE TABLE tb_links (
    link_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES tb_users(user_id),
    group_id INT REFERENCES tb_groups(group_id),
    link_title VARCHAR(30)
    shorted_url VARCHAR UNIQUE,
    origin_url VARCHAR,
    created_at DATE,
    updated_at TIMESTAMP
);
