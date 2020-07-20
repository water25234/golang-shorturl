# golang-shorturl
- 縮短網址專案

## Prerequisite
- GO v1.14
- Go Gin v1.5.0
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
);

CREATE INDEX index_shortener_url_encode
ON shortener(url_encode);

ALTER TABLE shortener
ADD CONSTRAINT shortener_url_encode_unique UNIQUE (url_encode);

```

## Build the project
- cp env-example .env
- docker-compose up -d


## Browers Work
- http://localhost:8080/shortener

![image](https://github.com/water25234/golang-shorturl/blob/master/asset/img/Screen%20Shot%202020-07-18%20at%203.22.27%20PM.png)


## RESTful API
- Create Shortener Url API
```
Request

curl --location --request POST 'http://127.0.0.1:8080/api/v1/shortener' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "url": "https://www.kkday.com/zh-tw"
    }'

Response

{
   "data":{
      "url_encode":"lVqELZ",
      "url":"https://www.cakeresume.com/dashboard",
      "full_url":"http://127.0.0.1:8080/api/v1/shortener/lVqELZ"
   },
   "metadata":{
      "desc":"success",
      "status":"0000"
   }
}
```
- Get Shortener Url API
```
Request

curl --location --request GET 'http://127.0.0.1:8080/api/v1/shortener/IVqErQ'

Response

{
    "data": {
        "shortenerURL": "https://www.kkday.com/zh-tw"
    },
    "metadata": {
        "desc": "success",
        "status": "0000"
    }
}
```
