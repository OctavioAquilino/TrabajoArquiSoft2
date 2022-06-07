import React, { useState , useEffect } from "react";
import './Prueba1.css';
import {Link} from "react-router-dom";
import { faBorderStyle } from "@fortawesome/free-solid-svg-icons";
//import { Link } from 'react-router-dom';

export const IndexLogin =(username,password)=>{

    const[token,setToken]= useState(false);
    
    const fetchApi = async()=>{

    const response = await fetch('http://localhost:8090/login', {
      method: 'OPTIONS',
      headers: {'Content-Type': 'application/json'},
      //credentials: 'include',
      body: JSON.stringify({
        username,
          password
      })
  }); 
  
  token = await response.json("token");
    setToken(true);
  
    };

  useEffect(()=>{
      fetchApi();
      },[])




const renderForm = (
  <div className="form">
    <div className="title">INGRESA TU CUENTA</div>
    <form onSubmit={fetchApi()}>
      <div className="input-container">
        <label>UserName </label>
        <input type="text" name="username" required />
        <label>Password </label>
        <input type="password" name="password" required />
      </div>
      <div className="button-container">
        <input type="submit" value="Send Request" />
      </div>
    </form>
  </div>
);

return (
  <div className="app">
    <div className="login-form">
      <div className="title"></div>
        {token? <Link replace to="/"/> : renderForm}
    </div>
  </div>
);
}
























/*async function getUserByID(username,userpass) {
  alert(username)
 return fetch('http://127.0.0.1:8090/login' , {
   method: 'POST',
   headers: {
     'Content-Type': 'application/json'
   },
   body: JSON.stringify({username, userpass}),
   
 })
 //.then((response) => response.json())
   .then(data => data.json())
}


export function Login() {

  const [userData, setUserData] = useState({});
  const [isUser, setIsUSer] = useState(false);
  
  const handleSubmit = async event => {
    //Prevent page reload
    event.preventDefault();

    // Obtenemos Textos
    var {uid, upass}  = document.forms[0];
    alert(upass.value)
    // Find user login info
    const user = await getUserByID(uid.value,upass.value);


    setUserData(user);
    setIsUSer(true);

  };
  const renderForm = (
    <div className="form">
      <div className="title">INGRESA TU CUENTA</div>
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>UserName </label>
          <input type="text" name="uid" required />
          <label>Password </label>
          <input type="password" name="upass" required />
        </div>
        <div className="button-container">
          <input type="submit" value="Send Request" />
        </div>
      </form>
    </div>
  );

  /*const Accept = (
      <Link to="/"> INICIO</Link> 
      <h1> {userData.name}</h1>
  )
*/
  



