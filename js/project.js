let dataBlog = [];

function addProject(event) {
  event.preventDefault();

  let projectName = document.getElementById("input-project").value
  let startDate = document.getElementById("input-start").value
  let endDate = document.getElementById("input-end").value
  let description = document.getElementById("description").value
  let image = document.getElementById("input-image").files


  let start = new Date(startDate)
  let end = new Date(endDate)
  let Distance = end - start
  console.log(Distance)

  let DistanceSecond = Math.floor(Distance/1000)
  let DistanceMinutes = Math.floor(DistanceSecond/60)
  let DistanceHours = Math.floor(DistanceMinutes/60)
  let DistanceDay = Math.floor(DistanceHours/24)
  let DistanceWeeks = Math.floor(DistanceDay/7)
  let DistanceMonth = Math.floor(DistanceWeeks/4)
  let DistanceYear = Math.floor(DistanceMonth/12)

  let DistanceTime = "";

  //pengkondisian//
  if (DistanceDay < 7) {
    DistanceTime = DistanceDay + "hari";
} else if (DistanceWeeks < 4) {
  DistanceTime = DistanceWeeks + "minggu";
} else if (DistanceMonth < 12) {
  DistanceTime = DistanceMonth + "bulan";
} else  {
  DistanceTime = DistanceYear + "tahun";
}
  


  const javascript =  `<i class="fa-brands fa-square-js"></i>`
  const reactJS = ` <i class="fa-brands fa-react"></i>`
  const golang = ` <i class="fa-brands fa-golang"></i>`
  const java = `<i class="fa-brands fa-java"></i>`

  let cbjavascript = document.getElementById("javascript").checked ? javascript : "";
  let cbreactJS = document.getElementById("reactJS").checked ? reactJS : "";
  let cbgolang = document.getElementById("golang").checked ? golang : "";
  let cbjava = document.getElementById("java").checked ? java : "";
  
  

  image = URL.createObjectURL(image[0]);
  console.log(image)

  //push project//
  let project = {
    projectName,
    DistanceTime,
    cbjavascript,
    cbreactJS,
    cbgolang,
    cbjava,
    description,
    postAt :new Date(),
    image,
  }
    dataBlog.push(project);
    console.log(project)
    console.log(dataBlog)

    renderProject();

}

//menampilkan blog//

function renderProject() {
  document.getElementById("contents").innerHTML = "";
  for (let i = 0; i < dataBlog.length; i++ ){
    document.getElementById("contents").innerHTML +=`
    <div class="row" id="contents">
      <div class="column">
        <div class="img">
          <img src="${dataBlog[i].image}"/>
        </div>
        <a href="projectdetail.html" style="color: black;">
        <h3>${dataBlog[i].projectName}</h3>
        </a>
        <p>${getFullTime(dataBlog[i].postAt)}</p>
        <p> Durasi : ${dataBlog[i].DistanceTime}
        <p>${dataBlog[i].description}</p>
        <div class="technologies-icon">
        ${dataBlog[i].cbjavascript}
        ${dataBlog[i].cbreactJS}
        ${dataBlog[i].cbgolang}
        ${dataBlog[i].cbjava}
        </div>        
        <div class="button-container">
          <button class="edit-button">Edit</button>
          <button class="delete-button">Delete</button>
        </div>       
      </div>  
`
  }


}
function getFullTime(time) {
  let bulan = ["Jan", "Feb","March", "Apr", "May", "Jun", "Jul", "Aug", "Sept","Oct", "Nov","Desc"];
  let tanggal = time.getDate();
  let indexBulan = time.getMonth();
  let tahun = time.getFullYear();
  let hours = time.getHours();
  let minutes = time.getMinutes();

  if (hours < 10) {
    hours = "0" + hours
  }
  if (minutes < 10) {
    minutes = "0" + minutes
  }

  
  

  return `${tanggal} ${bulan[indexBulan]}  ${tahun} ${hours}:${minutes} WIB`;
}