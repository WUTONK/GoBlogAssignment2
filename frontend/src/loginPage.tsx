import { useState } from 'react'
import './app.css'
import { Api } from './shared'

function LoginPage() {
  const [password, setPassword] = useState("")
  const [username, setUsername] = useState("")
  const [token, setToken] = useState("")

  return (
    <>
     <div className='app-container'>
      <div className='input-password'>
        <input
          type="text"
          value={username}
          placeholder="请输入用户名"
          onChange={(e) => setUsername(e.target.value)}
          className="username-input"
        />
      </div>
      <div>
        <input
            type="text"
            value={password}
            placeholder="请输入密码"
            onChange={(e) => setPassword(e.target.value)}
            className="password-input"
          />
      </div>
      <button
      className='login-button'
      title="登陆"
      onClick={()=>{
        Api.userLoginPost(
          {
            loginReq:{
              username,
              password
            } 
          }
        ).then((res)=>{
          setToken(res.token)
          // 将 token 缓存到本地
          localStorage.setItem("token",token)
          alert(token)
        }).catch((err) => {alert(err.message)})
      }}
      />
     </div>
    </>
  )
}

export default LoginPage
