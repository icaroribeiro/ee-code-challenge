import React, { Component } from 'react';
import { Redirect } from 'react-router-dom'
import { connect } from 'react-redux';
import { Container } from 'reactstrap';

import { getStatus } from '../../service/api';
import { getAllUserGithubStarredRepositories } from '../../service/api.js';
import { createUserRepository } from '../../service/api.js';
import { getAllUserRepositories } from '../../service/api.js';
import { deleteUserRepository } from '../../service/api.js';

import styles from './styles.module.css'; 
import RepositoryTable from '../../pages/Repository';

import { grantLogin } from '../../actions/login';
import { retrieveRepositories } from '../../actions/repository';

// This function is intended for dealing with retrieveing all data from repositories
// before implying the creation, editing or even removal of repositories from the database.
async function arrangeRepositories(username) {
    var response;

    var i, j;
    var isRegistered;
    
    var githubStarredRepositories = []
    var userRepositories = [];

    try {
        response = await getStatus();
    } catch (err) {
        console.log('Caught error: ', err);
        return;
    }

    response = await getAllUserGithubStarredRepositories(username);
    
    if (typeof response.data !== 'undefined' && response.data !== null) {
        githubStarredRepositories = response.data;
    }

    response = await getAllUserRepositories(username);

    if (typeof response.data !== 'undefined' && response.data !== null) {
        userRepositories = response.data;
    }

    for (i = 0; i < githubStarredRepositories.length; i++) {
        isRegistered = false;

        for (j = 0; j < userRepositories.length; j++) {
            if (githubStarredRepositories[i].id === userRepositories[j].id) {
                isRegistered = true;
                break;
            }
        }

        if (!isRegistered) {
            response = await createUserRepository(username, githubStarredRepositories[i]);
        }
    }

    for (i = 0; i < userRepositories.length; i++) {
        isRegistered = false;

        for (j = 0; j < githubStarredRepositories.length; j++) {
            if (userRepositories[i].id === githubStarredRepositories[j].id) {
                isRegistered = true;
                break;
            }
        }

        if (!isRegistered) {
            response = await deleteUserRepository(username, userRepositories[i].id);
        }
    }
}

class SubmitContainer extends Component {
  constructor(props) {
    super(props);

    this.state = {
      loading: true,
      progress: 100
    }
  }  

  componentDidMount() {
    var username = this.props.location.state.username;

    if (typeof(this.props.location.state) !== 'undefined') {
      arrangeRepositories(username);      

      setTimeout(function() {
        this.props.onGrantLogin();
        this.props.onRetrieveRepositories(username)
        .then(() => {
          this.setState({ loading: false })
        })
        .catch((err) => {
          console.log('Caught error: ', err);
        })
      }.bind(this), 3000);
    }
  }

  render() {
    if (typeof(this.props.location.state) === 'undefined') {
      return <Redirect to='/' />
    }

    const Loading = this.props.placeholder

    if (this.state.loading) {
      return (
        <Container className={styles.Container}>
          <Loading 
            progress={ this.state.progress }
          />
        </Container>
      );
    }

    var repositories = this.props.repositories;
    var username = this.props.location.state.username;

    if (repositories === "") {
      return <RepositoryTable data={ [] } />
    } else {
      return <RepositoryTable username = { username }
        data={ repositories } />
    }
  }
}

const mapStateToProps = (state) => {
  return {
    repositories: state.repositoryReducer.repositories
  };
}

const mapDispatchToProps = dispatch => {
  return {
    onGrantLogin: () => {
      dispatch(grantLogin())
    },
    onRetrieveRepositories: async (username) => {
      try {
        const response = await getAllUserRepositories(username);
        dispatch(retrieveRepositories(response.data));
      }
      catch (err) {
        console.log('Caught error: ', err);
      }
    }
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(SubmitContainer);