import React, { useContext, useEffect, useState} from "react";
import { ProductoItem } from "./ProductoItem";
import "bootstrap/dist/css/bootstrap.min.css";
import { ProductosBuscador } from "./buscador";

export const ProductosLista = ()=>{

    const [productos,setProductos] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/productRandom/9')
    .then((response) => response.json());
    setProductos(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])
    return(
        <>
        <ProductosBuscador/>
        
        <h2>Nuestro Productos Populares</h2>
        <div className="productos">
            {
                productos.map(producto =>(
                  <ProductoItem key={producto.id}
                  id={producto.id}
                  name={producto.name}
                  base_price={producto.base_price}
                  id_category={producto.id_category}
                  stock={producto.stock}
                  picture_url={producto.picture_url}
                  description={producto.description}
                  /> 
                ))
            }
            
        </div>
        </>
    )
}