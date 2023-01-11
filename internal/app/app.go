package app

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Run() error {
	return nil //Заглушка! После реализации метода Run будет возвращаться ошибка
}
