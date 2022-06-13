import React from 'react';
import {Header} from './components/Header';
import 'boxicons';
import {BrowserRouter as Router} from "react-router-dom";
import {Paginas} from './components/Paginas';
import {Carrito} from './components/Carrito/index';
import {Dataprovider} from './components/Carrito/Provider';
function App() {
  return (
    <Dataprovider>
    <div className="App">
      <Router>
        <Header/>
        <Carrito/>
        <Paginas/>
      </Router>
    </div>
  </Dataprovider>
  );
}

export default App;