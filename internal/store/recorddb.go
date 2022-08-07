package store

import(
	"context"
	"time"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	
	"getir-study-service/configs"
	"getir-study-service/internal/dto"
	"getir-study-service/models"
)

var dataCollection *mongo.Collection = configs.GetCollection(configs.DB, "records")

func FindRecords(request dto.RecordRequest) (*dto.RecordResponse, *dto.ErrorResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	var err error
	filter := bson.M{"createdAt": bson.M{"$gt": request.StartDate, "$lt": request.EndDate}, "totalCount": bson.M{"$lt": request.MaxCount, "$gt": request.MinCount}}
	cursor, err := dataCollection.Find(ctx, filter)
	if err != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: err, Message: "Failed to find records"}
	}
	
	// Iterate through the cursor
    var records []*models.Record
    err = cursor.All(ctx, &records)
	if err != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: err, Message: "Failed to find records"}
	}
	defer cursor.Close(ctx)
	
	return &dto.RecordResponse{
		Code:  0,
		Status: "Success",
		Records: records,
	}, nil
}