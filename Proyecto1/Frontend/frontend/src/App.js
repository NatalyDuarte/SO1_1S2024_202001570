import React, { useState, useEffect } from 'react';
import { Pie } from 'react-chartjs-2';
import { Line } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, registerables } from 'chart.js';

ChartJS.register(...registerables);

const API_URL_RAM = 'http://192.168.0.17:8080/tiempor/ram';
const API_URL_CPU = 'http://192.168.0.17:8080/tiempor/cpu';

function App() {
  const [data, setData] = useState({
    freeRam: null,
    cpuInfo: null,
  });
  const [chartDatahis, setChartDatahis] = useState({});  
  const [chartDatahiscp, setChartDatahiscp] = useState({});  

  useEffect(() => {
    const fetchData = async () => {
      try {
        const responseRam = await fetch(API_URL_RAM);
        if (responseRam.ok) {
          const dataRam = await responseRam.json();
          const responseCpu = await fetch(API_URL_CPU);
          if (responseCpu.ok) {
            const dataCpu = await responseCpu.json();
            setData({
              freeRam: dataRam.freeRam,
              cpuInfo: dataCpu,
            });
          } else {
            throw new Error('Error HTTP CPU: ' + responseCpu.status);
          }
        } else {
          throw new Error('Error HTTP RAM: ' + responseRam.status);
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };    
    const interval = setInterval(fetchData, 1000);
    return () => clearInterval(interval);
  }, []);

  useEffect(() =>{
    const fetchChartData = async () => {
      try {
        const response = await fetch('http://192.168.0.17:8080/tiempohis/ram');
        const data = await response.json();
  
        const labels = data.histrams.map(item => item.fech);
        const dataValues = data.histrams.map(item => item.histram);
  
        setChartDatahis({
          labels,
          datasets: [
            {
              label: 'RAM Usada',
              data: dataValues,
              borderColor: 'orange',
              backgroundColor: 'transparent',
              pointBorderColor: 'orange',
              pointBackgroundColor: 'rgba(255,150,0,0.5)',
              pointRadius: 5,
              pointHoverRadius: 10,
              pointHitRadius: 30,
              pointBorderWidth: 2,
              pointStyle: 'rectRounded',
            },
          ],
        });
      } catch (error) {
        console.error('Error fetching chart data:', error);
        // Handle errors appropriately, e.g., display an error message
      }
    };
    const interval = setInterval(fetchChartData, 10000);
    return () => clearInterval(interval);

  }, []);
//===============================================================
  useEffect(() =>{
    const fetchChartDatacp = async () => {
      try {
        const response = await fetch('http://192.168.0.17:8080/tiempohis/cpu');
        const data = await response.json();
  
        const labels = data.histcpus.map(item => item.fech);
        const dataValues = data.histcpus.map(item => item.histcpu);
  
        setChartDatahiscp({
          labels,
          datasets: [
            {
              label: 'CPU Usado',
              data: dataValues,
              borderColor: 'orange',
              backgroundColor: 'transparent',
              pointBorderColor: 'orange',
              pointBackgroundColor: 'rgba(255,150,0,0.5)',
              pointRadius: 5,
              pointHoverRadius: 10,
              pointHitRadius: 30,
              pointBorderWidth: 2,
              pointStyle: 'rectRounded',
            },
          ],
        });
      } catch (error) {
        console.error('Error fetching chart datacp:', error);
        // Handle errors appropriately, e.g., display an error message
      }
    };
    const interval = setInterval(fetchChartDatacp, 10000);
    return () => clearInterval(interval);

  }, []);

  const { freeRam, cpuInfo } = data;
  const totalRam = 16000000;
  const useRam = totalRam - freeRam;
  const porcentajeUsado = ((totalRam - freeRam) / totalRam) * 100;
  const cpuUso = cpuInfo?.cpuTotal - cpuInfo?.cpuPorcentaje;

  const cpuLibre = cpuInfo ? 100 - (cpuInfo.cpu_porcentaje / cpuInfo.cpu_total * 100) : 0;
  const cpuEnUso = cpuInfo ? (cpuInfo.cpu_porcentaje / cpuInfo.cpu_total * 100) : 0;
  const chartData = {
    labels: ['Memoria libre', 'Memoria en uso'],
    datasets: [
      {
        label: 'Memoria RAM',
        data: [freeRam, useRam],
        backgroundColor: ['#2ECC71', '#E74C3C'],
        borderColor: ['#2ECC71', '#E74C3C'],
        borderWidth: 1
      }
    ]
  };

  const cpuChartData = {
    labels: ['CPU libre', 'CPU en uso'],
    datasets: [
      {
        label: 'CPU',
        data: [cpuLibre, cpuEnUso],
        backgroundColor: ['#2ECC71', '#E74C3C'],
        borderColor: ['#2ECC71', '#E74C3C'],
        borderWidth: 1
      }
    ]
  };
  
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
            <h1><center>Monitoreo en tiempo real</center></h1>
              <div class="col-lg-6">
                <div class="p-2">
                  <h3>Memoria RAM</h3>
                  <Pie data={chartData} />
                </div>
              </div>
              <div class="col-lg-6">
                <div class="p-2">
                  <h3>CPU</h3>
                  <Pie data={cpuChartData} />
                </div>
              </div>
            </div>
          </div>
      </section>
      <section id="historico">
          <div class="container px-5">
            <div class="row gx-5 align-items-center">
            <h1><center>Monitoreo Historico</center></h1>
              <div class="col-lg-6">
                <div class="p-2">
                  <h3>Memoria RAM Historica</h3>
                  {chartDatahis.labels && chartDatahis.datasets && (
                      <Line data={chartDatahis} />
                    )}
                </div>
              </div>
              <div class="col-lg-6">
                <div class="p-2">
                  <h3>CPU Historico</h3>
                  {chartDatahiscp.labels && chartDatahiscp.datasets && (
                      <Line data={chartDatahiscp} />
                    )}
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