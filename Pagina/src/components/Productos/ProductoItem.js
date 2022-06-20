import React from "react";
import {addToCart} from "../cookies/add2car"
export const ProductoItem =(
    {id,
    name,
    picture_url,
    base_price,
    stock,
    description,
    id_category
})=>{


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
            {stock>0?<p>Stock: {stock}</p>:
            <p>Sin Stock</p>}
            
            <p className="price">U$S {base_price}</p>
        </div>
        <div className="buttom">
        {stock>0?
            <button className="btn" onClick={()=> addToCart(id)}>
                Agregar
            </button>:
            <p>Este Producto no esta Disponible</p>}
        </div>
        </div>
    )
}