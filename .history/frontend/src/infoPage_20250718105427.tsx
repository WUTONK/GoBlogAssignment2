import { useState } from 'react'
import './App.css'
import { Api } from './shared'

function InfoPage() {
  const [nickname, setNickname] = useState("")

  Api.userInfoGet({
    infoReq:{authorization: ""}
  })

  return (
    <>
      <p><strong>HI:</strong> {nickname}</p>
    </>
  )
}

export default InfoPage
