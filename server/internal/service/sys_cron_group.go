// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/sysin"
)

type ISysCronGroup interface {
	Delete(ctx context.Context, in sysin.CronGroupDeleteInp) error
	Edit(ctx context.Context, in sysin.CronGroupEditInp) (err error)
	Status(ctx context.Context, in sysin.CronGroupStatusInp) (err error)
	MaxSort(ctx context.Context, in sysin.CronGroupMaxSortInp) (*sysin.CronGroupMaxSortModel, error)
	View(ctx context.Context, in sysin.CronGroupViewInp) (res *sysin.CronGroupViewModel, err error)
	List(ctx context.Context, in sysin.CronGroupListInp) (list []*sysin.CronGroupListModel, totalCount int, err error)
	Select(ctx context.Context, in sysin.CronGroupSelectInp) (list sysin.CronGroupSelectModel, err error)
}

var localSysCronGroup ISysCronGroup

func SysCronGroup() ISysCronGroup {
	if localSysCronGroup == nil {
		panic("implement not found for interface ISysCronGroup, forgot register?")
	}
	return localSysCronGroup
}

func RegisterSysCronGroup(i ISysCronGroup) {
	localSysCronGroup = i
}