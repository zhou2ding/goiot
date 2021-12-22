package models

import "bubble/dao"

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CrateTodo(todo *Todo) (err error) {
	return dao.DB.Create(todo).Error
}

func ListAllToDo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(todoList).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateToDO (todo *Todo) (err error) {
	return dao.DB.Model(&todo).Where(Todo{Id: todo.Id}).Update("status", todo.Status).Error
}

func DeleteToDo (todo *Todo) (err error) {
	return dao.DB.Where("id=?", todo.Id).Delete(Todo{}).Error
}
