import './App.css';
import React, { useState } from 'react';

function App() {
   const url = "http://localhost:3001/data"
   const [data, setData] = useState('');

   const fetchApi = async () => {
     try {
       const response =  await fetch(url); 
       const textData =  await response.text();
       setData(textData);
     } catch (error) {
       console.log(error);
     }
   };

   return (
    <div className="App">
    <header className="App-header">
      <p>
      Tarea 1  - SO1 - 1s2024
      </p>
      <br></br>
      <button variant="contained" color="black"  onClick={fetchApi} >
          Mostrar Datos
      </button>
      <br></br>
      <textarea name="textdata" value={data} />
    </header>

  </div>
   );
 }

 export default App;