import React from 'react';
//import './App.css';
import {Header} from './components/Header';
import 'boxicons';
import {BrowserRouter as Router} from "react-router-dom";
import {Paginas} from './components/Paginas'
import{Dataprovider} from "./context/Dataprovider";

function App() {
  return (
    <Dataprovider>
    <div className="App">
      <Router>
        <Header/>
        <Paginas/>
      </Router>
    </div>
    </Dataprovider>

  );
}

export default App;
