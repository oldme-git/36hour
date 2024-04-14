create database "36hour_seat"

create table layout
(
    id          serial
        primary key,
    location_id integer      not null,
    policy_c_id integer,
    policy_l_id integer,
    layout_name varchar(100) not null,
    map         text,
    status      smallint,
    sort        integer,
    seats       integer,
    created_at timestamp(6) with time zone not null,
    updated_at timestamp(6) with time zone not null
);

comment on column layout.map is '布局信息';

comment on column layout.policy_c_id is '公共策略id，优先使用';

comment on column layout.policy_l_id is '属于自己的策略id，最后使用';

comment on column layout.status is '是否正常启用，1启用 2不启用';

comment on column layout.sort is '排序，越小越靠前';

comment on column layout.seats is '座位总数';

alter table layout
    owner to postgres;

create table policy_prepare
(
    id   serial
        primary key,
    name varchar(30) not null,
    info text        not null
);

comment on table policy_prepare is '策略预设';

comment on column policy_prepare.info is '策略信息';

alter table policy_prepare
    owner to postgres;

create unique index policy_prepare_name_unique
    on policy_prepare (name);

create table policy_common
(
    id   serial
        primary key,
    name varchar(30) not null,
    info text        not null
);

comment on table policy_common is '公共策略';

comment on column policy_common.info is '策略信息';

alter table policy_common
    owner to postgres;

create unique index policy_common_name_ununique
    on policy_common (name);

create table policy_layout
(
    id   serial
        primary key,
    info text not null
);

comment on table policy_layout is '策略预设';

comment on column policy_layout.info is '策略信息';

alter table policy_layout
    owner to postgres;

