CREATE TABLE IF NOT EXISTS todos (
    id          SERIAL  not null    constraint  todos_pkey   primary key,
    title       VARCHAR not null,
    description VARCHAR,
    is_done     BOOL    not null    default false,
    user_id     INT     not null,
    time_stamp  TIMESTAMPTZ
);