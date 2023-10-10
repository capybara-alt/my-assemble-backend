package job

import (
	"context"
)

type ICrawlJobUsecase interface {
	Execute(context.Context)
}
