package model

func (gb GetterBackends) FindUserByID(v int) (map[string]UserGroup, error) {
	r := map[string]UserGroup{}
	for _, b := range gb {
		lr, err := b.FindUserByID(v)
		if err != nil {
			return nil, err
		}
		r = mergeUserGroup(r, lr)
	}
	return r, nil
}
func (gb GetterBackends) FindUserByName(v string) (map[string]UserGroup, error) {
	r := map[string]UserGroup{}
	for _, b := range gb {
		lr, err := b.FindUserByName(v)
		if err != nil {
			return nil, err
		}
		r = mergeUserGroup(r, lr)
	}
	return r, nil
}
func (gb GetterBackends) FindGroupByID(v int) (map[string]UserGroup, error) {
	r := map[string]UserGroup{}
	for _, b := range gb {
		lr, err := b.FindGroupByID(v)
		if err != nil {
			return nil, err
		}
		r = mergeUserGroup(r, lr)
	}
	return r, nil
}
func (gb GetterBackends) FindGroupByName(v string) (map[string]UserGroup, error) {
	r := map[string]UserGroup{}
	for _, b := range gb {
		lr, err := b.FindGroupByName(v)
		if err != nil {
			return nil, err
		}
		r = mergeUserGroup(r, lr)
	}
	return r, nil
}
func (gb GetterBackends) Users() (map[string]UserGroup, error) {
	r := map[string]UserGroup{}
	for _, b := range gb {
		lr, err := b.Users()
		if err != nil {
			return nil, err
		}
		r = mergeUserGroup(r, lr)
	}
	return r, nil
}
func (gb GetterBackends) Groups() (map[string]UserGroup, error) {
	r := map[string]UserGroup{}
	for _, b := range gb {
		lr, err := b.Groups()
		if err != nil {
			return nil, err
		}
		r = mergeUserGroup(r, lr)
	}
	return r, nil
}
func (gb GetterBackends) HighestUserID() int {
	r := 0
	for _, b := range gb {
		lr := b.HighestUserID()
		if lr != 0 {
			r = lr
		}
	}
	return r
}
func (gb GetterBackends) LowestUserID() int {
	r := 0
	for _, b := range gb {
		lr := b.LowestUserID()
		if lr != 0 {
			r = lr
		}
	}
	return r
}
func (gb GetterBackends) HighestGroupID() int {
	r := 0
	for _, b := range gb {
		lr := b.HighestGroupID()
		if lr != 0 {
			r = lr
		}
	}
	return r
}
func (gb GetterBackends) LowestGroupID() int {
	r := 0
	for _, b := range gb {
		lr := b.LowestGroupID()
		if lr != 0 {
			r = lr
		}
	}
	return r
}
