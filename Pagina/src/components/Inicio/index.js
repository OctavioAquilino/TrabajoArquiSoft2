import React, {useEffect, useState } from "react";
import { CategoryItem } from "./CategoryItem";
import "bootstrap/dist/css/bootstrap.min.css";
import { ProductoItem } from "../Productos/ProductoItem";
import { Cookies } from "react-cookie";

async function GetProductByIdCategory(id) {
    return fetch('http://localhost:8090/productCategory/'+id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(data => data.json())
      /*.then(data => {
          console.log(data)
      })*/
   }


function RenderStart(){
    "No hay Productos en la Categoria"    
}
/*function RenderProducts(){

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
    alert("holaaaa")
}*/
/* return if(productos){
    RenderProducts

} else{
    RenderStart
}
*/
export const CategoryLista =()=>{
    //const Cookie = new Cookies();    
    
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
    /*useEffect(() => {
        Handle();
    }, [])*/

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
        <h1 className="title"> CATEGORIAS</h1>
        <div className="categorias">
            {
                categorias.map(categoria =>(
                  <button className="category"onClick={()=>Handle(categoria.id_category)}>
                      <CategoryItem key={categoria.id_category}
                  name={categoria.nombre}
                  description={categoria.descripcion}
                  />
                  </button>
                ))
            }
        </div>
        {productos? RenderStart():Render}  
        </>
    );
}


