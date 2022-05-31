import React from "react";
import Logo from "../../Imagenes/Logo1.png";

export const ProductoItem =(
    {id,
    title,
    price,
    image,
    category,
    img1,
    img2,
    cantidad
})=>{
    return(
        <div className="producto">
        <a href="#">
        <div className="producto_img">
        <img src={image} alt={title}/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>{title}</h1>
            <p>{category}</p>
            <p className="price">${price}</p>
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