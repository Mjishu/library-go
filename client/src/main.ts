import './style.css'
import typescriptLogo from './typescript.svg'
import viteLogo from '/vite.svg'
import signup from "./signup.ts"

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
  </div>
`

const app = document.getElementById("app")
signup(app)


//setupCounter(document.querySelector<HTMLButtonElement>('#counter')!)
