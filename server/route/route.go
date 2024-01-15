package route

import (
	"api/form-api/db"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Slug      string             `bson:"slug"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type CreateFormDto struct {
	Name string `json:"name"`
}

// Create form
func CreateForm(ctx *gin.Context) {
	var createFormData CreateFormDto
	if err := ctx.BindJSON(&createFormData); err != nil {
		return
	}

	newForm := Form{
		ID:        primitive.NewObjectID(),
		Name:      createFormData.Name,
		Slug:      strings.Replace(createFormData.Name, " ", "-", -1),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := db.InitDb().Collection(db.Collections.Form).InsertOne(context.TODO(), newForm)

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusCreated, result)
}

// get all forms
func GetAllForms(ctx *gin.Context) {
	curs, err := db.InitDb().Collection(db.Collections.Form).Find(context.TODO(), bson.D{})

	if err != nil {
		return
	}

	var results []Form
	curs.All(context.TODO(), &results)

	for _, result := range results {
		curs.Decode(&result)
	}

	ctx.IndentedJSON(http.StatusOK, results)
}

// add to form

// get a form
func GetForm(ctx *gin.Context) {

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err != nil {
		return
	}

	var result bson.M
	err = db.InitDb().Collection(db.Collections.Form).FindOne(
		context.TODO(),
		bson.M{"_id": id},
	).Decode(&result)

	ctx.IndentedJSON(http.StatusOK, result)
}

// delete a form
func DeleteForm(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err != nil {
		return
	}

	result, err := db.InitDb().Collection(db.Collections.Form).DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
	return
}
