import { useState } from 'react'
import './App.css'

function App() {
  const [password, setPassword] = useState("")
  const [username, setUsername] = useState("")

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
            placeholder="请输入用户名"
            onChange={(e) => setPassword(e.target.value)}
            className="password-input"
          />
      </div>

     </div>
    </>
  )
}

export default App
