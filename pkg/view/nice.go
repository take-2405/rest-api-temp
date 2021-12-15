package view

import (
	"backend-record/pkg/model/dto"
)

func ReturnNiceResopnse(nice *dto.Nice)dto.Nice{
	return dto.Nice{Nice: nice.Nice}
}
