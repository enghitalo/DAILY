package service

import (
	util "daily_project_module/src/utils"
	"fmt"
	"strconv"
)

func AuthVerify(token string) error {
	token, err := strconv.Unquote(token)
	if err != nil {
		return err
	}

	if user, err := util.DecodeJwtToken(token); err != nil || len(user) == 0 {
		return fmt.Errorf("token inválido %v", err)
	}
	return nil
}
