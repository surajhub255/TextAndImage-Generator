import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import Image from './Image'
import Text from './Text'
import {BrowserRouter as Router , Routes , Route} from 'react-router-dom'
import Landing from './Landing'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>

  <Router>
    <Routes>
      <Route path='/' element={<Landing/>}/>
      <Route path='/Text' element={<Text/>}/>
      <Route path='/Image' element={<Image/>}/>
    
    </Routes>
  </Router>


  </React.StrictMode>,
)
