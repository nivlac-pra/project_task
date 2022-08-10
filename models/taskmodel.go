package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/task-go/config"
	"github.com/golang/task-go/entities"
)

type TaskModel struct {
	conn *sql.DB
}

func NewTaskModel() *TaskModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &TaskModel{
		conn: conn,
	}
}

func (t *TaskModel) FindAll() ([]entities.Task, error) {

	rows, err := t.conn.Query("select * from task")
	if err != nil {
		return []entities.Task{}, err
	}
	defer rows.Close()

	var tasklist []entities.Task
	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id, &task.Task, &task.Assigne, &task.Deadline)

		deadline, _ := time.Parse("2006-01-02", task.Deadline)
		task.Deadline = deadline.Format("02-01-2006")

		tasklist = append(tasklist, task)
	}

	return tasklist, nil
}

func (t *TaskModel) Create(task entities.Task) bool {

	result, err := t.conn.Exec("insert into task (task, assigne, deadline) values(?, ?, ?) ",
		task.Task, task.Assigne, task.Deadline)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (t *TaskModel) Find(id int64, task *entities.Task) error {

	return t.conn.QueryRow("select * from task where id = ?", id).Scan(&task.Id, &task.Task, &task.Assigne, &task.Deadline)

}

func (t *TaskModel) Update(task entities.Task) error {

	_, err := t.conn.Exec(
		"update task set task = ?, assigne = ?, deadline = ? where id = ?",
	 	task.Task, task.Assigne, task.Deadline, task.Id)

	if err != nil {
		return err
	}

	return nil

}
