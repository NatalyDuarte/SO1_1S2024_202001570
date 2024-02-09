const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const { MongoClient } = require('mongodb');
const fs = require('fs');
const path = require('path');

const app = express();
const port = 3002;

const mongoURI = 'mongodb://localhost:27017/BaseTa'; // Reemplaza con tu URI de MongoDB

app.use(cors());
app.use(bodyParser.json({ limit: '10mb' }));
app.use(bodyParser.urlencoded({ extended: true, limit: '10mb' }));

app.post('/enviar', async (req, res) => {
  let client; // Declarar la variable client fuera del bloque try

  try {
    const { photo, date } = req.body;

    // Decodifica la imagen base64
    const imageData = photo.replace(/^data:image\/\w+;base64,/, '');
    const buffer = Buffer.from(imageData, 'base64');

    // Guarda la imagen en el servidor
    const fileName = `${date}_photo.png`;
    const filePath = path.join(__dirname, 'uploads', fileName);
    fs.writeFileSync(filePath, buffer);

    // Conéctate a la base de datos y almacena los datos
    client = new MongoClient(mongoURI, { useNewUrlParser: true, useUnifiedTopology: true });
    await client.connect();

    const database = client.db();
    const collection = database.collection('fotos');

    const result = await collection.insertOne({
      base64: imageData,
      fecha_toma: new Date(date),
      filePath: filePath // Opcional: puedes almacenar la ruta del archivo si lo deseas
    });

    console.log('Foto almacenada con éxito:', fileName);
    console.log('ID del documento en MongoDB:', result.insertedId);

    res.status(200).json({ success: true });
  } catch (error) {
    console.error('Error al procesar la foto:', error);
    res.status(500).json({ success: false, error: 'Internal Server Error' });
  } finally {
    if (client) {
      await client.close(); // Cierra la conexión a la base de datos al finalizar
    }
  }
});

app.listen(port, () => {
  console.log(`Servidor escuchando en el puerto ${port}`);
});
