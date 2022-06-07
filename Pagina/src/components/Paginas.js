import React from "react"
import {Routes, Route} from "react-router-dom";
import {ProductosLista} from "./Productos/index";
import Login from "./login/Login.tsx";
import {CategoryLista} from "./Inicio/index"
import './Inicio/inicio.css';
export const Paginas = ()=>{
    return(
        <section>
          <Routes>
          <Route path="/" exact element={<CategoryLista/>} />
        <Route path="/productos" exact element={<ProductosLista/>} />
        <Route path="/login" exact element={<Login/>} />
        </Routes>  
        </section>
    )
}
