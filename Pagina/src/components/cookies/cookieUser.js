import Cookies from "universal-cookie";

const Cookie= new Cookies()
export function CookieUser(id,token){
  
    let cookie = Cookie.get("user");
   
    if(cookie == undefined){
      Cookie.set("user", id + "," + token + ";", {path: "/"});
     
      return
    }
    /*let newCookie = ""
    let isNewItem = true
    let toCompare = cookie.split(";")
    let total = 0;
    toCompare.forEach((item) => {
      if(item != ""){
        let array = item.split(",")
        let item_id = array[0]
        let item_quantity = array[1]
        if(id == item_id){
          item_quantity = Number(item_quantity) + 1
          isNewItem = false
        }
        newCookie += item_id + "," + item_quantity + ";"
        total += Number(item_quantity);
      }
    });
    if(isNewItem){
      newCookie += id + ",1;"
      total += 1;
    }
    cookie = newCookie
    Cookie.set("cart", cookie, {path: "/"})
   */
    return
  }
   