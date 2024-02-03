import './App.css';
//import React, { useState } from 'react';

function App() {
var fechaHoraActual = new Date();
var fechaActual = fechaHoraActual.toLocaleDateString();

var horaActual = fechaHoraActual.toLocaleTimeString();

console.log("Fecha actual: " + fechaActual);
console.log("Hora actual: " + horaActual);
const fetchApi = async () => {
  const url = "http://localhost:3000/data"
  const [data, setData] = useState('');
  
  const fetchApi = async () => {
  try {
  const response = await fetch(url); 
  const textData = await response.text();
  console.log("Aqui"+response)
  setData(textData);
  } catch (error) {
  console.log(error);
  }
  };
  textData = textData + " Fecha: " + fechaActual + " Hora: " + horaActual +"</p>"
  if (elemento) {
    elemento.innerHTML = textData;
  }


return (
<div className="App">
<header className="App-header">
<p>
Tarea 1 - SO1 - 1s2024
</p>
<br></br>
<button variant="contained" color="black" onClick={fetchApi} >
Mostrar Datos
</button>
<br></br>
<div id="textdata"></div>
</header>

</div>
);
}

export default App;
