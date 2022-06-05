import React, { useContext, useEffect, useState} from "react";
import { ProductoItem } from "./ProductoItem";
import "bootstrap/dist/css/bootstrap.min.css";

export const ProductosLista = ()=>{

    const [productos,setProductos] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/product')
    .then((response) => response.json());
    setProductos(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])

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