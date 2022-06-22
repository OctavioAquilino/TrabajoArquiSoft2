import React from "react"
import Logo from "../../Imagenes/Logo1.png"
import {Link} from "react-router-dom";
import Cookies from "universal-cookie";

function LogOut(){
    var resultado = window.confirm('Estas seguro?');
if (resultado === true) {
    Cookie.set("user", "undefined,undefined", { path: "/" });
    window.location.reload();
} else { 
    window.alert('Pareces indeciso');
}
}

  
const Cookie = new Cookies();
    let cookie = Cookie.get("user")
    let id_user;
    if(cookie!=undefined){
    let array = cookie.split(",")
     id_user = array[0]
    }
    else{
         id_user = "undefined"
    }
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
               
                {id_user!="undefined"?
                <button className="btn" onClick={()=>LogOut()}>
                Log Out
            </button>:

            <li>
               <Link to="/login" className="botones_menu">
               <span></span>
                <span></span>
                <span></span>
                <span></span>
                Log In
                </Link> 
            </li>
            }
        </ul>

        <Link to="/cart">
        <div className="cart">
                <box-icon name='cart'></box-icon>
            </div>
        </Link>
        </header>
    )
}