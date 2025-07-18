import { useState } from 'react'
import './app.css'
import { Api } from './shared'

function InfoPage() {
  const [nickname, setNickname] = useState("")

  Api.userInfoGet(
    {}, // 或 { infoReq: { authorization: "" } }，看后端需不需要
    {
      headers: {
        Authorization: localStorage.getItem("token") || ""
      }
    }
  ).then(
    (res) => setNickname(res.nikeName)
  ).catch((err) => alert(err.message))

  // Api.userInfoGet({
  //   // 读取本地缓存 token，如果没有的话就返回空值
  //   // infoReq:{authorization: localStorage.getItem("token") || ""}
    
  // }).then(
  //   (res)=>{
  //     setNickname(res.nikeName)
  //   }
  // ).catch((err) => {alert(err.message)})

  return (
    <>      
      <p><strong>HI:</strong> {nickname}</p>
    </>
  )
}

export default InfoPage
