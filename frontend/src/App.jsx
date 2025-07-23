import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [message, setMessage] = useState("");
  useEffect(() => {
    fetch("http://localhost:8080/api/hello")
    .then((res) => res.json())
    .then((data) => setMessage(data.message))
    .catch((err) => setMessage("Error: " + err))

  }, [])

  return <div>{message}</div>
}

export default App
