import React from "react"
import Logo from "../../Imagenes/Logo1.png"
import {Link} from "react-router-dom";

export const Header = ()=>{
    return(
        <header>
        <Link to="/">
            <div className="logo">
                <img src={Logo} alt="logo" width="150"/>
            </div>
        </Link>
        <ul>
            <li>
               <Link to="/"> INICIO</Link> 
            </li>
            <li>
               <Link to="/productos"> PRODUCTOS</Link> 
            </li>
            <li>
               <Link to="/login"> LOG IN</Link> 
            </li>
        </ul>
        <div className="cart">
            <box-icon name='cart'></box-icon>
            <span className="item_total">0</span>
        </div>
        </header>
    )
}