package control

import (
	"GoProject/model"
)

type MemberData struct {
	data []model.Member
}

func NewMember() *MemberData {
	return &MemberData{}
}

// GetAllData    godoc
// @Summary      Get All Data
// @Description  get string data
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Member
// @Router       /getAllData [get]
func (m *MemberData) GetAllData() []model.Member {
	return m.data
}

// PostCreateData    godoc
// @Summary      Create Data
// @Description  Create  data
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Member
// @Router       /postCreateData [post]
func (m *MemberData) PostCreateData(data model.Member) model.Member {
	m.data = append(m.data, data)
	return data
}
