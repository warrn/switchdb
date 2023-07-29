-- +goose Up
-- +goose StatementBegin
CREATE TABLE manufacturers (
    name varchar(63) PRIMARY KEY
);

CREATE TYPE model_power AS ENUM (
    'Fixed Internal',
    'Replaceable Internal',
    'Replaceable Redundant Internal',
    'DC/External',
    'POE'
);

CREATE TABLE models (
    id public.xid PRIMARY KEY DEFAULT xid(),
    model_number varchar(127) NOT NULL,
    manufacturer_name varchar(63) NOT NULL references manufacturers(name),

    date_released date NOT NULL,
    date_end_of_life date,
    date_end_of_support date,

    supported_power model_power[] NOT NULL,

    poe_watts_budget int NOT NULL default 0,

    CHECK ( array_length(supported_power, 1) > 0 ),
    UNIQUE (manufacturer_name, model_number)
);

CREATE INDEX ON models USING gin (supported_power);

CREATE TYPE interface_speed AS ENUM (
    '10M',
    '100M',
    '1000M',
    '2500M',
    '5000M',
    '10G',
    '25G',
    '40G',
    '50G',
    '100G',
    '200G',
    '400G'
);

CREATE TYPE interface_form_factor AS ENUM(
    'Base-T',
    'SFP',
    'SFP+',
    'SFP28',
    'QSFP+',
    'QSFP28',
    'USB-A',
    'USB-B Mini',
    'USB-B Micro',
    'RJ45',
    'DB-9'
);

CREATE TYPE interface_protocol AS ENUM (
    'Ethernet',
    'Serial',
    'USB'
);

CREATE TABLE interface_bays (
    model_id public.xid NOT NULL references models(id),
    bay_number int NOT NULL,
    ports int NOT NULL,
    supported_speeds interface_speed[] NOT NULL,
    form_factors interface_form_factor[] NOT NULL,
    protocol interface_protocol NOT NULL,
    management bool NOT NULL,

    poe bool NOT NULL,
    poe_watts_per_port int NOT NULL,
    poe_watts_bay_max int NOT NULL,

    CHECK ( array_length(form_factors, 1) > 0 ),
    UNIQUE (model_id, bay_number)
);

CREATE INDEX ON interface_bays USING gin (supported_speeds);
CREATE INDEX ON interface_bays USING gin (form_factors);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS interface_bays;
DROP TYPE IF EXISTS interface_protocol;
DROP TYPE IF EXISTS interface_form_factor;
DROP TYPE IF EXISTS interface_speed;
DROP TABLE IF EXISTS models;
DROP TABLE IF EXISTS manufacturers;
-- +goose StatementEnd
