
interface userInfo {
    username : string,
    email: string,
    password: string
}

const userInfo: userInfo ={
    username: "",
    email: "",
    password: ""
}

const form = document.getElementById("sign-up-form") as HTMLFormElement;
const usernameInput = document.getElementById("username") as HTMLInputElement;
const passwordInput = document.getElementById("password") as HTMLInputElement;
const emailInput = document.getElementById("email") as HTMLInputElement

usernameInput.addEventListener("input",(e:Event) =>{
    const target = e.target as HTMLInputElement;
    userInfo.username = target.value;
})

emailInput.addEventListener("input",(e:Event) => {
    const target = e.target as HTMLInputElement
    userInfo.email = target.value;
})

passwordInput.addEventListener("input",(e:Event) => {
    const target = e.target as HTMLInputElement;
    userInfo.password = target.value
})

form.addEventListener("submit",(e:Event)=>{
    e.preventDefault();
    console.log(userInfo);
})
