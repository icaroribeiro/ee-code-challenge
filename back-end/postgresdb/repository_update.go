package postgresdb

import (
    "database/sql"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/lib/pq"
)

const (
    QueryUpdateUserReposity = "UpdateUserRepository"
    QueryUpdateReposity = "UpdateRepository"
)

func AddUpdateUserRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryUpdateUserReposity] = `
            UPDATE user_repositories
            SET
            tags = $1
            WHERE user_id = $2 and repository_id = $3;
        `
}

func AddUpdateRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryUpdateReposity] = `
            UPDATE repositories
            SET
            name = $1,
            description = $2,
            url = $3,
            language = $4
            WHERE id = $5;
        `
}

func (d *Datastore) UpdateUserRepository(userId string, repositoryId string, userRepository models.UserRepository) (int64, error) {
    var result sql.Result
    var err error
    var nRowsAffected int64

    result, err = d.Stmts[QueryUpdateUserReposity].Exec(pq.Array(userRepository.Tags),
                                            userId, 
                                            repositoryId)

    if err != nil {
        return 0, err
    }

    nRowsAffected, err = result.RowsAffected()

    if err != nil {
        return 0, err
    }

    return nRowsAffected, nil
}

func (d *Datastore) UpdateRepository(id string, repository models.Repository) (int64, error) {
    var result sql.Result
    var err error
    var nRowsAffected int64

    result, err = d.Stmts[QueryUpdateReposity].Exec(repository.Name,
                                            repository.Description,
                                            repository.URL,
                                            repository.Language,
                                            id)

    if err != nil {
        return 0, err
    }

    nRowsAffected, err = result.RowsAffected()

    if err != nil {
        return 0, err
    }

    return nRowsAffected, nil
}
