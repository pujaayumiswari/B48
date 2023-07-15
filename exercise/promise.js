//status promise: pending, fulfilled (resolved),rejected

let condition = false;
let promise = new Promise((resolve, reject) =>{
 if(condition){
  resolve("janji di tepati")
 }else{
  reject("janji ga di tepati")
 }
})

promise.then((value) =>{
  console.log(value)
}).catch((err) =>{
  console.log(err)
})