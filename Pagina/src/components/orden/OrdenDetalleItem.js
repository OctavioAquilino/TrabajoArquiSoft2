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
            <p>{nombre},{precio_unitario},{cantidad},{total}</p>
        </div>
    )
}