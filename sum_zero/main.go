package main

func main() {
	entrada := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}

	sumZero(entrada)

}

func sumZero(entrada []int) []int {

	counter := 0
	var valuePositivo int
	var valueNegativo int
	result := []int{}
	if len(entrada) == 0 {
		return result
	}

	for i := counter; i < len(entrada); i = counter {
		if i == len(entrada)-1 {
			if entrada[i] > 0 && valuePositivo > 0 {
				result = append(result, valuePositivo)
				result = append(result, entrada[i])
				break
			} else if entrada[i] > 0 && valueNegativo < 0 {
				result = append(result, (entrada[i] + valueNegativo))
			} else if entrada[i] < 0 && valuePositivo > 0 {
				result = append(result, (valuePositivo + entrada[i]))
			}
		}
		if entrada[i] >= 0 {

			valuePositivo = entrada[i] + valuePositivo
		} else {
			valueNegativo = entrada[i] + valueNegativo

		}
		for j := i + 1; j < len(entrada); j++ {
			if entrada[j] >= 0 {
				valuePositivo = entrada[j] + valuePositivo

			} else {
				valueNegativo = entrada[j] + valueNegativo
			}

			if valuePositivo+valueNegativo == 0 {
				valueNegativo = 0
				valuePositivo = 0
				counter = j + 1
				break
			} else if valuePositivo+valueNegativo < 0 {
				counter = j + 1
				break
			} else {
				valuePositivo = valuePositivo + valueNegativo
				valueNegativo = 0
				counter = j + 1
				break
			}
		}
	}
	return result
}
