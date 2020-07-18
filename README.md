# golang-shorturl

## Prerequisite
- GO v1.14
- Redis
- postgres

## Create Postgresql Table Schema

```
CREATE SEQUENCE public.shortener_id_sql
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;

CREATE TABLE public.shortener
(
    shortener_id integer NOT NULL DEFAULT nextval('shortener_id_sql'::regclass),
    create_date timestamp,
    modify_date timestamp,
    url_encode character varying(100), 
    url character varying ,
    is_close boolean NOT NULL DEFAULT false, 
    CONSTRAINT shortener_pkey PRIMARY KEY (shortener_id)
)

CREATE INDEX index_shortener_url_encode
ON shortener(url_encode);

ALTER TABLE shortener
ADD CONSTRAINT shortener_url_encode_unique UNIQUE (url_encode);

```

## Build the project
- docker-compose up -d