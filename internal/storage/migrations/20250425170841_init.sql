-- +goose Up
-- +goose StatementBegin
create table buckets(
    id serial primary key,
    name text not null,
    description text
);

create table problem_details(
    id serial primary key,
    description text not null,
    type text unique,
    title text,
    status integer,
    detail text,
    instance text
);

create table bucket_problem_details(
   id serial primary key,
   bucket_id integer REFERENCES buckets (id) NOT NULL,
   problem_detail_id integer REFERENCES problem_details (id) NOT NULL,

   unique (bucket_id, problem_detail_id)
);

create table labels(
    id serial primary key,
    name text not null unique
);

create table problem_details_labels(
    id serial primary key,
    label_id integer REFERENCES labels (id) NOT NULL,
    problem_detail_id integer REFERENCES problem_details (id) NOT NULL,

    unique (label_id, problem_detail_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table problem_details_labels;
drop table bucket_problem_details;
drop table labels;
drop table problem_details;
drop table buckets;
-- +goose StatementEnd
