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
        <div className="orden">
            <p>{monto_final},{fecha}</p>{
           OrdersDetalle.map((ordenDetalle) =>(
                <OrdenDetalleItem key={ordenDetalle.id}
                id={ordenDetalle.id}
                nombre={ordenDetalle.nombre}
                precio_unitario={ordenDetalle.precio_unitario}
                cantidad={ordenDetalle.cantidad}
                total={ordenDetalle.total}
                id_product={ordenDetalle.id_product}
                id_order={ordenDetalle.id_order}
                /> 
                ))
           }
        </div>
    )
}