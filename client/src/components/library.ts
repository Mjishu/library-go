function callApi(){
    fetch("/api/library")
    .then(res => res.json())
    .then(data => console.log(data))
}

function makeLibrary(){
    
}