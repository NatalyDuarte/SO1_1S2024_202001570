Ingenieria en Ciencias y Sistemas  
Sistemas Operativos 1  
Catedratico: Jesús Alberto Guzmán Polanco  
Aux: Daniel Velasquez
**Nombre:** Nataly Saraí Guzmán Duarte  
**Carnet:** 202001570  

# Documentacion Proyecto 2

## Introduccion
Se desarrollará un sistema de votaciones distribuido para un concurso de bandas de música. El sistema utilizará microservicios en Kubernetes, sistemas de mensajería para encolar tareas, Grafana para dashboards y bases de datos para almacenar datos y logs.


### Objetivos
- Implementar un sistema distribuido con microservicios en Kubernetes.
- Utilizar Grafana como interfaz gráfica de dashboards.
- Analizar el rendimiento de los servicios gRPC y Rust.
- Determinar en qué casos utilizar cada tecnología.

### Descripcion de cada tecnologia utilizada 
- Locust: Generador de tráfico que envía datos a los servicios en Kubernetes.
- gRPC: Framework para la comunicación entre servicios basado en RPC.
- Wasm: Tecnología que permite ejecutar código WebAssembly en entornos no web.
- Kafka: Sistema de mensajería distribuida para encolar tareas.
- Kubernetes: Plataforma de orquestación de contenedores para implementar y administrar microservicios.
- Redis: Base de datos NoSQL en memoria para almacenar contadores en tiempo real.
- MongoDB: Base de datos NoSQL para almacenar logs.
- Grafana: Plataforma para crear dashboards e visualizar métricas.
- Cloud Run: Plataforma para implementar y ejecutar aplicaciones sin servidor.

### Descripcion de cada deployment y service de k8s
- Productores:
    - gRPC: Servidor y cliente escritos en Golang que envían datos a Kafka.
    - Wasm: Servidor escrito en Rust que envía datos a Kafka utilizando WasmEdge.
- Servidor de Kafka: Recibe mensajes de los productores y los encola para que sean consumidos por el consumidor.
- Consumidor: Daemon escrito en Golang que consume mensajes de Kafka, procesa los datos y los almacena en Redis y MongoDB.
- Grafana: Se conecta a Redis para mostrar dashboards con las votaciones en tiempo real.
- Cloud Run:
    - API NodeJS: Permite consultar logs de MongoDB.
    - Webapp Vue.js: Muestra los últimos 20 logs de MongoDB.

### Conclusiones

