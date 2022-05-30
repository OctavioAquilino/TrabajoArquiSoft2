import React,{ createContext, useState, useEffect} from "react";
import Data from '../Data.js';
export const DataContext = createContext();

export const Dataprovider = (props)=>{
    const[productos,setProductos]= useState([])

    useEffect(() =>{
        const producto =Data.items
        if(producto){
            setProductos(producto)
        }else{
            setProductos([])
        }
        
    },[])

    const value = {
        productos:[productos]
    }
    return(
        <DataContext.Provider value = {value}>
            {props.children}
        </DataContext.Provider>
    )
}