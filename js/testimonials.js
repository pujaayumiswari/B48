//inheritance
//poymorphism
//abstraction
//encapsulation


// class Testimonials {
//   #quote =""
//   #Image =""
//   constructor(quote, Image){
//     this.#quote =quote
//     this.#Image =Image
//   }
 
//   get quote(){
//     return this.#quote
//   }

//   get Image(){
//     return this.#Image
//   }
  
//   get user(){
//     throw new error('there is must be user to make testimonials')
//   }

// get testimonialHtml(){
//   return`<div class="container" id="testimonials">
//   <div class="card">
//     <img src="${this.Image}" />
//     <div class="intro">
//     <p>"${this.quote}"</p>
//       <h4>${this.user}</h4>
//     </div>
//   </div>`
// }
// }
// class userTestimonial extends Testimonials{
//   #user =""
//   constructor(user, quote, Image){
//     super(quote, Image)
//     this.#user = user
//   }
//   get user(){
//     return "user : "+ this.#user
//   }
// }

// class companyTestimonial extends Testimonials{
//   #company =""
//   constructor(company, quote, Image){
//     super(quote, Image)
//     this.#company = company
//   }
//   get user(){
//     return "company : " + this.company
//   }
// }

// const testimonial1 = new userTestimonial("Nam Dosan", "첫 인상이 좋아요", src="/image/namdosan.jpg" )
// const testimonial2 = new userTestimonial("Bae Suzy", "멋있다", src="/image/suzy.jpg" )
// const testimonial3 = new userTestimonial("Han Ji pyeoung", "재킷이 잘 어울려요",src="/image/pyeoung.jpg" )

// let testimonialData = [testimonial1, testimonial2, testimonial3]
// let testimonialHtml =""

// for (let i = 0; i< testimonialData.length; i++){
//   testimonialHtml += testimonialData[i].testimonialHtml
// }

// document.getElementById("testimonials").innerHTML =testimonialHtml

// const testimonialData =[
//   {
//     user: "Nam Dosan",
//     quote: "첫 인상이 좋아요",
//     Image: src="image/namdosan.jpg",
//     rating: 5

//   },
//   {
//     user: "Bae Suzy",
//     quote: "첫 인상이 좋아요",
//     Image: src="image/suzy.jpg",
//     rating: 4
//   },
//   {user: "Han Ji Pyeoung",
//     quote: "첫 인상이 좋아요",
//     Image: src="image/Pyeoung.jpg",
//     rating: 3
// }
// ]



// function allTestimonial(){
//   let testimonialHtml =""

//   testimonialData.forEach((card,index) => {
//     testimonialHtml += `<div class="container" id="testimonials">
//     <div class="card">
//      <img src="${card.Image}" />
//      <div class="intro">
//      <p>"${card.quote}"</p>
//        <h4>${card.user}</h4>
//        <p class="author">${card.rating}<i class="fa-solid fa-star"></i></p>
//       </div>
//     </div>`
//   })
//   document.getElementById("testimonials").innerHTML = testimonialHtml

// }

// //eksekusi awal /default
// allTestimonial()
// function filterTestimonial (rating){
//   let filteredTestimonialHTML = ""
 
//    const filteredData = testimonialData.filter((card) =>{
//     return card.rating === rating
//    })

//    filteredData.forEach((card) =>{
//     filteredTestimonialHTML += `<div class="container" id="testimonials">
//     <div class="card">
//      <img src="${card.Image}" />
//      <div class="intro">
//      <p>"${card.quote}"</p>
//        <h4>${card.user}</h4>
//        <p class="author">${card.rating}<i class="fa-solid fa-star"></i></p>
//       </div>
//     </div>`
//    })

//    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
//   }

const promise = new Promise((resolve, reject) =>{
  const xhr = new XMLHttpRequest()
  xhr.open("GET", "https://api.npoint.io/4b21cc129aa29eac4b72", true)
  xhr.onload = function(){
    if (xhr.status === 200){
      resolve(JSON.parse(xhr.responseText))
    }else if (xhr.status >= 400){
      reject("error loading data")
    }
  }
  xhr.onerror = function(){
    reject("network error")
  }
  xhr.send()
})

 let testimonialData = []
 async function getData(rating){
  try{
    const response = await promise
    console.log(response)
    testimonialData = response
    allTestimonial()
  }catch (err){
    console.log(er)
  }
 }

 getData()

function allTestimonial(){
  let testimonialHTML = ""
  testimonialData.forEach((card, index) =>{
    testimonialHTML += `<div class="container" id="testimonials">
         <div class="card">
         <img src="${card.Image}" />
         <div class="intro">
         <p>"${card.quote}"</p>
           <h4>${card.user}</h4>
           <p class="author">${card.rating}<i class="fa-solid fa-star"></i></p>
          </div>
       </div>`
  })
  document.getElementById("testimonials").innerHTML = testimonialHTML
}



function filteredTestimonial(rating) {
  let filteredTestimonialHTML =""
  const filteredData = testimonialData.filter((card) =>{
    return card.rating === rating
  })

  filteredData.forEach((card) =>{
    filteredTestimonialHTML += `<div class="container" id="testimonials">
    <div class="card">
    <img src="${card.Image}" />
    <div class="intro">
    <p>"${card.quote}"</p>
      <h4>${card.user}</h4>
      <p class="author">${card.rating}<i class="fa-solid fa-star"></i></p>
     </div>
  </div>`
}) 
  document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}