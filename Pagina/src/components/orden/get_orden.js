import React, {useState} from "react";
import "bootstrap/dist/css/bootstrap.min.css";
//import { OrdenDetalleItem } from "./OrdenDetalleItem";
import { OrdenItem } from "./OrdenItem";
//import {IndexLogin} from "../login/IndexLogin"

async function GetOrdersByIdUser(id) {
    return fetch('http://localhost:8090/orderUser/' +id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(data => data.json())
   }
export const GetOrders = ()=>{

    const [ordenes,setOrdenes]=useState([]);
    //const [ordenReq,setOrdenReq]=useState([]);
    //const [ordenDetalle,setOrdenDetalle]=useState([]);
    
    /*const getOrdenes = async()=>{
        const response = await fetch('http://localhost:8090/orderUser/1')
        .then((response) => response.json());
        setOrdenes(response);
        //setOrdenDetalle(orden.OrdersDetalle)
        };
        const handleSubmit= (event)=>{
            event.preventDefault();
            //alert(busqueda)
           getOrdenes();
        };*/
        async function Handle (id) {
            const response = await GetOrdersByIdUser(id)
            if (response.status == 400) {
                alert("NO HA REALIZADO NINGUNA ORDEN")
                //window.location.reload();
             }else{
                setOrdenes(response)
                console.log(response);
             }
            //setOrdenes(response)
            //console.log(response);
            };
        return(
            <>
            <h1 className="title">MIS ORDENES</h1>
            <div>
                <div className="ordenes">
            <input className="mostrar" value = "Mostrar" type = "button" onClick = {()=>Handle(1)}/>
            {
                ordenes.map(orden =>(
                    <div>
                        <OrdenItem key={orden.id_order}
                    id_order = {orden.id_order}
                    monto_final = {orden.monto_final}
                    fecha = {orden.fecha}
                    id_user = {orden.id_user}
                    OrdersDetalle = {orden.OrdersDetalle}
                    /> 
                    </div>
                )
                )
            }
            </div>
            </div>
            
            </>      
        )
}






