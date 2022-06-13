import React from "react";
import { useContext } from "react";
import {Link} from "react-router-dom";
import { DataContext } from "../Carrito/Provider";
export const ProductoItem =(
    {id,
    name,
    picture_url,
    base_price,
    stock,
    description,
    id_category
})=>{

    const value= useContext(DataContext);
    const addCarrito = value.addCarrito;
    return(
        <div className="producto">
        <a href="#">
        <div className="producto_img">
            <img className="image" src={picture_url} alt=""/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>{name}</h1>
            <p>{description}</p>
            <p>Stock: {stock}</p>
            <p className="price">${base_price}</p>
        </div>
        <div className="buttom">
            <button className="btn" onClick={()=> addCarrito(id)}>
                Agregar
            </button>
        </div>
        </div>
    )
}