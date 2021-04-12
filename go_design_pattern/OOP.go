package go_design_pattern

type UserBo struct{
	ID string
	Name string
	Age string
}

func (u *UserBo)GetName()string{
	return u.Name
}
