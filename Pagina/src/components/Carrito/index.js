import React, { useState, useEffect, List, Checkbox} from "react";
import "./carrito.css"
import Cookies from "universal-cookie";
import { CrearOrden } from "../orden/crear_orden";
const Cookie = new Cookies();

const id_user = setUser()

function setUser (){
  let cookieUser = Cookie.get("user")

  if(cookieUser!=undefined){
  let array = cookieUser.split(",")
  return array[0]
  }else{
    return "undefined"
  }
 
  
}

function auxiliar(){

  CrearOrden();
  }

async function getProductById(id){
  return fetch("http://localhost:8090/product/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}




async function getCartProducts(){
 
  let items = []
  let a = Cookie.get("cart"+id_user).split(";")

  for (let i = 0; i < a.length; i++){
    let item = a[i];
    if(item != ""){
      let array = item.split(",")
      let id = array[0]
      let quantity = array[1]
      let product = await getProductById(id)
      product.quantity = quantity;
      items.push(product)
    }
  }
  return items
}


function getOptions(n){
  let options = []
  for(let i=1; i <= n; i++){
    options.push(i)
  }
  return options.map((option) =>
    <option value={option}> {option} </option>
)
}

function remove(n, p_id){

  let cookie = Cookie.get("cart"+id_user);
  let newCookie = ""
  let toCompare = cookie.split(";")
  let isEmpty = false
  toCompare.forEach((item) => {
    if(item != ""){
      let array = item.split(",")
      let item_id = array[0]
      let item_quantity = array[1]
      if(p_id == item_id){
        item_quantity = Number(item_quantity) - n
        if(item_quantity == 0){
          isEmpty = true
        }
      }
      if(p_id == item_id && !isEmpty){
        newCookie += item_id + "," + item_quantity + ";"
      }
      else if (p_id != item_id){
        newCookie += item_id + "," + item_quantity + ";"
      }
    }
  });
  cookie = newCookie
  Cookie.set("cart"+id_user, cookie, {path: "/"})
  window.location.reload()
  return
}


function showProducts(products){
  
  return products.map((product) =>

   <div className="producto">
   <div obj={product} key={product.id} >
   <div>
        <a href="#">
        <div className="producto_img">
            <img className="image" src={product.picture_url} alt=""/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>{product.name}</h1>
            <p>{product.description}</p>
            <p className="price">U$S {product.base_price}</p>
            <h3 className="Remove"> Remover </h3>
       <select id={"removeSelect" + product.id}>
        {getOptions(product.quantity)}
       </select>
       <button className="remove" onClick={() => remove(document.getElementById("removeSelect" + product.id).value, product.id)}>x</button>
       <h1 className="amount"> Cantidad: </h1>
       <h1 className="number"> {product.quantity} </h1>
       <h1 className="subtotal"> Subtotal: U$S {product.quantity * product.base_price} </h1>
        </div>
        </div>
        </div>
   </div>
 )
}

async function setCart(setter, setterTotal){
  let total = 0;
  await getCartProducts().then(response => {
    setter(response)
    response.forEach((item) => {
      total += item.base_price * item.quantity;
    });
    setterTotal(total)
  })
}

function Cart(){
  const [cartProducts, setCartProducts] = useState([]);
  const [total, setTotal] = useState(0);
 
  if (cartProducts.length <= 0 ){
    setCart(setCartProducts, setTotal)
  }

  

  const renderOrderButton = (
    <div className="emptySpace">
      <span> Total a Pagar: U$S {total} </span>
    </div>
  )

  return (
    <div>
      <h1 className="title"> TU CARRITO</h1>
      <div className="productos">
        {Cookie.get("cart"+id_user) ? showProducts(cartProducts) : <a></a>}
      </div>
      
      {cartProducts.length>=1 ? 
      <div className="pago">
        {renderOrderButton}
        <button className="payment"  onClick={()=> auxiliar()}> Payment</button>
      </div>:
      <div><h2>Tu Carrito esta vacio</h2></div>
}
</div>
  );
}

export default Cart;