package merchandiser_utils

import (
	"strconv"
)

func AssortmentRequestResult(assortmentRowsLen int, offsetCounter int) (isEmpty bool, message string) {
	if assortmentRowsLen == 0 {
		return true, "Товаров для обновления больше не осталось"
	}

	if assortmentRowsLen < 1000 && offsetCounter == 1 {
		return false, "Ассортимент успешно получен!\nВыполняется процесс обновления товаров..."
	}
	if assortmentRowsLen < 1000 && offsetCounter > 1 {
		return false, "Получена оставшаяся часть ассортимента!\nВыполняется процесс обновления товаров..."
	}

	return false, "Получена " + strconv.Itoa(offsetCounter) + "-ая тысяча товаров!\nВыполняется процесс обновления..."
}
