package backup

import (
	"os"

	"time"

	"fmt"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/pkg/errors"
)

type BackupDirectoryManager struct{}

func (BackupDirectoryManager) Create(path, name string, logger orchestrator.Logger, nowFunc func() time.Time) (orchestrator.Backup, error) {
	var (
		backupPath string
		err        error
	)

	directoryName := name + "_" + nowFunc().UTC().Format("20060102T150405Z")

	if path != "" {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		if !fileInfo.IsDir() {
			return nil, fmt.Errorf("artifact path %s is not a directory", path)
		}

		backupPath = fmt.Sprintf("%s/%s", path, directoryName)
	} else {
		backupPath = directoryName
	}

	err = os.Mkdir(backupPath, 0700)
	if err != nil {
		return nil, errors.New("failed creating artifact directory")
	}

	return &BackupDirectory{baseDirName: backupPath, Logger: logger}, nil
}

func (BackupDirectoryManager) Open(name string, logger orchestrator.Logger) (orchestrator.Backup, error) {
	_, err := os.Stat(name)
	return &BackupDirectory{baseDirName: name, Logger: logger}, errors.Wrap(err, "failed opening the directory")
}
