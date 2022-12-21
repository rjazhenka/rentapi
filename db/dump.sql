create table if not exists rent_turkey
(
    id                serial,
    title             varchar,
    rooms             integer,
    price             numeric,
    country           integer,
    city              integer,
    region            integer,
    district          integer,
    description       text,
    link              varchar,
    source            integer,
    status            integer,
    images_tg_ids     json,
    images_urls       json,
    rooms_label       varchar,
    address_label     varchar,
    address_elements  json,
    heating_gas_label varchar,
    is_heating_gas    boolean,
    is_furnished      boolean,
    contact_label     varchar,
    contact           bigint,
    price_lable       varchar,
    price_label       varchar,
    country_label     varchar,
    region_label      varchar,
    city_label        varchar,
    district_label    varchar,
    lat               double precision,
    long              double precision,
    external_id       varchar
        constraint rent_turkey_pk
            unique,
    tg_chat_id        bigint,
    added_at          timestamp default now(),
    tg_user_id        bigint
);

alter table rent_turkey
    owner to postgres;

create table if not exists rent_turkey_outbox
(
    id                 bigint not null
        constraint rent_turkey_outbox_pk
            primary key,
    is_sent            boolean   default false,
    sent_at            timestamp,
    created_at         timestamp default now(),
    tg_message_id      bigint,
    tg_message_desc_id bigint
);

alter table rent_turkey_outbox
    owner to postgres;

