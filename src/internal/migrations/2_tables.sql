-- +goose Up
CREATE TABLE member
(
    id integer not null primary key,
    uuid UUID DEFAULT uuid_generate_v4() NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    title varchar(20) NOT NULL,
    dob date,
    email varchar(255) NOT NULL,
    profile_url varchar(255),
    created_at timestamp(0) default NULL::timestamp without time zone
);

CREATE INDEX idx_member_uuid
    ON member (uuid);

CREATE UNIQUE INDEX uniq_26df0c148a90aba9
    ON member(uuid);

CREATE TABLE court
(
    id integer NOT NULL PRIMARY KEY,
    uuid UUID DEFAULT uuid_generate_v4() NOT NULL,
    court_number integer NOT NULL,
    alt_name varchar(255),
    surface varchar(20) NOT NULL
);

CREATE INDEX idx_court_uuid
    ON court (uuid);

CREATE UNIQUE INDEX uniq_26df0a9c148a90ab
    ON court(uuid);

-- +goose Down
DROP TABLE court;
DROP TABLE member;