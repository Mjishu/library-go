
interface userInfo {
    username: string,
    password: string
}

const userInfo: userInfo = {
    username: "",
    password: "",
}


const form = document.getElementById("sign-up-form") as HTMLFormElement;
const usernameInput = document.getElementById("username") as HTMLInputElement;
const passwordInput = document.getElementById("password") as HTMLInputElement;

usernameInput.addEventListener("input", (e:Event) => {
    const target = e.target as HTMLInputElement;
    userInfo.username = target.value
})

passwordInput.addEventListener("input", (e:Event)=>{
    const target = e.target as HTMLInputElement
    userInfo.password = target.value
})

form.addEventListener("submit",(e:Event)=>{
    e.preventDefault();

    console.log(userInfo)

    const fetchParams = {
        method: 'POST',
        headers:{
            "Content-Type":"application/json"
        },
        body: JSON.stringify(userInfo)
    }

    fetch("/api/login",fetchParams)
})