import React, { useState } from 'react';

const Text = () => {
  const [inputMessage, setInputMessage] = useState('');
  const [generatedText, setGeneratedText] = useState('');
  const [copySuccess, setCopySuccess] = useState(false);

  const handleGenerateText = async () => {
    try {
      const response = await fetch('http://localhost:8080/generate-text', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: inputMessage }),
      });

      const data = await response.json();
      setGeneratedText(data.result);
    } catch (error) {
      console.error('Error generating text:', error);
    }
  };

  const clearInput = () => {
    setInputMessage('');
    setGeneratedText('');
    setCopySuccess(false); // Reset copy success state
  };

  const handleCopyText = () => {
    const textArea = document.createElement('textarea');
    textArea.value = generatedText;
    document.body.appendChild(textArea);
    textArea.select();
    document.execCommand('copy');
    document.body.removeChild(textArea);
    
    setCopySuccess(true); 
    setTimeout(() => {
      setCopySuccess(false);
    }, 2000); 
  };

  return (
    <div>
      <h1 className='th'>Generate Text</h1>
      <div className='tb'>
        <div className='ts'>
          <h2>Text</h2>
          <textarea
            className='ti'
            type="text"
            id="inputMessage"
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
          />
          <br />
          <div className='tbutb'>
            <button type='reset' onClick={clearInput} className='tbut1'><b>Clear</b></button>
            <button className='tbut2' onClick={handleGenerateText}><b>Generate Text</b></button>
          </div>
        </div>
        <div className='tgb'>
          <h2>Generated Text:</h2>
          <textarea className='tp' cols="30" rows="10" value={generatedText}></textarea>
          <div>
            <button className='tbut3' onClick={handleCopyText}><b>Copy</b></button>
            {copySuccess && <p className="copy-notification">Copied to clipboard!</p>}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Text;
