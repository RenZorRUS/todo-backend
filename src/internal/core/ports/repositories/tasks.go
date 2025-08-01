package repositories

import (
	"context"
	"time"

	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/models"
)

type (
	BaseFindTaskParams struct {
		UserID      *uint64
		Title       *string
		Description *string
		DueDate     *time.Time
		Status      *models.TaskStatus
	}

	FindTaskParams struct {
		BaseFindTaskParams

		ID          *uint64
		ShowDeleted bool
	}

	FindTasksParams struct {
		BaseFindTaskParams

		Limit       *uint64
		Offset      *uint64
		ShowDeleted bool
	}

	DeleteTaskParams struct {
		ID     uint64
		UserID uint64
	}

	TaskRepository interface {
		GetTask(ctx context.Context, params *FindTaskParams) (*models.Task, error)
		GetTasks(ctx context.Context, params *FindTasksParams) ([]models.Task, error)
		CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
		UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
		SoftDeleteTask(ctx context.Context, params *DeleteTaskParams) error
		HardDeleteTask(ctx context.Context, params *DeleteTaskParams) error
	}
)
