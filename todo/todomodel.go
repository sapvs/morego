package todo

type Todoable interface {
	gethtml() string
}

type Todo struct {
	task     string
	subtodos []Todo
}

func (todo *Todo) gethtml() string {
	html := "<li>" + todo.task
	hastodos := len(todo.subtodos) != 0
	if hastodos {
		html += "<ul>"
	}

	for _, subtodo := range todo.subtodos {
		html += subtodo.gethtml()
	}

	if hastodos {
		html += "</ul>"
	}
	return html + "</li>"
}
