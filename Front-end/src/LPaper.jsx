import * as React from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import img1 from './assets/car.jpg';
import img2 from './assets/fox.png';
import img3 from './assets/sunset.jpg';
import img4 from './assets/game.jpg';


export default function LPaper() {
  return (
    <Box
      sx={{
        display: 'flex',
        flexWrap: 'wrap',
        '& > :not(style)': {
          m: 1,
          width: 355,
          height: 200,
        },
      }}
    >
      <Paper elevation={4}>
        <img src={img1} alt="Lazarus" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
      </Paper>
      <Paper elevation={4}>
        <img src={img2} alt="Lazarus" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
      </Paper>
      <Paper elevation={4}>
        <img src={img3} alt="Lazarus" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
      </Paper>
      <Paper elevation={4}>
        <img src={img4} alt="Lazarus" style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
      </Paper>
      
    </Box>
  );
}
