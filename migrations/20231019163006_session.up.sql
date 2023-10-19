CREATE TABLE IF NOT EXISTS session
(
    id          bigserial not null constraint session_pk primary key,
    usr_id     bigint                        unique         not null
        constraint session_fk_usr_id
            references usr
            on update cascade on delete cascade,
    token       varchar not null unique,
    expire_at   timestamp with time zone not null
);

CREATE INDEX session_index
ON session (token);