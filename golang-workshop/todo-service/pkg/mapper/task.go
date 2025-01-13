package mapper

import "todo-service/pkg/model"

func MapToDto(task model.Task) model.MongoTask {
	return model.MongoTask{
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
}

func MapToModel(task model.MongoTask) model.Task {
	return model.Task{
		ID:          task.ID.Hex(),
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
}
