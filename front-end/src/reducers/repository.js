import { RETRIEVE_REPOSITORIES } from '../actions/types.js';
import { RENEW_REPOSITORY } from '../actions/types.js';

const initialState = {
  repositories: []
}

export default function repositoryReducer(state = initialState, action) {
  switch (action.type) {
    case RETRIEVE_REPOSITORIES:
      return {
        ...state,
        repositories: action.payload.repositories
      };
    case RENEW_REPOSITORY:
      return {
        ...state,
        repositories: state.repositories.map(repository => {
          if (repository.id === action.payload.repository.id) {
            return action.payload.repository;
          }
          return repository;
        })
      };
    default:
      return state;
  }
}