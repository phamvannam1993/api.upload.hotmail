package logging

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log.autofarmer.go/config"
	CoreControllers "log.autofarmer.go/modules/core/controllers"
	CoreModels "log.autofarmer.go/modules/core/models"
	"log.autofarmer.go/util"
)

var (
	loggingModels = LoggingModels{
		CoreModels.MongoDBModels{
			Collection: "logging",
		},
	}
)

type (
	IndexQueryFields struct {
		Status string `query:"status"` // all | active | inactive
		Page   int64  `query:"page"`
	}
)
type LoggingController struct {
	CoreControllers.BaseControllers
}

func (controller *LoggingController) Index(c echo.Context) error {
	var (
		ctx         = c.Request().Context()
		queryValues = c.Get("query").(IndexQueryFields)
		query       = CoreModels.AppQuery{
			Status: queryValues.Status,
			Page:   queryValues.Page,
			Limit:  config.Limit20,
			Sort:   bson.M{"createdAt": -1},
		}
	)
	result, total := loggingModels.GetAll(ctx, query)
	return util.Response200(c, echo.Map{
		"pages": result,
		"total": total,
		"limit": query.Limit,
	}, "")
}

type CreateFieldJson struct {
	PageID      string  `json:"pageId"`
	URL         string  `json:"url"`
	Name        string  `json:"name"`
	MaxLiked    int     `json:"maxLiked"`
	CostPerLike float64 `json:"costPerLike"`
	CreatedAt   string  `json:"createdAt"`
	Email       string  `json:"email"`
}

//insert
func (controller *LoggingController) Store(c echo.Context) error {
	return util.Response200(c, "add-queue-insert", "")
	/*var (
		ctx     = c.Request().Context()
		payload = c.Get("payload").(interface{})
	)
	facebookPageBSON, err := loggingModels.Insert(ctx, payload)
	if err != nil {
		return util.Response400(c, echo.Map{}, err.Error())
	}
	return util.Response200(c, echo.Map{
		"facebookPageBSON": facebookPageBSON,
		"payload":          payload,
	}, "")*/
}
func (controller *LoggingController) Show(c echo.Context) error {
	var (
		ctx      = c.Request().Context()
		IDString = c.Param("id")
	)
	objID, _ := primitive.ObjectIDFromHex(IDString)
	data := loggingModels.FindOneByID(ctx, objID)
	return util.Response200(c, data, "")
}
func (controller *LoggingController) Update(c echo.Context) error {
	var (
		ctx        = c.Request().Context()
		IDString   = c.Param("id")
		payload    = c.QueryParams()
		filter     = bson.M{"_id": IDString}
		updateData = bson.M{
			"$set": payload,
		}
	)
	err := loggingModels.Update(ctx, filter, updateData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	return util.Response200(c, nil, "")
}
func (controller *LoggingController) Delete(c echo.Context) error {
	var (
		ctx      = c.Request().Context()
		IDString = c.Param("id")
	)
	objID, _ := primitive.ObjectIDFromHex(IDString)
	data := loggingModels.Delete(ctx, objID)
	return util.Response200(c, data, "")
}
func (controller *LoggingController) CheckExisted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx      = c.Request().Context()
			IDString = c.Param("id")
		)
		objID, _ := primitive.ObjectIDFromHex(IDString)
		page := loggingModels.FindOneByID(ctx, objID)
		if page == nil {
			return util.Response404(c, nil, "trang không tim thấy")
		}
		c.Set("page", page)
		return next(c)
	}
}
func (controller *LoggingController) InsertLog(params interface{}) {
	_, err := loggingModels.Insert(context.Background(), params)
	if err != nil {
	}
}
