import Cookies from "universal-cookie";

const Cookie= new Cookies()
export function CookieUser(id,token){
  
    let cookie = Cookie.get("user");
   
    if(cookie == undefined){
      Cookie.set("user", id + "," + token + ";", {path: "/"});
      return
    }
    return
  }
   