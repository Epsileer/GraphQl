package repos

import "graphql/models"

var ironMan = models.Hero{
	ID:              1,
	Name:            "Iron Man",
	Friends:         []int{2, 3},
	Characteristics: "Genius billionaire, philanthropist",
	Status:          "ACTIVE",
}

var captainAmerica = models.Hero{
	ID:              2,
	Name:            "Captain America",
	Friends:         []int{1, 3},
	Characteristics: "Super soldier with enhanced strength and agility",
	Status:          "ACTIVE",
}

var thor = models.Hero{
	ID:              3,
	Name:            "Thor",
	Friends:         []int{1, 2},
	Characteristics: "God of Thunder with a magical hammer (Mjolnir)",
	Status:          "ACTIVE",
}

func GetHeroes() []*models.Hero {
	heroes := []*models.Hero{&ironMan, &captainAmerica, &thor}
	return heroes
}

func GetHeroById(id int) *models.Hero {
	heroes := GetHeroes()
	for _, hero := range heroes {
		if hero.ID == id {
			return hero
		}
	}
	return nil
}

func GetHeroByStatus(status string) *models.Hero {
	heroes := GetHeroes()
	for _, hero := range heroes {
		if hero.Status == status {
			return hero
		}
	}
	return nil
}

func UpdateHeroNameById(id int, name string) *models.Hero {
	hero := GetHeroById(id)
	if hero != nil {
		hero.Name = name
	}
	return hero
}
