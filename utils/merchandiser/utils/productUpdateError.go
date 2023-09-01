package merchandiser_utils

func ProductUpdateError(productName string, progress string) string {
	return progress + " ▬ Товар " + productName + " не удалось обновить"
}
