package controllers

import (
    "context"
    "tutorials/configs"
    "tutorials/models"
    "tutorials/responses"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

var stateCollection *mongo.Collection = configs.GetCollection(configs.DB, "states")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var state models.State
		defer cancel()
	
	if err := c.BindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, responses.StateResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

			if validationErr := validate.Struct(&state); validationErr != nil {
				c.JSON(http.StatusBadRequest, responses.StateResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
				return
			}

			newState := models.State{
				Id:       primitive.NewObjectID(),
				Name:     state.Name,
				Capital: state.Capital,
			}

			result, err := stateCollection.InsertOne(ctx, newState)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.StateResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return		
			}

		c.JSON(http.StatusCreated, responses.StateResponse{Status: http.StatusCreated, Message: "OK", Data: map[string]interface{}{"data": result}})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var states []models.State
		defer cancel()
	
		results, err := stateCollection.Find(ctx, bson.M{})
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StateResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleState models.State
			if err = results.Decode(&singleState); err != nil {
				c.JSON(http.StatusInternalServerError, responses.StateResponse{Status: http.StatusInternalServerError, Message: "err", Data: map[string]interface{}{"data": err.Error()}})
			}

			states = append(states, singleState)
		}
		c.JSON(http.StatusOK, responses.StateResponse{Status: http.StatusOK, Message: "ok", Data: map[string]interface{}{"data": states}},)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		StateId := c.Param("stateid")
		var state models.State
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(StateId)

		err := stateCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&state)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StateResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.StateResponse{Status: http.StatusOK, Message: "ok", Data: map[string]interface{}{"data": state}})
	}
}