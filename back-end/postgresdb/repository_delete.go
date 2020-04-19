package postgresdb

import (
    "database/sql"
)

const (
    QueryDeleteUserRepository = "DeleteUserRepository"
    QueryDeleteRepository = "DeleteRepository"
)

func AddDeleteUserRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryDeleteUserRepository] = `
            DELETE
            FROM user_repositories
            WHERE user_id = $1 and repository_id = $2;
        `
}

func AddDeleteRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryDeleteRepository] = `
            DELETE 
            FROM repositories
            WHERE id = $1;
        `
}

func (d *Datastore) DeleteUserRepository(userId string, repositoryId string) (int64, error) {
    var err error
    var nRowsAffected int64
    var result sql.Result

    result, err = d.Stmts[QueryDeleteUserRepository].Exec(userId, repositoryId)

    if err != nil {
        return 0, err
    }

    nRowsAffected, err = result.RowsAffected()

    if err != nil {
        return 0, err
    }

    return nRowsAffected, nil
}

func (d *Datastore) DeleteRepository(id string) (int64, error) {
    var err error
    var nRowsAffected int64
    var result sql.Result

    result, err = d.Stmts[QueryDeleteRepository].Exec(id)

    if err != nil {
        return 0, err
    }

    nRowsAffected, err = result.RowsAffected()

    if err != nil {
        return 0, err
    }

    return nRowsAffected, nil
}
