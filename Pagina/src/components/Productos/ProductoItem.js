import React from "react";
import Logo from "../../Imagenes/Logo1.png";

export const ProductoItem =()=>{
    return(
        <div className="producto">
        <a href="#">
        <div className="producto_img">
        <img src={Logo} alt="logo"/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>Title</h1>
            <p>Categoria</p>
            <p className="price">$120</p>
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