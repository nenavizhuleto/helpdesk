package issue

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Issue struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	Title       string `form:"title" db:"title"`
	Description string `form:"description" db:"description"`
	Status      string `db:"status"`
}

type Table struct {
	db                                         *sqlx.DB
	insert, retrieve, update, listByUser, list *sqlx.NamedStmt
}

func NewTable(db *sqlx.DB) (*Table, error) {
	insert, err := db.PrepareNamed("INSERT INTO issues VALUES (:id, :user_id, :title, :description, :status);")
	if err != nil {
		return nil, err
	}

	retrieve, err := db.PrepareNamed("SELECT * FROM issues WHERE id = :id AND user_id = :user_id;")
	if err != nil {
		return nil, err
	}

	update, err := db.PrepareNamed("UPDATE issues SET title = :title, description = :description, status = :status WHERE id = :id AND user_id = :user_id;")
	if err != nil {
		return nil, err
	}

	listByUser, err := db.PrepareNamed("SELECT * FROM issues WHERE user_id = :user_id;")
	if err != nil {
		return nil, err
	}
	list, err := db.PrepareNamed("SELECT * FROM issues;")
	if err != nil {
		return nil, err
	}

	return &Table{
		db:         db,
		insert:     insert,
		retrieve:   retrieve,
		update:     update,
		listByUser: listByUser,
		list:       list,
	}, nil
}

func New(userId string) *Issue {
	return &Issue{
		ID:     uuid.NewString(),
		UserID: userId,
	}
}

func (it *Table) Insert(issue Issue) error {
	_, err := it.insert.Exec(&issue)
	return err
}

func (it *Table) Update(issue Issue) (int, error) {
	res, err := it.update.Exec(&issue)
	if err != nil {
		return 0, err
	}
	rows_affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rows_affected), nil
}

func (it *Table) ByUser(user_id string) ([]Issue, error) {
	data := []Issue{}
	err := it.listByUser.Select(&data, &Issue{UserID: user_id})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (it *Table) All() ([]Issue, error) {
	data := []Issue{}
	err := it.list.Select(&data, Issue{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (it *Table) Get(id string) (Issue, error) {
	i := Issue{}
	err := it.retrieve.Get(&i, &Issue{ID: id})
	if err != nil {
		return Issue{}, err
	}

	return i, err
}
