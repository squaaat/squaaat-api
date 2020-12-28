package model

func Load() []interface{} {
	return []interface{}{
		&User{},
		&UserSessionToken{},
	}
}
