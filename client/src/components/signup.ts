function signup(content){
    const signupHolder = document.createElement("div");
    
    createInputs(signupHolder)
    submitButton(signupHolder)
    content.append(signupHolder)
}

function createInputs(holder){
    const usernameLabel = document.createElement("label")
    const username = document.createElement("input")
    usernameLabel.htmlFor = "username"
    username.type = "text"
    username.name = "username" //* how to give name property?
    username.className = "input-field"

    const passwordLabel = document.createElement("label")
    const password = document.createElement("input")
    passwordLabel.htmlFor = "password"
    password.type = "password"
    password.name = "username"
    password.className = "input-field"

    holder.append(usernameLabel,username,passwordLabel,password)
}

function submitButton(holder){
    const button = document.createElement("button")
    button.type = "submit"
    button.className = "signup-button"
    button.setAttribute("id","signup-button")

    holder.append(button)
}


export default signup