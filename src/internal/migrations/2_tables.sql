-- +goose Up
CREATE TABLE member
(
    id integer not null primary key,
    uuid UUID DEFAULT uuid_generate_v4() NOT NULL,
    first_name varchar(255),
    last_name varchar(255),
    title varchar(20),
    dob date,
    email varchar(255),
    profile_url varchar(255),
    created_at timestamp(0) default NULL::timestamp without time zone
);

CREATE INDEX idx_member_uuid
    ON member (uuid);

CREATE UNIQUE INDEX uniq_26df0c148a90aba9
    ON member(uuid);

-- +goose Down
DROP TABLE member;