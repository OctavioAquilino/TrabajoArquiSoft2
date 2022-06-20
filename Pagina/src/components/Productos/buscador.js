import React, {useState} from "react";
import { ProductoItem } from "./ProductoItem";
import "bootstrap/dist/css/bootstrap.min.css";

export const ProductosBuscador = ()=>{

    const [productos,setProductos] = useState([]);
    const [busqueda, setBusqueda]= useState("");

    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({texto:busqueda})
    };

    const fetchApi = async()=>{
      
        const response = await fetch('http://localhost:8090/productText',requestOptions)
        .then((response) => response.json());
        if (response.status == 400) {
          alert("NO SE ENCONTRÓ DICHO PRODUCTO")
          window.location.reload();
       }else{
        
        setProductos(response)
        console.log(response);
       }
        };

    const handleChange=e=>{
        setBusqueda(e.target.value);
      };

      const handleSubmit= (event)=>{
        event.preventDefault();
        fetchApi();

    };
    
    return(
        <>
        <h1 className="title"> PRODUCTOS</h1>
        <div className="containerInput" >
        <input
           
          className="form-control inputBuscar"
          value={busqueda}
          placeholder="Buscador de Productos"
          onChange={handleChange}
         
        />
        <input
        value = "Buscar"
         type = "button"
        onClick = {handleSubmit}
        />
      </div>
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