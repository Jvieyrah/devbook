CREATE DATABASE IF NOT EXISTS devbook;

use devbook;

Drop table if exists usuarios;

CREATE TABLE usuarios(
    id int(11) not null primary key auto_increment,
    nome varchar(50) not null,
    nick varchar(50) not null,
    email varchar(50) not null,
    senha varchar(20) not null,
    criadoEm timestamp default current_timestamp()
) Engine=InnoDB;