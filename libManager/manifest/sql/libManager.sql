create database "36hour_lib";

create table floor
(
	id serial
		primary key,
	lib_id integer not null,
	floor_name varchar(100) not null,
	created_at timestamp(6) with time zone not null,
	updated_at timestamp(6) with time zone not null
);

alter table floor owner to postgres;

create table location
(
	id serial
		primary key,
	lib_id integer not null,
	floor_id integer not null,
	location_name varchar(100) not null,
	created_at timestamp(6) with time zone not null,
	updated_at timestamp(6) with time zone not null
);

alter table location owner to postgres;

create table lib
(
	id serial
		primary key,
	lib_name varchar(100) not null,
	address varchar(255) not null,
	active boolean not null,
	created_at timestamp(6) with time zone not null,
	updated_at timestamp(6) with time zone not null
);

comment on column lib.address is '地址';

comment on column lib.active is '是否正在使用';

alter table lib owner to postgres;
