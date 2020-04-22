import axios from 'axios';

const defaultOptions = {
  baseURL: `http://${process.env.REACT_APP_API_HOST}:8080`,
  headers: {
    'Content-Type': 'application/json',
  },
}

const api = axios.create(defaultOptions);

export const getStatus = () => { 
  try {
    return api.get(`/status`)
  } catch (err) {
    console.log('Caught error: ', err);
  }
}

export const getRepository = (id) => {
  try {
    return api.get(`/repositories/${id}`)
  } catch (err) {
    console.log('Caught error: ', err);
  }
}

export const getAllUserGithubStarredRepositories = (username) => { 
  try {
    return api.get(`/users/${username}/githubStarredRepositories`)
  } catch (err) {
    console.log('Caught error: ', err);
  }
}

export const createUserRepository = (username, repository) => { 
  try {
    return api.post(`/users/${username}/repository`, 
        JSON.stringify(repository)
      )
  } catch (err) {
    console.log('Caught error: ', err);
  }
}

export const getAllUserRepositories = (username) => { 
  try {
    return api.get(`/users/${username}/repositories`)
  } catch (err) {
    console.log('Caught error: ', err);
  }
}

export const updateUserRepository = (username, id, repository) => { 
  try {
    return api.put(`/users/${username}/repositories/${id}`, 
        JSON.stringify(repository)
      )
  } catch (err) {
    console.log('Caught error: ', err);
  }
}

export const deleteUserRepository = (username, id) => { 
  try {
    return api.delete(`/users/${username}/repositories/${id}`)
  } catch (err) {
    console.log('Caught error: ', err);
  }
}