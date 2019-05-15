package corll

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log/log"
	"myProject/db"
	"net/http"
	"strconv"
	"time"
)

func TestInsert(g *gin.Context) {
	rsp := new(Rsp)

	var tests []interface{}
	for i := 0; i < 1000000; i++ {
		newuser := new(Test)

		newuser.Email = strconv.Itoa(i)
		newuser.Username = strconv.Itoa(i)
		newuser.Password = strconv.Itoa(i)
		newuser.Phone = strconv.Itoa(i)
		newuser.Address = strconv.Itoa(i)
		newuser.CreateTime = int32(time.Now().Unix())
		newuser.Height = int32(i)
		tests = append(tests, newuser)
	}

	mgo := db.InitMongoDB()
	insertID, err := mgo.Collection(db.Test).InsertMany(context.Background(), tests, nil)
	fmt.Println(insertID)
	if err == nil {
		rsp.Msg = "success"
		rsp.Code = 200
		g.JSON(http.StatusOK, rsp)
		return
	} else {
		rsp.Msg = "faild"
		rsp.Code = 201
		g.JSON(http.StatusOK, rsp)
		return
	}
}
func Test2(g *gin.Context) {
	var Tests = make([]Test, 0)
	rsp := new(Rsp)
	mgo := db.InitMongoDB()

	opts := new(options.FindOptions)
	limit := int64(6)
	skip := int64((1 - 1) * 6) //
	opts.Limit = &limit
	opts.Skip = &skip

	sortMap := make(map[string]interface{})
	sortMap["height"] = -1
	opts.Sort = sortMap

	filter := bson.D{{"email", "1111"}}
	Txs, err := mgo.Collection(db.Test).Find(context.Background(), filter)
	if err != nil {
		rsp.Code = 500
		rsp.Msg = "get data error"
		g.JSON(http.StatusOK, rsp)
		return
	}
	for Txs.Next(context.Background()) {
		elem := new(Test)
		err := Txs.Decode(elem)
		if err != nil {
			log.Error(err)
		}

		Tests = append(Tests, *elem)
	}
	rsp.Data = Tests
	rsp.Code = 200
	rsp.Msg = "success"
	g.JSON(http.StatusOK, rsp)

	return
}