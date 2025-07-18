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
          type='text'
          placeholder='请输入用户名'
          onChange={(e) => setUsername(e.target.value)}
        />
      </div>
     </div>
    </>
  )
}

export default App
