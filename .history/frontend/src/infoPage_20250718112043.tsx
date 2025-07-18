import { useState } from 'react'
import './App.css'
import { Api } from './shared'

function InfoPage() {
  const [nickname, setNickname] = useState("")
  const [token, setToken] = useState("")

  Api.userInfoGet({
    infoReq:{authorization: ""}
  }).then(
    (res)=>{
      setNickname(res.nikeName)
    }
  )

  return (
    <>
      <input
          type="text"
          value={token}
          onChange={(e) => setToken(e.target.value)}
          placeholder="请输入城token"
          className="token-input"
        />
      <p><strong>HI:</strong> {nickname}</p>
    </>
  )
}

export default InfoPage
