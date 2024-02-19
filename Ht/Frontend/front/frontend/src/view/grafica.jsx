import React, { useEffect, useState } from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'react-chartjs-2';
import './styles.css';

ChartJS.register(ArcElement, Tooltip, Legend);
import { Greet } from '../../wailsjs/go/main/App';

export function Graficar() {
  const [totalRam, setTotalRam] = useState(8000000); 
  const [freeRam, setFreeRam] = useState(0); 
  const [useRam, setUseRam] = useState(0); 
  const [porcentajeUsado, setPorcentajeUsado] = useState(0);

  const updateFreeRam = (result) => {
    setFreeRam(result);
    setUseRam(totalRam - result);
    const actualizado = ((totalRam - result) / totalRam) * 100;
    setPorcentajeUsado(actualizado.toFixed(2)); 
  };

  useEffect(() => {
    const interval = setInterval(() => {
      Greet().then(updateFreeRam);
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="chart-container">
      <h2> Nataly Guzman</h2><br></br>
      <h1> 202001570</h1><br></br>
      <h1> RAM</h1><br></br>
      <Pie data={{
        labels: ['Libre', 'En uso'],
        datasets: [
          {
            label: '# of Votes',
            data: [freeRam, useRam],
            backgroundColor: [
              'rgba(255, 165, 0, 0.2)',
              'rgba(0, 255, 0, 0.2)', 
            ],
            borderColor: [
              'rgba(255, 165, 0, 0.2)', 
              'rgba(0, 255, 0, 0.2)', 
            ],
            borderWidth: 1,
          },
        ],
      }} />
    </div>
  );
}
