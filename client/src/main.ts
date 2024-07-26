import signup from "./components/signup"

console.log("connected")

const htmlContent = document.getElementById("content");

if (htmlContent){
    const title = document.createElement("h1");
    title.innerHTML = "This is the title of the library pages!"
    htmlContent.appendChild(title)
    signup(htmlContent)

  //  document.addEventListener("click",)

}else{
    console.log("issue accessing htmlContent")
}
