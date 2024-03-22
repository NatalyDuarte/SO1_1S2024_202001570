Ingenieria en Ciencias y Sistemas  
Sistemas Operativos 1  
Catedratico: Jesús Alberto Guzmán Polanco  
**Nombre:** Nataly Saraí Guzmán Duarte  
**Carnet:** 202001570  

# Proyecto1

## Base de datos
Se utilizo mysql para realizar la base de datos, en esta se almaceno la informacion de la CPU y de la memoria RAM. Esta base de datos se utilizo para poder mostrar las graficas con los datos historicos donde se muestra el uso de CPU y memoria RAM. A continuacion se muestran las tablas creadas con sus atributos:

* raminfo
    * freeram
    * boundran
    * time
* cpuinfo
    * freecpu
    * boundcpu
    * time


## Backend
Para el backend se utilizo el lenguaje de programacion GO que cuenta con funciones para el envio de datos a la base de datos, funciones para extraer el ultimo dato de las tablas de la base de datos y funciones para extraer todos los datos almacenados en las bases de datos. Para poder realizar la extraccion de los datos de los modulos se utilizaron go rutinas que se ejecutan cada 5 segundos


Los endpoints que se crearon fueron los siguientes
* /tiempo/ram
* /tiempo/cpu
* /tiempohis/ram
* /tiempohis/cpu

*fotos backend


## Fronted
Para el fronted se utilizo React junto a la libreria de chartjs para poder mostrar la informacion del CPU y de la memoria RAM. Para el uso y libre se utilizo un grafico circular y para los datos historicos se utilizo una grafica lineal. Para poder mostrar los datos en la grafica cada cierto tiempo se utilizo useEffect que se ejecuta cada 5 segundos.

*foto

## Dockerfile backend

Este archivo se utiliza para crear la imagen del backend realizado en GO. Aqui se copia la imagen base , los archivos go.mod y go.sum, luego descargamos las dependencias, despues compilamos y por ultimo exponemos el puerto a utilizar.


* foto de baackend dockerfile


## Docker-compose
El archivo define tres servicios: una base de datos MySQL, una API backend y una interfaz de usuario frontend. La base de datos se ejecuta en un contenedor con su propio volumen persistente para los datos. La API se construye a partir del Dockerfile actual y se conecta a la base de datos. La interfaz de usuario se construye a partir de un directorio separado y se conecta a la API. Los tres contenedores se ejecutan en una red interna y se exponen los puertos relevantes al host. El archivo también define variables de entorno para la configuración de la base de datos, la API y la interfaz de usuario.
 *foto


## App
Se utilizan componentes para los diferentes tipodes de graficas, como lo es: GraficaCpu, GraficaRam, HistoricoCpu, HistoricoRam, Home, entre otros. Como podemos ver obtenemos la informacion que nos devuelve los endpoint, y con esta informacion ya podemos graficar utilizando tanto la grafica de pie para los tiempo real, line para el tiempo real.


*foto
  


## Nginx
El comando docker run ejecuta una imagen de Docker en un contenedor. En este caso, se ejecuta una imagen de Nginx con la configuración del archivo nginx.conf. El comando expone el puerto 80 del contenedor al puerto 80 del host, monta los volúmenes /www/html y /etc/nginx/conf.d del host en el contenedor, y ejecuta el comando cat /etc/nginx/conf.d/default.conf para mostrar la configuración de Nginx.

*foto

## Modulos
Modulo Ram:
El módulo del kernel implementa un archivo en /proc que muestra la cantidad de memoria RAM libre del sistema. El módulo utiliza la estructura struct sysinfo para obtener información del sistema, incluyendo la memoria RAM libre. La función escribir_archivo se encarga de mostrar la cantidad de memoria RAM libre en el archivo /proc/ram_so1_1s2024. La función al_abrir se encarga de manejar la apertura del archivo. El módulo se carga y se elimina del kernel utilizando las funciones _inserty_remove respectivamente. Las funciones module_initymodule_exit se utilizan para indicar al kernel las funciones de carga y eliminación del módulo

Modulo Cpu:
Este módulo del kernel crea un archivo en /proc llamado "cpu_so1_1s2024" que proporciona información sobre el uso de la CPU y los procesos del sistema. El módulo itera a través de todos los procesos y recopila información como el tiempo de CPU, el uso de RAM, el estado del proceso (ejecutando, durmiendo, zombie, detenido) y el nombre del proceso. El módulo también calcula el uso total de CPU y la cantidad de procesos en cada estado.

La información se formatea en un JSON y se escribe en el archivo /proc cuando se lee.

## Comandos

Se utilizaron los siguientes comandos para la realizacion de los modulos

* sudo insmod modulo_cpu.ko
* sudo insmod modulo_ram.ko
* make clean
* make all

Se utilizaron los siguientes comandos para la parte de docker

* docker build -t nombre . 
* docker compose up -d
* docker compose stop
* docker images

Se utilizaron los siguientes comandos para la parte de docker hub

* sudo docker push user/imagen
* sudo docker pull user/imagen

