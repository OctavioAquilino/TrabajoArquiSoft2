import React, { useContext, useEffect, useState} from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import { OrdenDetalleItem } from "./OrdenDetalleItem";
import { OrdenItem } from "./OrdenItem";
export const GetOrders = ()=>{

    const [ordenes,setOrdenes]=useState([]);
    //const [ordenReq,setOrdenReq]=useState([]);
    //const [ordenDetalle,setOrdenDetalle]=useState([]);

   
    const getOrdenes = async()=>{
        const response = await fetch('http://localhost:8090/orderUser/1')
        .then((response) => response.json());
        setOrdenes(response);
        //setOrdenDetalle(orden.OrdersDetalle)
        };
        const handleSubmit= (event)=>{
            event.preventDefault();
            //alert(busqueda)
           getOrdenes();
    
        };

        return(
            <>
            <h1 className="title">MIS ORDENES</h1>
            <div className="containerInput" >
            <input
        value = "Mostrar"
         type = "button"
        onClick = {handleSubmit}
        />
            <div className="ordenes">
            {
                ordenes.map(orden =>(
                    <OrdenItem key={orden.id_order}
                    id_order = {orden.id_order}
                    monto_final = {orden.monto_final}
                    fecha = {orden.fecha}
                    id_user = {orden.id_user}
                    OrdersDetalle = {orden.OrdersDetalle}
                    /> 
                )
                )
              
            }
 
</div>
            
          </div>
    
           
            
            </>
        )

}






