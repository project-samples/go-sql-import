package app

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"time"

	. "github.com/core-go/io/import"
	q "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ApplicationContext struct {
	Import func(ctx context.Context) error
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}
	generateFileName := func() string {
		fileName := "20060102150402.csv"
		fullPath := filepath.Join("export", fileName)
		return fullPath
	}
	userType := reflect.TypeOf(User{})
	formatter, err := NewFixedLengthFormatter(userType)
	if err != nil {
		return nil, err
	}
	reader, err := NewFileReader(generateFileName)
	if err != nil {
		return nil, err
	}
	writer := q.NewStreamWriter(db, "users", userType, 500)
	importer := NewImporter(db, userType, formatter.ToStruct, func(ctx context.Context, data interface{}, endLineFlag bool) error {
		fmt.Println(data)
		ctx = context.Background()
		if endLineFlag {
			err = writer.Flush(ctx)
			if err != nil {
				return err
			}
		} else {
			if data != nil {
				err := writer.Write(ctx, data)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}, reader.Read)

	return &ApplicationContext{ Import: importer.Import }, nil
}

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" format:"%011s" length:"11" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Username    string     `json:"username" gorm:"column:username" bson:"username" length:"10" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Email       string     `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" length:"31" validate:"email,max=100"`
	Phone       string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" length:"20" validate:"required,phone,max=18"`
	Status      bool       `json:"status" gorm:"column:status" true:"1" false:"0" bson:"status" dynamodbav:"status" format:"%5s" length:"5" firestore:"status" validate:"required"`
	CreatedDate *time.Time `json:"createdDate" gorm:"column:createdDate" bson:"createdDate" length:"10" format:"dateFormat:2006-01-02" dynamodbav:"createdDate" firestore:"createdDate" validate:"required"`
}
