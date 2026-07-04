package cli

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"task-cli/internal/storage"
	"task-cli/internal/task"
)

type Handler struct {
	Storage *storage.Storage
}

func NewHandler(s *storage.Storage) *Handler {
	return &Handler{Storage: s}
}

func (h *Handler) getNextID(tasks []task.Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func (h *Handler) Add(description string) error {
	tasks, err := h.Storage.Load()
	if err != nil {
		return err
	}

	newTask := task.Task{
		ID:          h.getNextID(tasks),
		Description: description,
		Status:      task.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	if err := h.Storage.Save(tasks); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return nil
}

func (h *Handler) Update(id int, description string) error {
	tasks, err := h.Storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	if err := h.Storage.Save(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d updated successfully.\n", id)
	return nil
}

func (h *Handler) UpdateStatus(id int, status string) error {
	tasks, err := h.Storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	if err := h.Storage.Save(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as %s.\n", id, status)
	return nil
}

func (h *Handler) Delete(id int) error {
	tasks, err := h.Storage.Load()
	if err != nil {
		return err
	}

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := h.Storage.Save(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully.\n", id)
	return nil
}

func (h *Handler) List(statusFilter string) error {
	tasks, err := h.Storage.Load()
	if err != nil {
		return err
	}

	var filtered []task.Task
	for _, t := range tasks {
		if statusFilter == "" || t.Status == statusFilter {
			filtered = append(filtered, t)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDescription\tStatus\tCreated At\tUpdated At")
	for _, t := range filtered {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			t.ID,
			t.Description,
			t.Status,
			t.CreatedAt.Format("2006-01-02 15:04:05"),
			t.UpdatedAt.Format("2006-01-02 15:04:05"),
		)
	}
	w.Flush()
	return nil
}
