#!/bin/bash

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv

docker compose up -d

docker exec -it postgres sh

psql -U postgres

create table tasks (
  id integer, 
  task varchar(255),
  is_completed boolean,
  created_at timestamp,
  updated_at timestamp
);