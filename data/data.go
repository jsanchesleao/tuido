package data

import (
	"tuido/components/todolist"
)

type SerializedData struct {
  Items []SerializedItem
}

type SerializedItem struct {
  Text string
  Done bool
}

func ToSerialized(list todolist.TodoList) SerializedData {
  serialized := SerializedData{ Items: []SerializedItem{} }
  for _, item := range list.Items {
    serialized.Items = append(serialized.Items, SerializedItem{
      Text: item.Text,
      Done: item.Done,
    })
  }
  return serialized
}


func ToMainModel(s SerializedData) todolist.TodoList {
  list := todolist.NewTodoList()
  for _, item := range s.Items {
    list.Items = append(list.Items, todolist.TodoItem{
      Text: item.Text,
      Done: item.Done,
    })
  }
  return list
}

