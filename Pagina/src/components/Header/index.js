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
               <Link to="/" className="botones_menu">
                <span></span>
                <span></span>
                <span></span>
                <span></span>
               INICIO
               </Link> 
            </li>
            <li>
               <Link to="/productos" className="botones_menu">
               <span></span>
                <span></span>
                <span></span>
                <span></span>
               PRODUCTOS
               </Link> 
            </li>
            <li>
               <Link to="/order" className="botones_menu"> 
               <span></span>
                <span></span>
                <span></span>
                <span></span>
                MIS ORDENES
                </Link> 
            </li>
            <li>
               <Link to="/login" className="botones_menu">
               <span></span>
                <span></span>
                <span></span>
                <span></span>
                LOG IN
                </Link> 
            </li>
        </ul>

        <Link to="/cart">
        <div className="cart">
                <box-icon name='cart'></box-icon>
            </div>
        </Link>
        </header>
    )
}