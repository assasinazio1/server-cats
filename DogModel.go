package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	EarLength int16    `json:"ear_length" pg:"ear_length"`
	Color     string   `json:"color" pg:"color"`
}

func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}

func CreateDog(cat Dog) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func FindDogById(dogId string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dog).Where("id = ?", dogId).First()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func DropDogById(dogId string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	dog = FindDogById(dogId)

	_, err := pgConnect.Model(&dog).Where("id = ?", dogId).Delete()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name
	oldDog.EarLength = dog.EarLength
	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Set("name = ?", oldDog.Name).
		Set("is_strip = ?", oldDog.EarLength).
		Set("color = ?", oldDog.Color).
		Where("id = ?", oldDog.ID).
		Update()

	if err != nil {
		panic(err)
	}
	pgConnect.Close()
	return oldDog
}
