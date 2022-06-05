import React, { useState } from "react";
import './Prueba1.css';
import {Link} from "react-router-dom";

async function getUserByID(id) {
 return fetch('http://127.0.0.1:8090/user/' + id, {
   method: 'GET',
   headers: {
     'Content-Type': 'application/json'
   }
 })
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

    alert(upass.value);
    // Find user login info
    const user = await getUserByID(uid.value);
    
    
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

  const Accept = (
      <Link to="/"> INICIO</Link> 
  )

  return (
    <div className="app2">
      <div className="login-form">
          {isUser? Accept : renderForm}
      </div>
    </div>
  );
}


