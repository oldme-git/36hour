drop table if exists seat;
create table layout
(
	id serial,
	hall_id integer not null,
	layout_name varchar(100) not null
);

alter table layout owner to postgres;
