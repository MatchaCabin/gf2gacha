package util

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"
)

func BackupDB() error {
	dbData, err := os.ReadFile("gf2gacha.db")
	if err != nil {
		return errors.WithStack(err)
	}

	backupName := fmt.Sprintf("gf2gacha_%d.db", time.Now().UnixNano())
	err = os.WriteFile(backupName, dbData, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
