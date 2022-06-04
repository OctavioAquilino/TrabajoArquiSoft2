import React,{ createContext, useState, useEffect} from "react";

export const DataContext = createContext();

export const Dataprovider= (props)=>{

    const [productos,setProductos] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/product')
    .then((response) => response.json());
    setProductos(response);
    };
    useEffect(()=>{
    fetchApi();
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