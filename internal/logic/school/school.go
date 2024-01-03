package school

import (
	"context"
	"gfdemo/internal/dao"
	"gfdemo/internal/service"
)

type sSchool struct{}

func init() {
	service.RegisterSchool(&sSchool{})
}

func (s *sSchool) GetOne(ctx context.Context) (data interface{}, err error) {

	var (
		m = dao.School.Ctx(ctx)
	)
	one, err := m.WherePri(1).One()
	data = one
	return
}
