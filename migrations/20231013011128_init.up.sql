do
$$
begin
    execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
end;
$$;


CREATE TABLE IF NOT EXISTS usr
(
    id          bigserial not null constraint usr_pk primary key,
    name        VARCHAR(500) NOT NULL,
    username    VARCHAR(500) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR NOT NULL,
    u_group     TEXT,
    status_id   SMALLINT,
    type_id     SMALLINT
);


CREATE TABLE IF NOT EXISTS score
(
    id         bigserial not null constraint score_pk primary key,
    usr_id     bigint                        unique         not null
        constraint score_fk_usr_id
            references usr
            on update cascade on delete cascade,
    current      jsonb default '{"easy":0,"medium":0,"hard":0,"total":0}'::jsonb not null ,
    footprint    jsonb default '{"easy":0,"medium":0,"hard":0,"total":0}'::jsonb not null,
    active       boolean default false
);

CREATE TABLE IF NOT EXISTS event
(
    id          bigserial not null constraint event_pk primary key,
    usr_id     bigint                              not null
        constraint event_fk_usr_id
            references usr
            on update cascade on delete cascade,
    start_time timestamp with time zone not null,
    end_time    timestamp with time zone not null,
    goal        integer  not null,
    condition   text     not null,
    status_id   integer  not null
);

CREATE TABLE IF NOT EXISTS usr_event
(
    id          bigserial not null constraint usr_event_pk primary key,
    usr_id      bigint                                 not null
        constraint usr_event_fk_usr_id
            references usr
            on update cascade on delete cascade,
    event_id    bigint                                 not null
        constraint usr_event_fk_event_id
            references event
            on update cascade on delete cascade,
    active      boolean    default true,
    winner      boolean default false
);

CREATE TABLE IF NOT EXISTS tournament
(
    id          bigserial not null constraint tournament_pk primary key,
    usr_id     bigint                        unique         not null
        constraint tournament_fk_usr_id
            references usr
            on update cascade on delete cascade,
    start_time timestamp with time zone not null,
    end_time    timestamp with time zone not null,
    goal        integer  not null,
    cost        integer  not null,
    prize_pool  integer,
    status_id   integer  not null
);

CREATE TABLE IF NOT EXISTS usr_tournament
(
    id          bigserial not null constraint usr_tournament_pk primary key,
    usr_id      bigint                                 not null
        constraint usr_tournament_fk_usr_id
            references usr
            on update cascade on delete cascade,
    event_id    bigint                                 not null
        constraint usr_tournament_fk_event_id
            references event
            on update cascade on delete cascade,
    active      boolean    default true,
    winner      boolean    default false
);