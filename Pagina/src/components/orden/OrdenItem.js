import React from "react";
import { OrdenDetalleItem } from "./OrdenDetalleItem";
export const OrdenItem =(
    {id_order,
        monto_final,
        fecha,
        id_user,
        OrdersDetalle
})=>{
//poner la tabla
    return(
      <div>
        <div className="orden">
        <table>
        <thead>
           <tr>
             <th>Producto</th>
             <th>Precio Unitario</th>
             <th>Cantidad</th>
             <th>Total</th>
           </tr>
         </thead>
         <tbody>
         {
           OrdersDetalle.map((ordenDetalle) =>(
                <tr>
                <td>{ordenDetalle.nombre}</td>
                <td>{ordenDetalle.precio_unitario}</td>
                <td>{ordenDetalle.cantidad}</td>
                <td>{ordenDetalle.total}</td>
              </tr>
                ))
           }
         </tbody>
       </table>
          <table>
         <thead>
           <tr>
             <th>Monto Final</th>
             <th>Fecha</th>
           </tr>
         </thead>
         <tbody>
            <tr>
               <td>{monto_final}</td>
               <td>{fecha}</td>
             </tr>
         </tbody>
       </table>
       <br></br>
        </div>
        <div>
        </div>
        </div>
    )
}