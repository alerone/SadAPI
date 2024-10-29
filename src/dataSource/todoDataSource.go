package dataSource

import (
	"errors"
	"fmt"
	"mysadapi/models"
	"time"
)

func GetToDos() ([]models.ToDo, error) {
	rows, err := db.Query("SELECT * FROM toDo")
	if err != nil {
		return nil, errors.New("error trying to execute the query")
	}
	defer rows.Close()

	var todos []models.ToDo
	for rows.Next() {
		var todo models.ToDo
		var CreatedAt time.Time
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &CreatedAt)
		if err != nil {
			return nil, errors.New("error trying to scan the results")
		}
		todo.CreatedAt = CreatedAt.Format("Mon, 02 Jan 2006 15:04:05 MST")
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("error at rows")
	}
	return todos, nil
}

func GetToDosWhere(condition string, args ...interface{}) ([]models.ToDo, error) {
	// Preparamos la consulta SQL con la cláusula WHERE dinámica
	query := "SELECT * FROM toDo WHERE " + condition + " $1"

	// Ejecutamos la consulta con los argumentos proporcionados
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute the query: %v", err)
	}
	defer rows.Close()

	var todos []models.ToDo
	for rows.Next() {
		var todo models.ToDo
		var CreatedAt time.Time
		// Escaneamos los resultados
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &CreatedAt)
		if err != nil {
			return nil, errors.New("error trying to scan the results")
		}
		todo.CreatedAt = CreatedAt.Format("Mon, 02 Jan 2006 15:04:05 MST")
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("error at rows")
	}

	// Retornamos las tareas encontradas
	return todos, nil
}

func CreateToDo(title, description string, completed bool) (models.ToDo, error) {
	todo := models.ToDo{ID: 0, Title: title, Description: description, Completed: completed, CreatedAt: ""}
	var createdAt time.Time
	query := "INSERT INTO toDo (title, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := db.QueryRow(query, title, description, completed).Scan(&todo.ID, &createdAt)
	if err != nil {
		return todo, errors.New("error trying to create the toDo")
	}
	todo.CreatedAt = createdAt.Format("Mon, 02 Jan 2006 15:04:05 MST")

	return todo, nil
}

func UpdateToDo(id int, field string, value interface{}) error {
	// Crea la consulta SQL basada en el campo especificado
	query := fmt.Sprintf("UPDATE toDo SET %s=$1 WHERE id=$2", field)

	// Ejecuta la consulta
	_, err := db.Exec(query, value, id)
	if err != nil {
		return fmt.Errorf("error trying to update: %v", err)
	}

	return nil
}

func DeleteToDo(id int) error {
	query := "DELETE FROM toDo WHERE id=$1"
	_, err := db.Exec(query, id)

	if err != nil {
		return errors.New("error trying to delete the toDo")
	}

	return nil
}
