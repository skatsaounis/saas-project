package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/skatsaounis/askmeanything/db"
	"github.com/skatsaounis/askmeanything/models"
	"github.com/skatsaounis/askmeanything/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var questionCollection *mongo.Collection = db.GetCollection(db.DB, "questions")
var validate = validator.New()

func CreateQuestion() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var question models.Question
		defer cancel()

		// Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.QuestionResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate.Struct(&question); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.QuestionResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newQuestion := models.Question{
			Id:       primitive.NewObjectID(),
			Question: question.Question,
			Answer:   question.Answer,
		}
		result, err := questionCollection.InsertOne(ctx, newQuestion)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.QuestionResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAQuestion() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		questionId := params["questionId"]
		var question models.Question
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(questionId)

		err := questionCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&question)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.QuestionResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": question}}
		json.NewEncoder(rw).Encode(response)
	}
}

func EditAQuestion() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		questionId := params["questionId"]
		var question models.Question
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(questionId)

		// Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.QuestionResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate.Struct(&question); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.QuestionResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		update := bson.M{"question": question.Question, "answer": question.Answer}
		result, err := questionCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Get updated question details
		var updatedQuestion models.Question
		if result.MatchedCount == 1 {
			err := questionCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedQuestion)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.QuestionResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedQuestion}}
		json.NewEncoder(rw).Encode(response)
	}
}

func DeleteAQuestion() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		questionId := params["questionId"]
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(questionId)

		result, err := questionCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			rw.WriteHeader(http.StatusNotFound)
			response := responses.QuestionResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Question with specified ID not found!"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.QuestionResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Question successfully deleted!"}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAllQuestions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var questions []models.Question
		defer cancel()

		results, err := questionCollection.Find(ctx, bson.M{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleQuestion models.Question
			if err = results.Decode(&singleQuestion); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.QuestionResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}
			questions = append(questions, singleQuestion)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.QuestionResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": questions}}
		json.NewEncoder(rw).Encode(response)
	}
}
