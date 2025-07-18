import { useState } from 'react'
import './App.css'
import { Api } from './shared'

function InfoPage() {
  const [nickname, setNickname] = useState("")

  Api.userInfoGet({
    infoReq:{authorization: localStorage.getItem("token") || ""}
  }).then(
    (res)=>{
      setNickname(res.nikeName)
    }
  ).catch((err) => {alert(err.message)})

  return (
    <>      
      <p><strong>HI:</strong> {nickname}</p>
    </>
  )
}

export default InfoPage
