import React, { useContext, useEffect, useState} from "react";
import { ProductoItem } from "./ProductoItem";
import "bootstrap/dist/css/bootstrap.min.css";
import { ProductosBuscador } from "./buscador";

export const ProductosLista = ()=>{

    const [productos,setProductos] = useState([]);
    //const [prodseach,setProdseach] = useState([]);
    const [busqueda, setBusqueda]= useState("");
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/productRandom/9')
    .then((response) => response.json());
    setProductos(response);
    //setProdseach(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])
    /*const handleChange=e=>{
        setBusqueda(e.target.value);
        filtrar(e.target.value);
      }
      const filtrar=(terminoBusqueda)=>{
        var resultadosBusqueda=prodseach.filter((elemento)=>{
          if(elemento.name.toString().toLowerCase().includes(terminoBusqueda.toLowerCase()))
          {
            return elemento;
          }
        });
        setProductos(resultadosBusqueda);
      }
*/
    return(
        <>
        <ProductosBuscador/>
        
        

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