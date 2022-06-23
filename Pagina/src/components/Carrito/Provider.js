import React,{ createContext, useState, useEffect} from "react";

export const DataContext = createContext();

export const Dataprovider= (props)=>{
    const [menu,setMenu]=useState(false);
    const [carrito,setCarrito]=useState([]);
    const [productos,setProductos] = useState([]);
    const [total,setTotal] = useState(0);
    const fetchApiProd = async()=>{
        const response = await fetch('http://localhost:8090/product')
        .then((response) => response.json());
        setProductos(response);
        };
    
        useEffect(()=>{
        fetchApiProd();
        },[])

    const addCarrito=(id)=>{
        const check = carrito.every(item=>{
            return item.id!== id;
        })
        if(check){
            const data= productos.filter(producto=>{
                return producto.id === id
            })
            setCarrito([...carrito,...data])
        }else{
            alert("El producto ya fue agregado")
        }
    }
    useEffect(()=>{
        const getTotal=()=>{
            const res = carrito.reduce((prev,item)=>{
                return prev + (item.base_price * item.stock)
            },0)
            setTotal(res)
        }
        getTotal()
    },[carrito])

/*----En esta parte tendria que ir la memoria del carrito */
    /*useEffect(()=>{
        const dataCarrito=JSON.parse(localStorage.getItem('dataCarrito'))
        if(dataCarrito){
            setCarrito(dataCarrito)
        }
    },[])
    useEffect(()=>{
        localStorage.setItem('dataCarrito',JSON.stringify(carrito))
    },[carrito])*/

    const value = {
        menu:[menu,setMenu],
        addCarrito:addCarrito,
        carrito:[carrito,setCarrito],
        total:[total,setTotal]
    }
    return(
        <DataContext.Provider value = {value}>
            {props.children}
        </DataContext.Provider>
    )
}