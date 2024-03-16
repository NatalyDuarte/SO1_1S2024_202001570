import './App.css';
import React, { useState, useEffect } from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'react-chartjs-2';


const API_URL = 'http://localhost:8080/ram_info';

function App() {
  const [freeRam, setFreeRam] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(API_URL);
        if (response.ok) {
          const data = await response.json();
          setFreeRam(data.freeRam);
        } else {
          throw new Error('Error HTTP: ' + response.status);
        }
      } catch (error) {
        console.error('Error fetching data:', error);
        // Handle errors here (e.g., display an error message to the user)
      }
    };

    fetchData();
  }, []);

  return (
    <div className="App">
      <header class="masthead text-center text-white">
        <div class="masthead-content">
          <div class="container px-5">
              <h1 class="masthead-heading mb-0">Proyecto1 -1S2024</h1>
              <h2>Estudiante: Nataly Saraí Guzmán Duarte</h2>
          </div>
        </div>
        <div class="bg-circle-1 bg-circle"></div>
        <div class="bg-circle-2 bg-circle"></div>
        <div class="bg-circle-3 bg-circle"></div>
      </header>
      <nav class="navbar navbar-expand-lg navbar-dark navbar-custom fixed-top">
        <div class="container px-5">
            <a class="navbar-brand" href="#page-top">Proyecto1</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>
            <div class="collapse navbar-collapse" id="navbarResponsive">
                <ul class="navbar-nav ms-auto">
                    <li class="nav-item"><a class="nav-link" href="#tiemporeal">Monitoreo en tiempo real</a></li>
                    <li class="nav-item"><a class="nav-link" href="#historico">Monitoreo historico</a></li>
                    <li class="nav-item"><a class="nav-link" href="#arbol">Arbol de Procesos</a></li>
                    <li class="nav-item"><a class="nav-link" href="#estados">Diagrama de estados</a></li>
                </ul>
            </div>
        </div>
      </nav>
      <section id="tiemporeal">
          <div class="container px-5">
            <div class="row gx-5 align-items-center">
              <div class=" order-lg-1">
                <div class="p-2">
                <h2>Monitoreo en tiempo real</h2>
                <h1> RAM</h1><br></br>
                console.log(freeRam);
                </div>
              </div>
            </div>
          </div>
      </section>
      <section id="historico">
          <div class="container px-5">
            <div class="row gx-5 align-items-center">
              <div class=" order-lg-1">
                <div class="p-2">
                <h2>Monitoreo Historico</h2>
                </div>
              </div>
            </div>
          </div>
      </section>
      <section id="arbol">
          <div class="container px-5">
            <div class="row gx-5 align-items-center">
              <div class=" order-lg-1">
                <div class="p-2">
                <h2>Arbol de procesos</h2>
                </div>
              </div>
            </div>
          </div>
      </section>
      <section id="estados">
          <div class="container px-5">
            <div class="row gx-5 align-items-center">
              <div class=" order-lg-1">
                <div class="p-2">
                <h2>Diagrama de estados</h2>
                </div>
              </div>
            </div>
          </div>
      </section>
    </div>
  );
}

export default App;