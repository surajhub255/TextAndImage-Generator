import React from 'react';
import { useNavigate } from 'react-router-dom';
import Nav1 from './Nav1';
import LPaper from './LPaper';


const Landing = () => {
  const navigate = useNavigate();

  return (
    <div>
      <Nav1 />
      <h1 className='lh'>
        <b>
          What if AI can <span style={{ color: 'blue' }}>create any <br />Text or Image</span> in Seconds?
        </b>
      </h1>
      <h3 className='lh'><b><span style={{ color: 'grey' }}>Meet LAZARUS, the AI Image and Text generetor who turns whatever you can imagine into unique <br /> image and photos in seconds.Finnaly the perfect picture to match any message.</span> </b></h3>
      <div className='lbut'>
        <button className='lbut1' onClick={() => navigate("/Text")}>
          <b>Text Generator</b>
        </button>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <button className='lbut1' onClick={() => navigate("/Image")}>
         <b> Image Generator</b>
        </button>
      </div>
      <br /><br /><br /><br /><br /><br />
      <LPaper />
    </div>
  );
};

export default Landing;
