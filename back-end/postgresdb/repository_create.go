package postgresdb

import (
    "database/sql"
    "fmt"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/lib/pq"
)

const (
    QueryCreateUserRepository = "CreateUserRepository"
    QueryCreateRepository = "CreateRepository"
)

func AddCreateUserRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryCreateUserRepository] = `
            INSERT INTO 
            user_repositories (user_id, repository_id, tags)
            VALUES ($1, $2, $3);
        `
}

func AddCreateRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryCreateRepository] = `
            INSERT INTO 
            repositories (id, name, description, url, language)
            VALUES ($1, $2, $3, $4, $5);
        `
}

func (d *Datastore) CreateUserRepository(userRepository models.UserRepository) (models.UserRepository, error) {
    var result sql.Result
    var err error
    var nRowsAffected int64

    result, err = d.Stmts[QueryCreateUserRepository].Exec(userRepository.UserID, 
                userRepository.RepositoryID,
                pq.Array(userRepository.Tags))

    if err != nil {
        return userRepository, err
    }

    nRowsAffected, err = result.RowsAffected()

    if err != nil {
        return userRepository, err
    }

    if nRowsAffected == 0 {
        return userRepository, fmt.Errorf("it wasn't possible to generate the record")
    }

    return userRepository, nil
}

func (d *Datastore) CreateRepository(repository models.Repository) (models.Repository, error) {
    var result sql.Result
    var err error
    var nRowsAffected int64

    result, err = d.Stmts[QueryCreateRepository].Exec(repository.ID, 
                repository.Name,
                repository.Description,
                repository.URL,
                repository.Language)

    if err != nil {
        return repository, err
    }

    nRowsAffected, err = result.RowsAffected()

    if err != nil {
        return repository, err
    }

    if nRowsAffected == 0 {
        return repository, fmt.Errorf("it wasn't possible to generate the record")
    }

    return repository, nil
}
