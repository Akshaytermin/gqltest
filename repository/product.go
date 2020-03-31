package repository

import (
	"context"
	"log"

	"github.com/Akshaytermin/gqltest/graph/model"
)

func Create(ctx context.Context, doc model.Product) (model.Product, error) {
	var resp model.Product

	db := model.FetchConnection()
	defer db.Close()

	err := db.Create(&doc).Preload("Ingredients").Find(&resp).Error
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	return resp, nil
}

func Update(ctx context.Context, query model.Product, updates model.Product) (model.Product, error) {
	var product model.Product

	db := model.FetchConnection()
	defer db.Close()

	err := db.Model(model.Product{}).Where(query).First(&product).Updates(updates).Preload("Ingredients").Error
	if err != nil {
		log.Fatal(err)
		return product, err
	}

	return product, nil
}

func UpdateIngredients(ctx context.Context, query model.Ingredient, updates model.Ingredient) (model.Ingredient, error) {

	var resp model.Ingredient

	db := model.FetchConnection()
	defer db.Close()

	err := db.Model(model.Ingredient{}).Where(query).First(&resp).Updates(updates).Error
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}

func Delete(ctx context.Context, id *int) (model.Product, error) {
	db := model.FetchConnection()
	defer db.Close()

	var product model.Product
	//Fetch based on ID and delete
	err := db.Where("id = ?", *id).First(&product).Delete(&product).Error
	if err != nil {
		log.Fatal(err)
		return product, err
	}
	var products []*model.Product
	db.Preload("Ingredients").
		Find(&products)

	return product, nil
}

func DeleteIngregients(ctx context.Context, id []int) error {
	db := model.FetchConnection()
	defer db.Close()
	var ing model.Ingredient
	err := db.Where("id IN (?)", id).First(&ing).Delete(ing).Error
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func FindAll(ctx context.Context) ([]*model.Product, error) {

	db := model.FetchConnection()

	defer db.Close()

	var products []*model.Product

	err := db.Preload("Ingredients").Find(&products).Error
	if err != nil {
		log.Fatal(err)
		return products, err
	}

	return products, err
}

func FindByID(ctx context.Context, id *int) ([]model.Ingredient, error) {

	db := model.FetchConnection()

	defer db.Close()

	var resp []model.Ingredient
	err := db.Find(&resp).Where("product_id = ?").Error
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	return resp, nil
}
