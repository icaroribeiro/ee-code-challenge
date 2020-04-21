import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import { Container } from 'reactstrap';

import styles from './styles.module.css'; 
import CredentialsForm from "./CredentialsForm";

import { RequestLogin } from '../../actions/login.js';

class Home extends Component {
    constructor(props) {
        super(props);

        this.state = {
          username: ""
        };
    }

  validateForm() {
    return this.state.username.length > 0;
  }

  handleChange = event => {
    this.setState({
      [ event.target.name ]: event.target.value
    });
  }

  handleSubmit = event => {
    event.preventDefault();

    if (this.state.username.trim()) {
      this.props.history.push({
        pathname: '/repositories',
        state: {
          username: this.state.username
        }
      });
    }
  }

  componentDidMount() {
    this.props.onRequestLogin();
  }

  render() {
    return (
      <Container className={styles.Container}>
        <CredentialsForm
          handleChange={ this.handleChange.bind(this) }
          handleSubmit={ this.handleSubmit.bind(this) }
          validateForm={ this.validateForm.bind(this) }
          username={ this.state.username }
        />
      </Container>
    );
  }
}

const mapDispatchToProps = dispatch => {
  return {
    onRequestLogin: () => {
      dispatch(RequestLogin());
    }
  };
};

export default withRouter(connect(null, mapDispatchToProps)(Home));