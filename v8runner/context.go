package v8runner

type Контекст struct {
	КлючСоединенияСБазой  string
	Пользователь          string
	Пароль                string
	КлючРазрешенияЗапуска string
	КодЯзыка              string
	КодЯзыкаСеанса        string
}

func НовыйКонтекст() *Контекст {

	return newContext()

}

func newContext() *Контекст {
	return &Контекст{
		"",
		"",
		"",
		"",
		"",
		"",
	}
}
