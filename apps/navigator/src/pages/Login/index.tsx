import api from '@/models/api'

export function Login() {
  const username = prompt("Enter username");
  const password = prompt("Enter password");
  
  api.authenticate({ username, password }).then(() => {
    window.location.href = "/"
  }).catch((err) => {
    console.log("ERR", err)
    alert("Login failed")
    window.location.reload()
  })
  
	return <div />
}
