
import React,{ useState} from "react"
import "./Prueba1.css"
import Cookies from "universal-cookie";
import {CookieUser} from "../cookies/cookieUser"
const Cookie =new Cookies();
export default function IndexLogin(){
    const[user,setUser]= useState("");
    const[password,setPassword] = useState("");
  
    const onChangeUser =  (user)=>{
        setUser(user.target.value);
        
    }
    
    const onChangePas = (password)=>{
    setPassword(password.target.value)};

    
    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({user_name : user, password : password })
    };

    const login = async()=>{
        fetch('http://localhost:8090/login',requestOptions)
        .then(response => {if (response.status == 400) {
           alert("Datos incorrectos")
           window.location.reload();
           return response.json()
        }
        if(response.status==201){
          window.location.replace("/")
          return response.json()
        }
        return response.json()})
        .then(response => {
            Cookie.set("user", response.id_user + "," + response.token+ ";", {path: "/"})
    })
   
    };
   
    const handleSubmit= (event)=>{
        event.preventDefault();
        login();

    };

    return(
      <div className="app2">
        <div className="login-form">
        <form onSubmit={handleSubmit} >
        <ul className="ul">
        
        <h1 className="login">LOGIN</h1>
        
          <li>
        <input id="user" type={"text"} placeholder="user" onChange={onChangeUser} value ={user} required></input>
          </li>
        <li>
        <input id="password" type={"password"} placeholder="password" onChange={onChangePas} value={password} required></input>
        </li>
        <input type="submit" value="Log In"></input>
        </ul>
        
        </form>
        </div>
        </div>

    );




}