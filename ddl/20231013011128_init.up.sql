do
$$
begin
    execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
end;
$$;


CREATE TABLE IF NOT EXISTS usr
(
    id           bigserial not null constraint usr_pk primary key,
    name        VARCHAR NOT NULL,
    nickname    VARCHAR NOT NULL,
    secret      VARCHAR NOT NULL
);


CREATE TABLE IF NOT EXISTS footprint
(
    id         bigserial not null constraint footprint_pk primary key,
    usr_id     bigint                        unique         not null
        constraint footprint_fk_usr_id
            references usr
            on update cascade on delete cascade,
    easy       bigint   not null,
    medium     bigint   not null,
    hard       bigint   not null,
    total      bigint   not null,
    active     boolean default false
);

CREATE TABLE IF NOT EXISTS event
(
    id         bigserial not null constraint event_pk primary key,
    usr_id     bigint                                 not null
        constraint event_fk_usr_id
            references usr
            on update cascade on delete cascade,
    goal        smallint not null,
    category    varchar  not null,
    tags        varchar  not null,
    status_id   smallint not null,
    start_time  timestamp with time zone  not null,
    end_time    timestamp with time zone  not null
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
    active      boolean    default true
)