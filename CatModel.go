package main

type Cat struct {
	tableName struct{} `pg:"cats"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsStrip   bool     `json:"is_strip" pg:"is_strip"`
	Color     string   `json:"color" pg:"color"`
}

//FindAllCats Получить всех котиков из таблицы.

func FindAllCats() []Cat {
	var cats []Cat
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cats).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cats
}

// Добавить котика в таблицу.
func CreateCat(cat Cat) Cat {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&cat).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cat
}
