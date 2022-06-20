import React from 'react';
import {Header} from './components/Header';
import 'boxicons';
import {BrowserRouter as Router} from "react-router-dom";
import {Paginas} from './components/Paginas';
function App() {
  return (
    
    <div className="App">
      <Router>
        <Header/>
        <Paginas/>
      </Router>
    </div>
 
  );
}

export default App;