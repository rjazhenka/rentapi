create table public.rent_turkey (
  id integer not null default nextval('rent_turkey_id_seq'::regclass),
  title character varying,
  rooms integer,
  price numeric,
  country integer,
  city integer,
  region integer,
  district integer,
  description text,
  link character varying,
  source integer,
  status integer,
  images_tg_ids json,
  images_urls json,
  rooms_label character varying,
  address_label character varying,
  address_elements json,
  heating_gas_label character varying,
  is_heating_gas boolean,
  is_furnished boolean,
  contact_label character varying,
  contact bigint,
  price_lable character varying,
  price_label character varying,
  country_label character varying,
  region_label character varying,
  city_label character varying,
  district_label character varying,
  lat double precision,
  long double precision,
  external_id character varying,
  tg_chat_id bigint
);

