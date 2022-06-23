import React, { useContext, useEffect, useState} from "react";
import "bootstrap/dist/css/bootstrap.min.css";

import Cookies from "universal-cookie";

const Cookie = new Cookies();

function vaciarCarrito(){
    alert("vaciando carro")
    document.cookie = "cart=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    alert("Carrito vaciado")
}
export const CrearOrden = ()=>{
    
    let cookie = Cookie.get("user")
    let id_user = parseInt(cookie.split(",")[0]);
    let orderDetail ={'id_product':0,'cantidad':0}
    let ordersDetail = []
    
    let a = Cookie.get("cart").split(";")
  
    for (let i = 0; i < a.length; i++){
      let item = a[i];
      if(item != ""){
        let array = item.split(",")
         orderDetail.id_product = parseInt(array[0])
         orderDetail.cantidad =parseInt(array[1])
        ordersDetail.push(orderDetail)
      }
    }
   

    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({
            
             id_user: id_user,
             OrdersDetalle:ordersDetail
        

    })
    };
    const crearOrden = async()=>{
        fetch('http://localhost:8090/order',requestOptions)
        .then(response => {if (response.status != 201) {
            alert("Error en la compra")
           
            window.location.reload();
            return response.json()
         }else{
            alert("Compra realizada con exito")
            vaciarCarrito()
            window.location.replace("/")
            return response.json()
          }})

        
        };
        
    crearOrden();

}