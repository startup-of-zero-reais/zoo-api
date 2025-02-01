package helpers

import "fmt"

func GetAge(text string) (string, error) {
	ages := map[string]string{
		"Neonato": "neonate",
		"Filhote": "cub",
		"Jovem":   "young",
		"Adulto":  "adult",
		"Senil":   "senile",
	}

	age, ok := ages[text]
	if !ok {
		return "", fmt.Errorf("a idade é inválida")
	}

	return age, nil
}

func GetGender(text string) (string, error) {
	genders := map[string]string{
		"Feminino":   "female",
		"Masculino":  "male",
		"Indefinido": "undefined",
	}

	gender, ok := genders[text]
	if !ok {
		return "", fmt.Errorf("o gênero é inválido")
	}

	return gender, nil
}
