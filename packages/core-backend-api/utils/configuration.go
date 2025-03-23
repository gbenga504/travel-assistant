package utils

import (
	"fmt"
	"os"

	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
)

func LookupEnv(name string) string {
	env, ok := os.LookupEnv(name)

	errMsg := fmt.Sprintf("Failed to load %s", name)

	if !ok {
		logger.Error(errMsg, logger.ErrorOpt{
			Name:    errors.Name(errors.ErrEnvNotLoaded),
			Message: errors.Message(errors.ErrEnvNotLoaded),
		})
	}

	return env
}
