import React, { useContext, useEffect, useState } from "react";
import { CategoryItem } from "./CategoryItem";
import "bootstrap/dist/css/bootstrap.min.css";

export const CategoryLista =()=>{

    const [categorias,setCategorias] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/category')
    .then((response) => response.json());
    setCategorias(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])
    
    return(
       
        <>
        <h1 className="title"> Categorias</h1>
        <div className="categorias">
            {
                categorias.map(categoria =>(
                  <CategoryItem key={categoria.id_category}
                  name={categoria.nombre}
                  description={categoria.descripcion}
                  />  
                ))
            }
        </div>
        </>
    )
}


