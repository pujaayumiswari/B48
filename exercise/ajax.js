const xhr = new XMLHttpRequest()


//crud
xhr.open('get', "https://your.url", true) //true = asyncronous, false = syncronous

xhr.onload = function(){} //mengecek status
xhr.onerror = function(){} // menampilkan error ketika internet mati atau dari server
xhr.send()