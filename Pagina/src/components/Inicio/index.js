import React, {useEffect, useState } from "react";
import { CategoryItem } from "./CategoryItem";
import "bootstrap/dist/css/bootstrap.min.css";
import { ProductoItem } from "../Productos/ProductoItem";

async function GetProductByIdCategory(id) {
    return fetch('http://localhost:8090/productCategory/' +id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(data => data.json())
   }

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

    const [productos,setProductos]=useState([]);
    async function Handle (id) {
    const response = await GetProductByIdCategory(id)
    setProductos(response)
    console.log(response);
    };
const Render =(
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
)
    return(
        <>
        <div className="categorias">
            {
                categorias.map(categoria =>(
                  <button className="category"onClick={()=>Handle(categoria.id_category)}>
                      <CategoryItem key={categoria.id_category}
                  name={categoria.nombre}
                  />
                  </button>
                ))
                
            }
        </div>
          <div>
              {productos>0? "No hay Productos en esta categoria": Render}
          </div>
        </>
    );
}