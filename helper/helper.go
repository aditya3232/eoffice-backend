package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Response struct {
	Meta       Meta        `json:"meta"`
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Pagination struct {
	Page          int `form:"page" json:"page" binding:"required"`
	Limit         int `form:"limit" json:"limit" binding:"required"`
	Total         int `json:"total"`
	TotalFiltered int `json:"total_filtered"`
}

func NewPagination(page int, limit int) Pagination {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

type Sort struct {
	Sort  string `form:"sort" json:"sort"`
	Order string `form:"order" json:"order"`
}

func NewSort(sort string, order string) Sort {
	if sort == "" {
		sort = "id"
	}

	if order == "" {
		order = "asc"
	}

	return Sort{
		Sort:  sort,
		Order: order,
	}
}

func APIResponse(message string, code int, pagination Pagination, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
	}

	jsonResponse := Response{
		Meta:       meta,
		Pagination: pagination,
		Data:       data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func UnmarshalError(err error) []string {
	var errors []string
	errors = append(errors, "Format Json Error, cek kembali penulisan format json", "Atau Json.UnmarshalTypeError, cek kembali value tipe data dibeberapa field")
	return errors
}

func FormatError(err error) []string {
	var errors []string
	errors = append(errors, fmt.Sprintf("%s", err))
	return errors
}

func StringToDate(dateString string) time.Time {
	StringToDate, _ := time.Parse("2006-01-02", dateString)
	return StringToDate
}

func StringToDateTime(dateTimeString string) time.Time {
	StringToDateTime, _ := time.Parse("2006-01-02 15:04:05", dateTimeString)
	return StringToDateTime

}

func DateTimeToString(t time.Time) string {
	date := t.Format("2006-01-02 15:04:05")
	date = strings.Replace(date, "T", " ", -1)
	date = strings.Replace(date, "Z", "", -1)
	return date
}

// function to check if a string is an email and valid
func IsEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func StrToInt(str string) int {
	var i int
	fmt.Sscanf(str, "%d", &i)
	return i
}

func InArray(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		value := reflect.ValueOf(haystack)
		for i := 0; i < value.Len(); i++ {
			if reflect.DeepEqual(value.Index(i).Interface(), needle) {
				return true
			}
		}
	}
	return false
}

func GetJSONTags(t reflect.Type) []string {
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
		if field.Type.Kind() == reflect.Struct {
			embeddedTags := GetJSONTags(field.Type)
			tags = append(tags, embeddedTags...)
		}
	}
	return tags
}

func QueryParamsToMap(c *gin.Context, s interface{}) map[string]string {
	params := make(map[string]string)
	queryParams := c.Request.URL.Query()

	jsonTags := GetJSONTags(reflect.TypeOf(s))

	exclude := map[string]bool{
		"comment": true,
		"limit":   true,
		"page":    true,
		"sort":    true,
		"order":   true,
		"dir":     true,
	}
	include := []string{
		"_all_",
	}

	if len(include) > 0 {
		jsonTags = append(jsonTags, include...)
	}

	_all_ := false
	if _, ok := queryParams["_all_"]; ok {
		_all_ = true
		for _, jsonTag := range jsonTags {
			queryParams[jsonTag] = []string{queryParams["_all_"][0]}
		}
	}

	for fieldName, fieldValue := range queryParams {
		if _, ok := exclude[fieldName]; ok {
			continue
		}

		if !InArray(fieldName, jsonTags) && !_all_ {
			continue
		}

		fieldName = getGormColumnNameFromJSON(s, fieldName)
		params[fieldName] = fieldValue[0]
	}

	return params
}

func getGormColumnNameFromJSON(s interface{}, jsonName string) string {
	// Get the type of the struct
	t := reflect.TypeOf(s)

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the value of the `json` tag
		jsonTag := field.Tag.Get("json")

		// Check if the `json` tag matches the specified name
		if jsonTag == jsonName {
			// Get the value of the `gorm` tag
			gormTag := field.Tag.Get("gorm")

			// Get the value of the `column` tag within the `gorm` tag
			switch {
			case strings.HasPrefix(gormTag, "column:"):
				columnTag := strings.Split(gormTag, ";")[0]
				return strings.TrimPrefix(columnTag, "column:")
			default:
				return jsonName
			}
		}
	}

	return jsonName
}

func ConstructWhereClause(query *gorm.DB, filter map[string]string) *gorm.DB {
	var allValue string

	// check if filter has _all_ key
	if v, ok := filter["_all_"]; ok {
		allValue = v
		delete(filter, "_all_")
	}

	for key, value := range filter {
		if value == "" {
			continue
		}

		if allValue != "" {
			value = allValue
		}

		switch {
		case isValidDateRange(value):
			dateRange := strings.Split(value, "|")
			startDate := strings.TrimPrefix(dateRange[0], "startDate:")
			endDate := strings.TrimPrefix(dateRange[1], "endDate:")
			query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", key), startDate, endDate)

		case strings.HasPrefix(value, "startDate:"):
			startDate := strings.TrimPrefix(value, "startDate:")
			query = query.Where(fmt.Sprintf("%s >= ?", key), startDate)

		case strings.HasPrefix(value, "endDate:"):
			endDate := strings.TrimPrefix(value, "endDate:")
			query = query.Where(fmt.Sprintf("%s <= ?", key), endDate)

		case strings.Contains(value, "_"):
			value = strings.Replace(value, "_", "%", -1)
			if allValue != "" {
				query = query.Or(fmt.Sprintf("CAST(%s AS TEXT) ILIKE ?", key), allValue)
			} else {
				query = query.Where(fmt.Sprintf("CAST(%s AS TEXT) ILIKE ?", key), value)
			}

		default:
			if allValue != "" {
				query = query.Or(fmt.Sprintf("%s = ?", key), allValue)
			} else {
				query = query.Where(fmt.Sprintf("%s = ?", key), value)
			}
		}
	}

	query = query.Where("deleted_at IS NULL")
	return query
}

func ConstructOrderClause(query *gorm.DB, sort Sort) *gorm.DB {
	if sort.Sort != "" {
		query = query.Order(fmt.Sprintf("%s %s", sort.Sort, sort.Order))
	}
	return query
}

func ConstructPaginationClause(query *gorm.DB, pagination Pagination) *gorm.DB {
	query = query.Limit(pagination.Limit)
	query = query.Offset((pagination.Page - 1) * pagination.Limit)
	return query
}

func isValidDateRange(value string) bool {
	dateRegex := regexp.MustCompile(`^startDate:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}|endDate:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)
	return dateRegex.MatchString(value)
}

func TimePointer(t time.Time) *time.Time {
	return &t
}

func RecoverPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
}

func GetTableName(data interface{}) string {
	// Get the reflect.Type of the input data
	dataType := reflect.TypeOf(data)

	// Make sure the input is a struct
	if dataType.Kind() != reflect.Struct {
		panic("GetTableName() only accepts structs")
	}

	// Check if the struct has a TableName() method
	method, ok := dataType.MethodByName("TableName")
	if !ok {
		panic("GetTableName() only accepts structs with a TableName() method")
	}

	// Call the TableName() method on a zero value of the struct type
	// to get the table name
	v := reflect.New(dataType).Elem()
	tableName := method.Func.Call([]reflect.Value{v})[0].String()

	return tableName
}

func JSONToStruct(data interface{}, s interface{}) error {
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Convert the JSON to a struct
	err = json.Unmarshal(jsonData, s)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func StructToJSON(s interface{}) (string, error) {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
