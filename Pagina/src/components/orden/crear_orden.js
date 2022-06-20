import React, { useContext, useEffect, useState} from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import { OrdenDetalleItem } from "./OrdenDetalleItem";
export const CrearOrden = ()=>{

    const [ordenRes,setOrdenRes]=useState([]);
    //const [ordenReq,setOrdenReq]=useState([]);
    //const [ordenDetalle,setOrdenDetalle]=useState([]);

    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({
             //{"id_user":1,"OrdersDetalle":[{"id_product":15,"cantidad":2},{"id_product":10,"cantidad":1}]}'
             id_user: 1,
             OrdersDetalle:[{"id_product":15,"cantidad":2},{"id_product":10,"cantidad":1}]
        })
    };
    const crearOrden = async()=>{
        const response = await fetch('http://localhost:8090/order',requestOptions)
        .then((response) => response.json());
        setOrdenRes(response);
        //setOrdenDetalle(orden.OrdersDetalle)
        };
        const handleComprar= (event)=>{
            event.preventDefault();
            //alert(busqueda)
            crearOrden();
    
        };

    return(
            <>
            <h1 className="title"> ORDEN</h1>
        <div className="containerInput" >   
        <input value = "Comprar" type = "button" onClick = {handleComprar}/>
        </div>
            </>
        )
}