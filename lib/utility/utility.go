package utility

import (
	"github.com/eifzed/antre-app/internal/entity/repo/antre"
)

func StringExistInSlice(item string, itemSlice []string) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

func RoleExistInSlice(item antre.MstRole, itemSlice []antre.MstRole) bool {
	for _, i := range itemSlice {
		if i.Name == item.Name {
			return true
		}
	}
	return false
}

func IntExistInSlice(item int, itemSlice []int) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

func Int32ExistInSlice(item int32, itemSlice []int32) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}

func Int64ExistInSlice(item int64, itemSlice []int64) bool {
	for _, i := range itemSlice {
		if i == item {
			return true
		}
	}
	return false
}
