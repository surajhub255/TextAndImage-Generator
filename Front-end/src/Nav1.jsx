import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Container from '@mui/material/Container';
import img1 from './assets/lazarus.jpeg'


function Nav1() {
  return (
    <AppBar position="static">
      <Container maxWidth="xl" className='nav1'>
        <Toolbar disableGutters>
          <img src={img1} alt="img1"  sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} width='80px' height='50px'/>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default Nav1;