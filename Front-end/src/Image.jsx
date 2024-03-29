import React, { useState } from 'react';
import axios from 'axios';

const apiKey = "";

function Image() {
  const [textInput, setTextInput] = useState('');
  const [imageSrc, setImageSrc] = useState('');
  const [loading, setLoading] = useState(false);

  const generateImage = async () => {
    setLoading(true);

    try {
      const response = await axios.post(`http://localhost:8080/generate?text=${encodeURIComponent(textInput)}`, {
        headers: {
          Authorization: `Bearer ${apiKey}`,
        },
      });

      setImageSrc(response.data.imageSrc);
    } catch (error) {
      console.error('Error generating image:', error);
    } finally {
      setLoading(false);
    }
  };

  const clearInput = () => {
    setTextInput('');
    setImageSrc('');
  };

  return (
    <div>
      <h1 className='head1'>Text to Image Generator</h1>
      <div className='ib'>
        <div className='ii'>
          <h1 className='head1'>Turn your Ideas into Picture</h1>
          <h3 className='head1'>Easily create an image from scratch with our AI Image <br /> generator by entering descriptive text </h3>
          <input className='in' type="text" value={textInput} onChange={(e) => setTextInput(e.target.value)} /> <br /> <br />
          <button className='ibut1' onClick={generateImage}>Generate Image</button> &nbsp;
          <button className='ibut2' onClick={clearInput}>Clear</button>
          {loading && <div className="loader"></div>}
        </div>
        <div className='ig'>
          <h2 className='head1'>Generated Image</h2>
          {imageSrc && <img className='img1' src={imageSrc} alt="Generated Image" />}
        </div>
      </div>
    </div>
  );
}

export default Image;
