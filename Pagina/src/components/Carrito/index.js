import React, { useContext, useEffect, useState } from 'react'
import Logo from "../../Imagenes/Logo1.png"
import {DataContext} from './Provider'
export const Carrito = () => {
    const value= useContext(DataContext)
    const [menu,setMenu]= value.menu;
    const [carrito,setCarrito] = value.carrito;
    const [total]=value.total;

    const tooglefalse = ()=>{
        setMenu(false);
    }

    const show1 = menu? "carritos show" : "carritos";
    const show2 = menu? "carrito show" : "carrito";

    const removeProducto =id=>{
        if(window.confirm){
            carrito.forEach((item,index)=>{
                if(item.id=== id){
                item.stock=1;
                carrito.splice(index,1)  
                }

            })
            setCarrito([...carrito])
        }
        
    }
    const resta = id =>{
        carrito.forEach((item)=>{
            if(item.id=== id){
                item.stock ===1?item.stock=1: item.stock -=1
            }
            setCarrito([...carrito])
        })
    }
    /*---En la suma tendria que iniar en uno, pero no se como*/
    const suma = id =>{
        carrito.forEach((item)=>{
            if(item.id=== id){
                item.stock+=1
            }
            setCarrito([...carrito])
        })
    }

  return (
    <div className={show1}>
        <div className={show2}>
            <div className='carrito_close' onClick={tooglefalse}>
                <box-icon name="x"></box-icon>
            </div>
            <h2>SU CARRITO</h2>
            
                
            <div className='carrito_center'>
                {
                    carrito.length ===0? <h2>CARRITO VACIO</h2>:
                    <>
                    {
                    carrito.map((producto)=>(
                <div className='carrito_item'>
                    <img  src={producto.picture_url} className='logo' alt=''/>
                    <div>
                    <h3>{producto.name}</h3>
                   <p className='price'>${producto.base_price}</p>
                    </div>
                <div>
                <box-icon name="up-arrow" type="solid" onClick={()=>suma(producto.id)}></box-icon>
                <p className='cantidad'>{producto.stock}</p>
                <box-icon name="down-arrow" type="solid" onClick={()=>resta(producto.id)}></box-icon>
                </div>
                <div className='remove_item' onClick={()=>removeProducto(producto.id)}>
                    <box-icon name="trash"></box-icon>
                </div>
                </div>
                ))}
                </>
                }
            </div>
            
            <div className='carrito_footer'>
                <h3>Total: ${total}</h3>
                <button className='btn'>Payment</button>
            </div>
        </div>
    </div>
  )
}
