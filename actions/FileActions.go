package actions

import (
	"encoding/json"
	"io/ioutil"
	"tuido/components/todolist"
	"tuido/data"
)

func Save(path string, list todolist.TodoList) {
  s := data.ToSerialized(list)
  data, err := json.MarshalIndent(s, "", "  ")
  if err != nil {
    panic(err)
  }
  err = ioutil.WriteFile(path, data, 0644)
  if err != nil {
    panic(err)
  }
}

func Load(path string) todolist.TodoList {
  content, err := ioutil.ReadFile(path)
  if err != nil {
    return todolist.NewTodoList()
  }
  var list data.SerializedData
  err = json.Unmarshal(content, &list)
  if err != nil {
    return todolist.NewTodoList()
  }
  return data.ToMainModel(list)
}
