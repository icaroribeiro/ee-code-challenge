import React, { Component } from 'react';
import { Redirect } from 'react-router-dom'
import { connect } from 'react-redux';
import { Container } from 'reactstrap';

import { getStatus } from '../../service/api';
import { updateRepository } from '../../service/api';
import { getAllUserGithubStarredRepositories } from '../../service/api.js';
import { createUserRepository } from '../../service/api.js';
import { getAllUserRepositories } from '../../service/api.js';
import { deleteUserRepository } from '../../service/api.js';

import styles from './styles.module.css'; 
import RepositoryTable from '../../pages/Repository';

import { grantLogin } from '../../actions/login';
import { retrieveRepositories } from '../../actions/repository';

// This function is to arrange all user repositories when comparing those currently starred on Github website
// with those that have already been registered in the database. The result of this evaluation will imply 
// the creation, editing and even removing of repositories from the database.
async function arrangeRepositories(username) {
  var response;

  var i;

  var githubStarredRepositories = [];
  var repositoryMap1 = new Map();
  
  var userRepositories = [];
  var repositoryMap2 = new Map();
  

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

  for (i = 0; i < githubStarredRepositories.length; i++) {
    repositoryMap1.set(githubStarredRepositories[i].id, githubStarredRepositories[i]);
  }

  response = await getAllUserRepositories(username);

  if (typeof response.data !== 'undefined' && response.data !== null) {
    userRepositories = response.data;
  }

  for (i = 0; i < userRepositories.length; i++) {
    repositoryMap2.set(userRepositories[i].id, userRepositories[i]);

    if (!repositoryMap1.has(userRepositories[i].id)) {
      response = await deleteUserRepository(username, userRepositories[i].id);
    }
  }

  for (var [key, value] of repositoryMap1) {
    if (repositoryMap2.has(key)) {
      var repository = repositoryMap2.get(key);

      if (value.name !== repository.name ||
        value.description !== repository.description ||
          value.url !== repository.url ||
            value.language !== repository.language) {
              response = await updateRepository(key, value);
            }
    } else {
      response = await createUserRepository(username, value);
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