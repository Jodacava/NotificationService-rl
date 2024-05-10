DROP TABLE IF EXISTS type_notification;
create table type_notification
(
    type_id    varchar not null
        constraint type_notification_pk
            primary key,
    rl_rule_id integer not null
);

create index type_notification_type_id_index
    on type_notification (type_id);

DROP TABLE IF EXISTS rate_limit_rule;
create table rate_limit_rule
(
    rl_rule_Id    integer not null
        constraint rate_limit_rule_pk
            primary key,
    max_shipments integer not null,
    time_shipment varchar not null
);

create index rate_limit_rule_time_shipment_index
    on rate_limit_rule (time_shipment desc);

DROP TABLE IF EXISTS notification_attempt;
create table notification_attempt
(
    email_recipient   varchar not null,
    type_id           varchar not null,
    shipment_count    integer not null,
    last_notification varchar not null,
    constraint notification_attempt_pk
        primary key (email_recipient, type_id)
);

alter table notification_attempt
    owner to postgres;

