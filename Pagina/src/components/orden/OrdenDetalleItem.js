import React from "react";
export const OrdenDetalleItem =(
    {id,
    precio_unitario,
    cantidad,
    total,
    nombre,
    id_product,
    id_order
})=>{

    return(
        <div className="ordenDetalle">
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
            <tr>
               <td>{nombre}</td>
               <td>{precio_unitario}</td>
               <td>{cantidad}</td>
               <td>{total}</td>
             </tr>
         </tbody>
       </table>
        </div>
    )
}