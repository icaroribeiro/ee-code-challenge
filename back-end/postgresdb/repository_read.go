package postgresdb

import (
    "database/sql"
    "database/sql/driver"
    "fmt"
    "github.com/icaroribeiro/ee-code-challenge/back-end/models"
    "github.com/lib/pq"
)

const (
    QueryGetUserRepository = "GetUserRepository"
    QueryGetAllUserRepositoriesByUserId = "GetAllUserRepositoriesByUserId"
    QueryGetAllUserRepositoriesByRepositoryId = "GetAllUserRepositoriesByRepositoryId"
    QueryGetRepository = "GetRepository"
)

func AddGetAllUserRepositoriesStatements(unpreparedStmts map[string]string) {
    var stmts map[string]string
    var key string
    var value string

    stmts = map[string]string{
        QueryGetUserRepository: `
            SELECT user_id, repository_id, tags
            FROM user_repositories
            WHERE user_id = $1 and repository_id = $2;
        `,

        QueryGetAllUserRepositoriesByUserId: `
            SELECT user_id, repository_id, tags
            FROM user_repositories
            WHERE user_id = $1;
        `,

        QueryGetAllUserRepositoriesByRepositoryId: `
            SELECT user_id, repository_id, tags
            FROM user_repositories
            WHERE repository_id = $1;
        `,
    }

    for key, value = range stmts {
        unpreparedStmts[key] = value
    }
}

func AddGetRepositoryStatement(unpreparedStmts map[string]string) {
    unpreparedStmts[QueryGetRepository] = `
            SELECT id, name, description, url, language
            FROM repositories
            WHERE id = $1;
        `
}

func (d *Datastore) GetUserRepository(userId string, repositoryId string) (models.UserRepository, error) {
    var err error
    var userRepository models.UserRepository
    var dataArray []sql.NullString
    var data sql.NullString
    var isOK bool
    var tag string
    var value driver.Value

    err = d.Stmts[QueryGetUserRepository].QueryRow(userId, repositoryId).Scan(&userRepository.UserID, 
                                                &userRepository.RepositoryID,
                                                pq.Array(&dataArray))
    
    // Complete the list of all tags.
    for _, data = range dataArray {
        value, err = data.Value()

        tag, isOK = value.(string)

        if !isOK {
            return userRepository, fmt.Errorf("it wasn't possible to get data from the list of all tags")
        }

        userRepository.Tags = append(userRepository.Tags, tag)
    }
    
    if err != nil {
        if err != sql.ErrNoRows {
            return userRepository, err
        }
    }

    return userRepository, nil
}

func (d *Datastore) GetAllUserRepositoriesByUserId(userId string) ([]models.UserRepository, error) {
    var rows *sql.Rows
    var err error
    var userRepository models.UserRepository
    var dataArray []sql.NullString
    var data sql.NullString
    var driverValue driver.Value
    var tag string
    var isOK bool
    var userRepositories []models.UserRepository
    
    rows, err = d.Stmts[QueryGetAllUserRepositoriesByUserId].Query(userId)

    defer rows.Close()

    for rows.Next() {
        err = rows.Scan(&userRepository.UserID, 
                    &userRepository.RepositoryID, 
                    pq.Array(&dataArray))

        // Complete the list of the tags.
        for _, data = range dataArray {
            driverValue, err = data.Value()
    
            tag, isOK = driverValue.(string)
    
            if !isOK {
                return userRepositories, 
                    fmt.Errorf("it wasn't possible to get data from the list of all tags")
            }
    
            userRepository.Tags = append(userRepository.Tags, tag)
        }

        if err != nil {
            return userRepositories, err
        }

        userRepositories = append(userRepositories, userRepository)
    }

    return userRepositories, nil
}

func (d *Datastore) GetAllUserRepositoriesByRepositoryId(repositoryId string) ([]models.UserRepository, error) {
    var rows *sql.Rows
    var err error
    var userRepository models.UserRepository
    var dataArray []sql.NullString
    var data sql.NullString
    var driverValue driver.Value
    var tag string
    var isOK bool
    var userRepositories []models.UserRepository
    
    rows, err = d.Stmts[QueryGetAllUserRepositoriesByRepositoryId].Query(repositoryId)

    defer rows.Close()

    for rows.Next() {
        err = rows.Scan(&userRepository.UserID, 
                    &userRepository.RepositoryID, 
                    pq.Array(&dataArray))

        // Complete the list of the tags.
        for _, data = range dataArray {
            driverValue, err = data.Value()
    
            tag, isOK = driverValue.(string)
    
            if !isOK {
                return userRepositories, 
                    fmt.Errorf("it wasn't possible to get data from the list of all tags")
            }
    
            userRepository.Tags = append(userRepository.Tags, tag)
        }

        if err != nil {
            return userRepositories, err
        }

        userRepositories = append(userRepositories, userRepository)
    }

    return userRepositories, nil
}

func (d *Datastore) GetRepository(id string) (models.Repository, error) {
    var err error
    var repository models.Repository

    err = d.Stmts[QueryGetRepository].QueryRow(id).Scan(&repository.ID, 
                                                &repository.Name, 
                                                &repository.Description,
                                                &repository.URL, 
                                                &repository.Language)
    
    if err != nil {
        if err != sql.ErrNoRows {
            return repository, err
        }
    }

    return repository, nil
}
