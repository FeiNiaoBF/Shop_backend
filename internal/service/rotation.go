// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goBack/internal/model"
)

type (
	IRotation interface {
		Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error)
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.RotationUpdateInput) error
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}
