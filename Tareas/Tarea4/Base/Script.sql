Create Database tarea4;
use tarea4;

create table proyecto(
    id_album INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(200),
    album varchar(200),
    year varchar(200),
    ranks varchar(200)
);

select * from proyecto;
DELETE from proyecto; 