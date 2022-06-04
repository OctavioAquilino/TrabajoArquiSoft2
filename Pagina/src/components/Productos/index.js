import React, { useContext } from "react";
import { DataContext } from "../../context/Dataprovider";
import { ProductoItem } from "./ProductoItem";
import "bootstrap/dist/css/bootstrap.min.css";

export const ProductosLista = ()=>{

    const value = useContext(DataContext)
    const [productos] = value.productos
    console.log(productos)

    return(
        <>
        <h1 className="title"> PRODUCTOS</h1>
        
        <div className="productos">
            {
                productos.map(producto =>(
                  <ProductoItem key={producto.id}
                  id={producto.id}
                  name={producto.name}
                  base_price={producto.base_price}
                  stock={producto.stock}
                  description={producto.description}
                  />  
                ))
            }
        </div>
        </>
    )
}