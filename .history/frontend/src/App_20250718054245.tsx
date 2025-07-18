import { useState } from 'react'
import './App.css'

function App() {
  const [password, setPassword] = useState(0)
  const [username, setUsername] = useState(0)

  return (
    <>
     <div className='app-container'>
      <div className='input-password'>
        <input
          type='text'
        />
      </div>
     </div>
    </>
  )
}

export default App
