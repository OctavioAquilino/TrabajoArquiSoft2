import React,{useContext, useState} from "react"
import Logo from "../../Imagenes/Logo1.png"
import {Link} from "react-router-dom";
import { DataContext } from "../Carrito/Provider";

export const Header = ()=>{
    const value = useContext(DataContext);
    const [menu,setMenu]= value.menu;
    const [carrito]= value.carrito

    const toogleMenu = ()=>{
        setMenu(!menu);
    }
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
            <li>
               <Link to="/order"> MIS ORDENES</Link> 
            </li>
        </ul>
        <div className="cart" onClick={toogleMenu}>
            <box-icon name='cart'></box-icon>
            <span className="item_total"> {carrito.length}</span>
        </div>
        </header>
    )
}