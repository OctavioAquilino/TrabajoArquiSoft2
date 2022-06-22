import React, {useState, useEffect} from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import Cookies from "universal-cookie";
import { OrdenItem } from "./OrdenItem";


const Cookie = new Cookies();

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
    let cookie = Cookie.get("user")
    let id_user;
    if(cookie!=undefined){
    let array = cookie.split(",")
     id_user = array[0]
    }
    else{
         id_user = "undefined"
    }
    const [ordenes,setOrdenes]=useState([]);
        async function Handle (id) {
            const response = await GetOrdersByIdUser(id)
            if (response.status == 400) {
                alert("NO HA REALIZADO NINGUNA ORDEN")
                
             }else{
                setOrdenes(response)
             }
            
            };
    useEffect(()=>{
        if(id_user!="undefined"){
        Handle(id_user);
        }
        else{
            alert("Debe loguearse")
            window.location.replace("/login")
        }
        })
        return(
            <>
            <h1 className="title">MIS ORDENES</h1>
            <div>
                <div className="ordenes">
            
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
   







