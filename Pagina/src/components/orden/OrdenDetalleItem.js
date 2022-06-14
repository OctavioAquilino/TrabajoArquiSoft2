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
         
         <tbody>
            <tr>
               <td>{nombre}----    </td>
               <td>{precio_unitario}  -----    </td>
               <td>{cantidad} ----     </td>
               <td>{total}    </td>
             </tr>
         </tbody>
       </table>
        </div>
    )
}