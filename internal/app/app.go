package app

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"time"

	imp "github.com/core-go/io/import"
	"github.com/core-go/io/reader"
	v "github.com/core-go/io/validator"
	"github.com/core-go/log"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/writer"
	_ "github.com/lib/pq"
)

type ApplicationContext struct {
	Import func(ctx context.Context) (int, int, error)
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}
	userType := reflect.TypeOf(User{})
	csvType := reader.DelimiterType
	filename := ""
	test := ""
	if csvType == reader.DelimiterType {
		filename = "delimiter.csv"
		test = "10,abraham59E,rory30@example.com,975-283-2267,TRUE,2019-02-20"
	} else {
		filename = "fixedlength.csv"
		test = "00000000001 abraham59             rory30@example.com        975-283-2267 true2019-02-20"
	}
	generateFileName := func() string {
		fullPath := filepath.Join("export", filename)
		return fullPath
	}
	formatter, err := reader.NewFormater(userType, csvType)
	if err != nil {
		return nil, err
	}
	// test formatter ToStruct
	var user User
	formatter.ToStruct(ctx, test, &user)
	fmt.Println("user", user)
	//reader, err := NewFixedlengthFileReader(generateFileName)
	reader, err := reader.NewDelimiterFileReader(generateFileName)
	if err != nil {
		return nil, err
	}
	mp := map[string]interface{}{
		"app": "import users",
		"env": "dev",
	}
	logError := imp.NewErrorHandler(log.ErrorFields, "fileName", "lineNo", &mp)
	writer := writer.NewStreamWriter(db, "userimport", userType, 6)
	// writer := q.NewInserter(db, "userimport", userType)
	validator := v.NewValidator()
	importer := imp.NewImporter(userType, formatter.ToStruct, reader.Read, logError.HandlerException, validator.Validate, logError.HandlerError, filename, writer.Write, writer.Flush)
	return &ApplicationContext{Import: importer.Import}, nil
}

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" format:"%011s" length:"11" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Username    string     `json:"username" gorm:"column:username" bson:"username" length:"10" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Email       string     `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" length:"31" validate:"email,max=100"`
	Phone       string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" length:"20" validate:"required,max=18"`
	Status      bool       `json:"status" gorm:"column:status" true:"1" false:"0" bson:"status" dynamodbav:"status" format:"%5s" length:"5" firestore:"status"`
	CreatedDate *time.Time `json:"createdDate" gorm:"column:createdDate" bson:"createdDate" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate" firestore:"createdDate" validate:"required"`
	// Test        string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" length:"1" validate:"required,max=18"`
}
