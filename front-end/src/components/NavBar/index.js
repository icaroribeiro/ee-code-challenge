import React from 'react';
import { Navbar, NavbarBrand, Nav, NavItem, NavLink } from 'reactstrap';

import styles from './styles.module.css'; 

const NavBar = (props) => {
  const displayHomeLink = (auth) => {
    return (
      auth === true && (
      <Nav className="ml-auto" navbar>
        <NavItem>
          <NavLink href="/">home</NavLink>
        </NavItem>
      </Nav>
      )
    );
  };

  return (
    <div>
      <Navbar 
        className={styles.Navbar}
        color="white" 
        light 
        expand="md">
        <NavbarBrand
          className={styles.NavbarBrand}>
          githubstars
        </NavbarBrand>
        { displayHomeLink(props.isLogged) }
      </Navbar>
    </div>
  );
};

export default NavBar;