package db

import (
	"bytes"
	"context"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"

	apiv1alpha1 "github.com/percona/percona-server-mysql-operator/api/v1alpha1"
	"github.com/percona/percona-server-mysql-operator/pkg/clientcmd"
)

type SwitchoverSQLManager struct {
	db *db
}

func NewSwitchoverSQLManager(pod *corev1.Pod, cliCmd clientcmd.Client, user apiv1alpha1.SystemUser, pass, host string) *SwitchoverSQLManager {
	return &SwitchoverSQLManager{db: newDB(pod, cliCmd, user, pass, host)}
}

func (m *SwitchoverSQLManager) ExecSQL(ctx context.Context, query string) error {
	var errb, outb bytes.Buffer

	err := m.db.exec(ctx, query, &outb, &errb)
	if err != nil {
		return errors.Wrap(err, "ExecSQL failed")
	}

	return nil
}
