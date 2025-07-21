import { useState } from 'react'
import './app.css'
import { Button, Input, Space } from 'antd'
import { Api } from './shared'

function LoginPage() {
  const [password, setPassword] = useState("")
  const [username, setUsername] = useState("")
  const [token, setToken] = useState("")

  return (
    <>
     <div className='app-container'>
     <Space direction="vertical" size="large" style={{ innerHeight: '100%' }}>
      <Space direction="vertical" size="small" style={{ innerHeight: '100%' }}>
      <div className='input-password'>
        <Input
          type="text"
          value={username}
          placeholder="请输入用户名"
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setUsername(e.target.value)}
        />
      </div>
      <div>
        <Input
            type="text"
            value={password}
            placeholder="请输入密码"
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)}
            className="password-input"
          />
      </div>
      </Space>
      
      {/* 登录按钮 */}
      <Button
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
      >
        登录
      </Button>
      </Space>
     </div>
    </>
  )
}

export default LoginPage
