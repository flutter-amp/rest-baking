package repository

import (
	"fmt"

	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/recipe"
	"github.com/jinzhu/gorm"
)

type RecipeGormRepo struct {
	conn *gorm.DB
}

func NewRecipeGormRepo(db *gorm.DB) recipe.RecipeRepository {
	return &RecipeGormRepo{conn: db}
}

func (recipeRepo *RecipeGormRepo) Recipes() ([]entity.Recipe, []error) {
	recipes := []entity.Recipe{}
	errs := recipeRepo.conn.Find(&recipes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return recipes, errs
}

func (recipeRepo *RecipeGormRepo) Recipe(id uint) (*entity.Recipe, []error) {
	recipe := entity.Recipe{}
	errs := recipeRepo.conn.First(&recipe, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &recipe, errs
}

func (recipeRepo *RecipeGormRepo) DeleteRecipe(id uint) (*entity.Recipe, []error) {
	rcpe, errs := recipeRepo.Recipe(id)

	if len(errs) > 0 {
		return nil, errs
	}
	errs = recipeRepo.conn.Delete(rcpe, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return rcpe, errs
}
func (recipeRepo *RecipeGormRepo) UpdateRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	rsp := recipe
	errs := recipeRepo.conn.Save(rsp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return rsp, errs
}

func (recipeRepo *RecipeGormRepo) StoreRecipe(recipe *entity.Recipe) (*entity.Recipe, []error) {
	rcpe := recipe
	errs := recipeRepo.conn.Create(rcpe).GetErrors()
	if len(errs) > 0 {
		fmt.Println("found errors........................")
		return nil, errs
	}
	return rcpe, errs
}

// func (recipeRepo *RecipeGormRepo) updateImage(recipe *entity.Recipe, imagePath string) (*entity.Recipe, []error) {

// 	rcpe := recipe
// 	errs := recipeRepo.conn.Model(&rcpe).UpdateColumn("imageUrl", imagePath).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return rcpe, errs
// }

func (recipeRepo *RecipeGormRepo) UserRecipes(uid uint) ([]entity.Recipe, []error) {
	usrRecipes := []entity.Recipe{}
	errs := recipeRepo.conn.Where("user_id = ?", uid).Find(&usrRecipes).GetErrors()
	//errs := recipeRepo.conn.Model(user).Related(&usrRecipes, "Orders").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrRecipes, errs
}
