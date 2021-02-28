package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/recipe"

	"github.com/julienschmidt/httprouter"
)

type RecipeHandler struct {
	recipeService recipe.RecipeService
}

func NewRecipeHandler(rspService recipe.RecipeService) *RecipeHandler {
	fmt.Println("recipe handler created")
	return &RecipeHandler{recipeService: rspService}
}

func (rh *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	recipes, errs := rh.recipeService.Recipes()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipes, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rh *RecipeHandler) GetIngredients(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, errr := strconv.Atoi(ps.ByName("id"))

	if errr != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	ingredients, errs := rh.recipeService.Ingredients(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(ingredients, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rh *RecipeHandler) GetUserRecipes(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	uid, err := strconv.Atoi(ps.ByName("uid"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipes, errs := rh.recipeService.UserRecipes(uint(uid))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipes, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (rh *RecipeHandler) PostRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("recipe handelr")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	recipe := &entity.Recipe{}
	fmt.Println("in post recipe 2")
	err := json.Unmarshal(body, recipe)
	// for i := 0; i < recipe..length; i++ {
	// 	fmt.Println("gooooooooooooooooooooooooo")
	// }
	fmt.Println(recipe)

	if err != nil {
		fmt.Println("HEEEEEEEEEEE222EEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.StoreRecipe(recipe)
	fmt.Println("my recipeee")
	fmt.Println(recipe)
	if len(errs) > 0 {
		//fmt.Println("HEEEEEEEEEEEEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(output)
	p := fmt.Sprintf("/recipes/add/%d", recipe.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return
}

//post image of recipe
func (rh *RecipeHandler) PostImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("image here")
	r.ParseForm()
	fmt.Println("steppp1")
	file, handler, err := r.FormFile("file")
	fmt.Println("steppp2")
	rid := ps.ByName("id")
	fmt.Println(rid)
	fmt.Println(file != nil)
	fmt.Println(handler.Filename)
	if err != nil {
		fmt.Println("HEEEEEEEEEEE222EEEEEE")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, _ := strconv.Atoi(rid)
	if err != nil {
		fmt.Println("conversion")
	}
	recipet, errs := rh.recipeService.Recipe(uint(id))
	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(recipet)
	dst, err := os.Create(filepath.Join("./images/", filepath.Base(rid+""+handler.Filename)))
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		fmt.Println("erorrrrrrrrrrrrrrrrr")
		return
	}
	// recipe, errs := rh.recipeService.updateImage(recipet, dst)

	// if len(errs) > 0 {
	// 	//fmt.Println("HEEEEEEEEEEEEEEEE")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }

	// output, err := json.MarshalIndent(recipe, "", "\t\t")

	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }

	// p := fmt.Sprintf("/recipes/add/%d", recipe.ID)
	// w.Header().Set("Location", p)
	// w.WriteHeader(http.StatusCreated)
	// w.Write(output)
	// return
	return
} // GetSinglerecipe handles
func (rh *RecipeHandler) GetSingleRecipe(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.Recipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (rh *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rh.recipeService.DeleteRecipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

func (rh *RecipeHandler) PutRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipe, errs := rh.recipeService.Recipe(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &recipe)
	recipe, errs = rh.recipeService.UpdateRecipe(recipe)
	fmt.Println(recipe)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(recipe, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
