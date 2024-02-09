import React, { useRef } from 'react';
import './App.css';

function App() {
  const videoRef = useRef(null);
  const photoRef = useRef(null);

  const startCamera = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ video: true });

      if (videoRef.current) {
        videoRef.current.srcObject = stream;
      }
    } catch (error) {
      console.error('Error al acceder a la cámara:', error);
    }
  };

  const takePhoto = () => {
    if (videoRef.current) {
      const canvas = document.createElement('canvas');
      const video = videoRef.current;
  
      canvas.width = video.videoWidth / 2;
      canvas.height = video.videoHeight / 2;
  
      const ctx = canvas.getContext('2d');
      ctx.drawImage(video, 0, 0, canvas.width, canvas.height);
  
      // Convierte la imagen a base64
      const dataURL = canvas.toDataURL('image/png');
  
      if (photoRef.current) {
        photoRef.current.style.width = `${canvas.width}px`;
        photoRef.current.style.height = `${canvas.height}px`;
        photoRef.current.src = dataURL;
  
        // Envía la foto a la API
        sendPhotoToAPI(dataURL);
      }
    }
  };
  
  const sendPhotoToAPI = async (photoBase64) => {
    try {
      const currentDate = new Date().toISOString();
  
      const response = await fetch('http://localhost:3002/enviar', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          photo: photoBase64,
          date: currentDate,
        }),
      });
  
      if (response.ok) {
        console.log('Foto enviada con éxito a la API.');
      } else {
        console.error('Error al enviar la foto a la API:', response.statusText);
      }
    } catch (error) {
      console.error('Error en la solicitud POST:', error);
    }
  };
  

  return (
    <div className="App">
      <header className="App-header">
        <center><h1>Tarea 2</h1></center>
        <center><h2>Nataly Guzman - 202001570</h2></center>
        <button onClick={startCamera}>Habilitar Cámara</button><br></br>
        <video ref={videoRef} autoPlay muted playsInline className="small-video" />
        <button onClick={takePhoto}>Tomar Foto</button><br></br>
        <img ref={photoRef} alt="Foto" />
      </header>
    </div>
  );
}

export default App;
