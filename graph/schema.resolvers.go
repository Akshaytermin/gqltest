package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/Akshaytermin/gqltest/graph/generated"
	model "github.com/Akshaytermin/gqltest/graph/model"
	"github.com/Akshaytermin/gqltest/repository"
)

//create the product
func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct, ingredients []*model.NewIngredient) (*model.Product, error) {
	var resp model.Product

	product := model.Product{
		Name:  input.Name,
		Price: *input.Price,
	}

	product.Ingredients = make([]model.Ingredient, len(ingredients))

	for i, item := range ingredients {
		product.Ingredients[i] = model.Ingredient{Name: item.Name}
	}

	//Create by passing the pointer to the product
	resp, err := repository.Create(ctx, product)
	if err != nil {
		log.Fatal(err)
		return &resp, err
	}

	return &resp, nil
}

//updating the product
func (r *mutationResolver) UpdateProduct(ctx context.Context, id *int, input *model.NewProduct, ingredients []*model.NewIngredient) (*model.Product, error) {

	var product model.Product

	db := model.FetchConnection()
	defer db.Close()

	product.Ingredients = make([]model.Ingredient, len(ingredients))
	for index, item := range ingredients {
		product.Ingredients[index] = model.Ingredient{Name: item.Name}
	}

	query := model.Product{}

	query.ID = *id

	updates := model.Product{
		Name:        input.Name,
		Price:       *input.Price,
		Ingredients: product.Ingredients,
	}

	product, err := repository.Update(ctx, query, updates)
	if err != nil {
		log.Fatal(err)
		return &product, err
	}

	return &product, nil
}

//deleting the product
func (r *mutationResolver) DeleteProduct(ctx context.Context, id *int) ([]*model.Product, error) {

	var resp []*model.Product

	_, err := repository.Delete(ctx, id)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	return resp, nil
}

//Finding all product details
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {

	var products []*model.Product

	products, err := repository.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
		return products, err
	}

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
