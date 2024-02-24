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
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Slug      string             `bson:"slug" json:"slug"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
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

type AddFormDataDto struct {
	FormId string      `json:"formId"`
	Data   primitive.M `json:"data"`
}

// add to form
func CreateFormData(ctx *gin.Context) {

	var addFromDataDto AddFormDataDto
	if err := ctx.BindJSON(&addFromDataDto); err != nil {
		return
	}

	// formId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	formId, err := primitive.ObjectIDFromHex(addFromDataDto.FormId)

	if err != nil {
		println("Error converting formId to objectId")
		return
	}

	// check from
	cursor := db.InitDb().Collection(db.Collections.Form).FindOne(context.TODO(), bson.M{"_id": formId})

	var formJsonData bson.M
	if err := cursor.Decode(&formJsonData); err != nil {
		return
	}

	if formJsonData == nil {
		println("Cannot added data because form does not exist")
		return
	}

	// add form data
	var formData db.FormData = db.FormData{
		ID:        primitive.NewObjectID(),
		FormId:    formId,
		Data:      addFromDataDto.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := db.InitDb().Collection(db.Collections.FormData).InsertOne(context.TODO(), formData)

	if err != nil {
		println("Error inserting form data")
		return
	}

	ctx.IndentedJSON(http.StatusCreated, result)
}

// get a forms data
func GetFromData(ctx *gin.Context) {
	// form id
	formId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		println("from not found")
		return
	}

	filter := bson.M{"formId": formId}
	cursor, err := db.InitDb().Collection(db.Collections.FormData).Find(context.TODO(), filter)

	if err != nil {
		println("Error geting from data")
		return
	}

	var allResults []db.FormData
	if err := cursor.All(context.TODO(), &allResults); err != nil {
		println("error decoding value")
		return
	}

	for _, result := range allResults {
		cursor.Decode(&result)
	}

	ctx.IndentedJSON(http.StatusOK, allResults)
}

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
