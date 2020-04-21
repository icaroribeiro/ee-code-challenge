import { RETRIEVE_REPOSITORIES } from './types.js';
import { RENEW_REPOSITORY } from './types.js';

export const retrieveRepositories = (repositories) => {
  return {
    type: RETRIEVE_REPOSITORIES,
    payload: { 
      repositories: repositories
    }
  }
};

export const renewRepository = (repository) => {
  return {
    type: RENEW_REPOSITORY,
    payload: { 
      repository: repository
    }
  }
};