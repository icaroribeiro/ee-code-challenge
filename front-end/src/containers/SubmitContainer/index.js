import React, { Component } from 'react';
import { Redirect } from 'react-router-dom'
import { connect } from 'react-redux';
import { Container } from 'reactstrap';

import { getStatus } from '../../service/api';
import { getGithubStarredRepositories } from '../../service/api';
import { getRepositories } from '../../service/api';
import { createRepository } from '../../service/api';
import { updateRepository } from '../../service/api';
import { deleteRepository } from '../../service/api';

import styles from './styles.module.css'; 
import RepositoryTable from '../../pages/Repository';

import { grantLogin } from '../../actions/login';
import { retrieveRepositories } from '../../actions/repository';

//
// This function is intended for dealing with retrieveing all data from repositories
// before implying the creation, editing or even removal of repositories from the database.
//
async function arrangeRepositories(username) {
  var response;

  var i, j;
  var isRegistered;
  
  var githubStarredRepositories = []
  var userRepositories = [];
  var commonRepositories = [];

  var commonRepositoriesToCreate = [];
  var commonRepositoriesToDelete = [];

  try {
    response = await getStatus();
  } catch (err) {
    console.log('Caught error: ', err);
    return;
  }

  response = await getGithubStarredRepositories(username);
  githubStarredRepositories = response.data;

  response = await getRepositories(username);
  userRepositories = response.data;

  for (i = 0; i < githubStarredRepositories.length; i++) {
    isRegistered = false;

    for (j = 0; j < userRepositories.length; j++) {
      if (githubStarredRepositories[i].id === userRepositories[j].id) {
        isRegistered = true;
        break;
      }
    }

    if (!isRegistered) {
      response = await createRepository(username, githubStarredRepositories[i]);
      commonRepositoriesToCreate.push(githubStarredRepositories[i]);
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
      response = await deleteRepository(username, userRepositories[i].id);
      commonRepositoriesToDelete.push(userRepositories[i]);
    }
  }

  response = await getRepositories(undefined);
  commonRepositories = response.data;

  if (commonRepositoriesToCreate.length > 0) {
    for (i = 0; i < commonRepositoriesToCreate.length; i++) {
      isRegistered = false;

      for (j = 0; j < commonRepositories.length; j++) {
        if (commonRepositoriesToCreate[i].id === commonRepositories[j].id) {
          isRegistered = true;
          break;
        }
      }

      if (!isRegistered) {
        response = await createRepository(undefined, commonRepositoriesToCreate[i]);
      } else {
        response = await updateRepository(undefined, commonRepositoriesToCreate[i].id, commonRepositoriesToCreate[i]);
      }
    }
  }

  if (commonRepositoriesToDelete.length > 0) {
    for (i = 0; i < commonRepositoriesToDelete.length; i++) {
      isRegistered = false;

      for (j = 0; j < commonRepositories.length; j++) {
        if (commonRepositoriesToDelete[i].id === commonRepositories[j].id) {
          isRegistered = true;
          break;
        }
      }

      if (!isRegistered) {
        response = await deleteRepository(undefined, commonRepositoriesToDelete[i].id);
      }
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
      dispatch(GrantLogin())
    },
    onRetrieveRepositories: async (username) => {
      try {
        const response = await getRepositories(username);
        dispatch(retrieveRepositories(response.data));
      }
      catch (err) {
        console.log('Caught error: ', err);
      }
    }
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(SubmitContainer);