package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	merchandiser_goods "merchandiser/goods"
	merchandiser_moysklad "merchandiser/moysklad"
	merchandiser_utils "merchandiser/utils"

	"github.com/joho/godotenv"
)

func init() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	err := godotenv.Load(filepath.Join(exPath, ".env"))
	if err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	token := os.Getenv("MOYSKLAD_AUTH_TOKEN")

	offset := 0
	offsetCounter := 1
	for {
		fmt.Println("Выполняется запрос ассортимента МойСклад...")
		assortmentResponse, err := merchandiser_moysklad.Assortment(token, offset)
		if err != nil {
			log.Fatal("Не удалось получить ассортимент")
		}

		assortmentIsEmpty, assortmentRequestResultMessage := merchandiser_utils.AssortmentRequestResult(len(assortmentResponse.Rows), offsetCounter)
		fmt.Println(assortmentRequestResultMessage)
		if assortmentIsEmpty {
			break
		}

		for i, assortmentRow := range assortmentResponse.Rows {
			assortmentRowParams := strings.Split(assortmentRow.Name, "-{}-")
			if len(assortmentRowParams) == 3 {
				continue
			}

			searchPayload := merchandiser_goods.SearchPayload{
				SearchText: assortmentRowParams[0],
			}
			searchResponse, err := merchandiser_goods.Search(searchPayload)
			if err != nil {
				fmt.Println(merchandiser_utils.ProductUpdateError(assortmentRow.Name, merchandiser_utils.Progress(i+1, len(assortmentResponse.Rows))))
			}

			if len(searchResponse.Items) > 0 {
				productCardPayload := merchandiser_goods.ProductCardPayload{
					GoodsId: searchResponse.Items[0].Goods.GoodsId,
				}
				productCardResponse, err := merchandiser_goods.ProductCard(productCardPayload)
				if err != nil {
					fmt.Println(merchandiser_utils.ProductUpdateError(assortmentRow.Name, merchandiser_utils.Progress(i+1, len(assortmentResponse.Rows))))
				}

				productUpdatePayload := merchandiser_moysklad.ProductUpdatePayload{
					Name:        productCardResponse.Item.Goods.Title + "-{}-" + productCardResponse.Item.CollectionName + "-{}-updated",
					Description: merchandiser_utils.ProductDescription(productCardResponse.Item.Description, productCardResponse.Item.Goods.Attributes),
					Images: []merchandiser_moysklad.ProductUpdatePayloadImage{{
						FileName: productCardResponse.Item.Goods.ImageLink,
						Content:  merchandiser_utils.Base64ByUrl(productCardResponse.Item.Goods.ImageLink),
					}},
				}
				productUpdateResponse, err := merchandiser_moysklad.ProductUpdate(token, assortmentRow.Id, productUpdatePayload)
				if err != nil {
					fmt.Println(merchandiser_utils.ProductUpdateError(assortmentRow.Name, merchandiser_utils.Progress(i+1, len(assortmentResponse.Rows))))
				}

				if productUpdateResponse == true {
					fmt.Println(merchandiser_utils.Progress(i+1, len(assortmentResponse.Rows)) + " ✓ Товар " + assortmentRow.Name + " успешно обновлен!")
				}
			} else {
				fmt.Println(merchandiser_utils.ProductUpdateError(assortmentRow.Name, merchandiser_utils.Progress(i+1, len(assortmentResponse.Rows))))
			}
		}

		offset = offset + 1000
		offsetCounter++
	}

	fmt.Println("Обновление ассортимента завершено!")
}
