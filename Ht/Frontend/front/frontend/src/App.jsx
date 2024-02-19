import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';

import {Greet} from "../wailsjs/go/main/App";

import {Graficar} from './view/grafica';

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);


    function greet() {
        Greet(name).then(updateResultText);
    }

    

    return (
        <div id="App">
            
            
            <Graficar></Graficar>
            
        </div>
    )
}

export default App
