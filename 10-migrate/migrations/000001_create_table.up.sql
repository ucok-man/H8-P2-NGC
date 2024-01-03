CREATE TABLE users (
    user_id serial PRIMARY KEY,
    email varchar(255) NOT NULL,
    password_hash bytea NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NOT NULL
);

CREATE TABLE profiles (
    profile_id serial PRIMARY KEY,
    user_id integer NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    phone_number varchar(255) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);