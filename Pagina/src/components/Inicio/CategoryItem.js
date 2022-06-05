import React from "react";

export const CategoryItem =(
    {id,
    name,
    description
})=>{
    return(
        <div className="categoria">
        <div className="categoria_footer">
            <h1>{name}</h1>
            <p>{description}</p>
        </div>
        </div>
    )
}