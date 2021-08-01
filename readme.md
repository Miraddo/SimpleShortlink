# Simple Short Link

a simple app to make links/urls shorter and return the short value to get the main url


## Install and Configure DataBase with Docker
I used postgresql in this project.

### steps
- Install Docker ^_^
- run this command `docker pull postgres:latest` to download the last version of postgres
- after that to have storage create volume `docker volume create pg-volume`
- then run the container with this command `docker run -d --name=pgdb -p 5432:5432 -v pg-volume:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123456 postgres`
  - with `--name` set the name for your container here I set `pgdb` you can change it 
  - with `POSTGRES_PASSWORD` you can set the password for your database is important I set `123456` ^_^ you should use a strong password not like me  
- if you don't get any error so probably every thing correctly sets now you can use `docker container ls` or `docker ps` to see the container is run
- now you ready to add the sql below here to your database
- to add the sql below here first use this command  `docker exec -it pgdb psql -U postgres` then paste the sql below here
```postgresql
create database shortlink;
create table urls
(
    id  serial
    constraint urls_pk
    primary key,
    url text,
    key char(8)
);

alter table urls
owner to postgres;

create unique index urls_key_uindex
on urls (key);
```

after all you should open/edit `config/config.go` file and set your database info correctly

## How to use this app
for use this app we had three command after build the code or run the code as like :
```bash
# first 
# run an simple http server you can access to it with localhost:8080
go run cmd/main.go serv 
# in this part you can send request to localhost:8080/short?url=<your url>
# to get the main url with key send request to localhost:8080/url?key=<your key>

# second
# if you don't want use the http server just use the command line interface 
go run cmd/main.go short <your url> # to return the key
# Key : #######

go run cmd/main.go url <your key> # to return the main url based on key
# Url : #######
```

## Libraries
Libraries that I used in this app

- [Google UUID](https://github.com/google/uuid) 
- [PostgresSQL Library](https://github.com/lib/pq) 
- [Cobra CommandLine](https://github.com/spf13/cobra)
- [ZAP Logger](https://pkg.go.dev/go.uber.org/zap)
