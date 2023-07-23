//DOM = document object model
//(manipulasi elemen html di javascript)
//= sama dengan 
//== Perbandingan
//=== Identitas harus sama


function submitData() {
  
  let name = document.getElementById("name").value
  let email = document.getElementById("email").value
  let phone = document.getElementById("Phone").value
  let subject = document.getElementById("input-subject").value
  let message = document.getElementById("message").value
  
  if(name === ""){
    return alert('nama tidak boleh kosong!')
  }else if(email ===""){
    return alert('email harus di isi!')
  }
  else if(phone ===""){
    return alert('no telepon harus di isi!')
  }
  else if(subject ===""){
    return alert('subject harus di isi!')
  }
  else if(message ===""){
    return alert('pesan harus di isi!')
  }

  let objectData = {
    name: name,
    email,
    phone,
    subject,
    message
}

let arrayData = [name, email, phone, subject, message]

console.log(objectData)

  let emailreceiver="payumiswari@gmail.com";
  let a = document.createElement("a");
  a.href=`mailto:${emailreceiver}?subject=${subject}&body=halo, nama saya ${name}, ${message}.silahkan kontak saya di nomor ${phone}`;
  a.click();
  
  console.log(name);
  console.log(email);
  console.log(phone);
  console.log(subject);
  console.log(message);
  
}