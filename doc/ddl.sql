drop table if exists tbl_login_users cascade;
drop table if exists mst_roles cascade;
drop table if exists tbl_chat_hist cascade;

create table tbl_login_users (
  del_flg boolean default false not null
  , exec_date timestamp(6) without time zone default now() not null
  , exec_user_id character varying(50) not null
  , user_id character varying(50) not null
  , password character varying(50) not null
  , user_name character varying(100) not null
  , role_id character varying(2) not null
  , icon character varying(20)
  , primary key (user_id)
);

create table mst_roles (
  del_flg boolean default false not null
  , exec_date timestamp(6) without time zone default now() not null
  , exec_user_id character varying(50) not null
  , role_id character varying(50) not null
  , role_name character varying(50) not null
  , primary key (role_id)
);

create table tbl_chat_hist (
  del_flg boolean default false not null
  , exec_date timestamp(6) without time zone default now() not null
  , exec_user_id character varying(50) not null
  , user_id_from character varying(50) not null
  , user_id_to character varying(50) not null
  , send_date timestamp(6) without time zone default now() not null
  , msg_str character varying(1000)
  , msg_stmp character varying(20)
  , chk_read boolean default false not null
  , primary key (user_id_from, send_date)
);