package models

type Status struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

func DummyStatuses() []Status {
	return []Status{
		{Id: 1, Name: "В пути"},
		{Id: 2, Name: "На складе"},
		{Id: 3, Name: "Продан"},
		{Id: 4, Name: "Снят с продажи"},
	}
}
