CREATE TABLE users (
    id bigserial primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);

CREATE TABLE estimations (
    date_create date not null,
    user_id int not null,
    estimate int
);

ALTER TABLE "estimations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");