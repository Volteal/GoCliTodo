# GoCliTodo

**GoCliTodo** is a simple CLI todo app written in Go. running the -add command creates a todo-list.json in the current folder allowing you to create files based on the project you are working on.

## Installation

```
make install
```

## Available Commands

```
todo -list
  List all list items!

todo -add <Task Name, Task Note>
  Add a new list item!

todo -complete=X
  Mark a todo item as complete where X is the item's number.

todo -del=X
  Delete an item from the list where X is the item's number.
```
