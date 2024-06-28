-- +goose Up
CREATE SEQUENCE member_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
CREATE TABLE member
(
    id integer NOT NULL PRIMARY KEY DEFAULT nextval('member_id_seq'),
    uuid UUID DEFAULT uuid_generate_v4() NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    title varchar(20) NOT NULL,
    dob date,
    email varchar(255) NOT NULL,
    profile_url varchar(255),
    created_at timestamp(0) default NULL::timestamp without time zone
);
ALTER SEQUENCE member_id_seq OWNED BY member.id;
CREATE UNIQUE INDEX uniq_member_uuid ON member(uuid);


CREATE SEQUENCE court_slot_set_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
CREATE TABLE court_slot_set
(
    id integer NOT NULL PRIMARY KEY DEFAULT nextval('court_slot_set_id_seq'),
    uuid UUID DEFAULT uuid_generate_v4() NOT NULL,
    name varchar(255) NOT NULL,
    description text,
    slot_times jsonb DEFAULT '[]'
);
ALTER SEQUENCE court_slot_set_id_seq OWNED BY court_slot_set.id;
CREATE UNIQUE INDEX uniq_court_slot_set_uuid ON court_slot_set(uuid);


CREATE SEQUENCE court_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
CREATE TABLE court
(
    id integer NOT NULL PRIMARY KEY DEFAULT nextval('court_id_seq'),
    uuid UUID DEFAULT uuid_generate_v4() NOT NULL,
    court_number integer NOT NULL,
    alt_name varchar(255),
    surface varchar(20) NOT NULL,
    slot_set_id UUID NOT NULL,
    CONSTRAINT fk_court_court_slot_set
        FOREIGN KEY(slot_set_id)
            REFERENCES court_slot_set(uuid)
);
ALTER SEQUENCE court_id_seq OWNED BY court.id;
CREATE UNIQUE INDEX uniq_court_uuid ON court(uuid);
CREATE INDEX idx_court_slot_set_id on court(slot_set_id);

-- +goose Down
DROP TABLE member;
DROP TABLE court;
DROP TABLE court_slot_set;
