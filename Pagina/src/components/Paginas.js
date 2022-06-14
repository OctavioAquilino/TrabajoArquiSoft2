import React from "react"
import {Routes, Route} from "react-router-dom";
import {ProductosLista} from "./Productos/index";
import IndexLogin from "./login/IndexLogin";
import {CategoryLista} from "./Inicio/index"
import Cart from "./Carrito";
import {GetOrders} from "./orden/get_orden"

export const Paginas = ()=>{
    return(
        <section>
          <Routes>
          <Route path="/" exact element={<CategoryLista/>} />
        <Route path="/productos" exact element={<ProductosLista/>} />
        <Route path="/login" exact element={<IndexLogin/>} />
        <Route path="/cart"  element={<Cart/>} />
        <Route path="/order" exact element={<GetOrders/>} />
        </Routes>  
        </section>
    )
}
