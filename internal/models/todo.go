package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Volteal/todo-app/internal/common"
	"github.com/alexeyco/simpletable"
)

type item struct {
	TaskName    string
	TaskNote    string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []item

func (t *TodoList) Show() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Task #"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Note"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Creates At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	var cells [][]*simpletable.Cell
	for i, item := range *t {
		idx := i + 1
		taskName := common.Blue(item.TaskName)
		taskNote := common.Blue(item.TaskNote)
		taskDone := common.Red("\u274C Incomplete")
		if item.Done {
			taskName = common.Green(item.TaskName)
			taskNote = common.Green(item.TaskNote)
			taskDone = common.Green("\u2705 Complete")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", idx)},
			{Text: taskName},
			{Text: taskNote},
			{Text: taskDone},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 6, Text: common.Red(fmt.Sprintf("You have: [%d] tasks pending.", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *TodoList) Add(name, note string) {
	todo := item{
		TaskName:    name,
		TaskNote:    note,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *TodoList) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	trueIndex := index - 1

	ls[trueIndex].CompletedAt = time.Now()
	ls[trueIndex].Done = true
	fmt.Printf("Task: %v. Has been marked as complete!\n", ls[trueIndex].TaskName)

	return nil
}

func (t *TodoList) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	trueIndex := index - 1

	*t = append(ls[:trueIndex], ls[index:]...)

	return nil
}

func (t *TodoList) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoList) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoList) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
