
const Redis = require('ioredis');
const conexion = new Redis ({
  host: '10.12.128.3',
  port: 6379,
  connectTimeout: 5000,
});

function funcionPub() {
  const mensaje = '{msg: "Hola esta es la Tarea3"}'
  conexion.publish('test', mensaje)
    .then(() => {
      console.log("Mensaje publicado con exito");
    })
    .catch((err) => {
      console.error("Ocurrio un error al publicar el mensaje: ", err);
    });
}

setInterval(funcionPub, 3000);
