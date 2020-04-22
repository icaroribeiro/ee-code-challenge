import React, { Component } from "react";
import { connect } from 'react-redux';

import NavBar from "../NavBar";
import Footer from "../Footer";

class LayoutNav extends Component{
  render(){
    let renderData;
    
    renderData = (
      this.props.children
    );

    const isLogged = this.props.isLogged;

    return(
      <div>
        <div>
          <NavBar
            isLogged={ isLogged }
          />
        </div>
        <div>
          {renderData}
        </div>
        <div>
          <Footer
          />
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    isLogged: state.loginReducer.isLogged
  };
}

export default connect(mapStateToProps)(LayoutNav);