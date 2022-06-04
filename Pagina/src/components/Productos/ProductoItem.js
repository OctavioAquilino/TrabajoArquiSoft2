import React from "react";

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
        </div>
        </a>
        <div className="producto_footer">
            <h1>{name}</h1>
            <p>{description}</p>
            <p>Stock: {stock}</p>
            <p className="price">${base_price}</p>
        </div>
        <div className="buttom">
            <button className="btn">
                Agregar
            </button>
            <div>
              <a href="#"className="btn">Vista</a>  
            </div>
        </div>
        </div>
    )
}