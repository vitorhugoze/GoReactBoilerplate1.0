export type UserData = {
  user: string,
  pess: string,
  mail: string,
  pass: string
}

const signupHandler = async function(user: UserData): Promise<boolean> {

  const res = await fetch('/signup', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(user)
  })

  if(res.status === 200){
    return true
  } else {
    return false
  }

}

export default signupHandler