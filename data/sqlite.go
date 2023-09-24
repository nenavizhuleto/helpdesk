package data

import (
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const DBNAME = "app.db"

type App struct {
    Db *sql.DB
    insert, update, retrieve, listClient, list *sql.Stmt
    Subscribers Subscriber
}

var DB *App

func NewDB() (*App, error) {
    db, err := sql.Open("sqlite3", DBNAME)

    if err != nil {
        return nil, err
    }

    insert, err := db.Prepare("INSERT INTO issues VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);")
    if err != nil {
        return nil, err
    }

    retrieve, err := db.Prepare("SELECT * FROM issues WHERE id = ? AND client = ?;")
    if err != nil {
        return nil, err
    }

    update, err := db.Prepare("UPDATE issues SET company = ?, department = ?, name = ?, phonenumber = ?, innernumber = ?, description = ?, status = ? WHERE id = ? AND client = ?;")
    if err != nil {
        return nil, err
    }

    listClient, err := db.Prepare("SELECT * FROM issues WHERE client = ?;")
    if err != nil {
        return nil, err
    }
    list, err := db.Prepare("SELECT * FROM issues;")
    if err != nil {
        return nil, err
    }

    app := &App{
        Db: db,
        insert: insert,
        retrieve: retrieve,
        update: update,
        listClient: listClient,
        list: list,
        Subscribers: *NewSubscriber(),
    }

    return app, nil
}

func (a *App) InsertIssue(issue Issue) error {
    _, err := a.insert.Exec(
        issue.ID,
        issue.ClientID,
        issue.Company,
        issue.Department,
        issue.Name,
        issue.PhoneNumber,
        issue.InnerNumber,
        issue.Description,
        issue.Status,
    )
    a.Subscribers.Notify(issue.ClientID)
    return err
}

func (a *App) UpdateIssue(issue Issue) (int, error) {
    res, err := a.update.Exec(
        issue.Company,
        issue.Department,
        issue.Name,
        issue.PhoneNumber,
        issue.InnerNumber,
        issue.Description,
        issue.Status,
        issue.ID,
        issue.ClientID,
    )
	if err != nil {
		return 0, err
	}
	rows_affected, err := res.RowsAffected()
    if err != nil {
		return 0, err
	}

    a.Subscribers.Notify(issue.ClientID)
	return int(rows_affected), nil
}

func (a *App) ListClientIssues(client string) ([]Issue, error) {
	rows, err := a.listClient.Query(client)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

    data := []Issue{}
	for rows.Next() {
		i := Issue{}
		err = rows.Scan(
            &i.ID, 
            &i.ClientID, 
            &i.Company, 
            &i.Department,
            &i.Name,
            &i.PhoneNumber,
            &i.InnerNumber,
            &i.Description,
            &i.Status,
        )
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}

func (a *App) ListIssues() ([]Issue, error) {
	rows, err := a.list.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

    data := []Issue{}
	for rows.Next() {
		i := Issue{}
		err = rows.Scan(
            &i.ID, 
            &i.ClientID, 
            &i.Company, 
            &i.Department,
            &i.Name,
            &i.PhoneNumber,
            &i.InnerNumber,
            &i.Description,
            &i.Status,
        )
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}

func (a *App) RetrieveIssue(id string, client string) (Issue, error) {
	row := a.retrieve.QueryRow(id, client)

	i := Issue{}
	var err error
    if err = row.Scan(
        &i.ID, 
        &i.ClientID, 
        &i.Company, 
        &i.Department,
        &i.Name,
        &i.PhoneNumber,
        &i.InnerNumber,
        &i.Description,
        &i.Status,
    ); err == sql.ErrNoRows {
        return Issue{}, err
    }

	return i, err
}
