package general

import (
	"bytes"
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/dto"
	res "compass_mini_api/pkg/util/response"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func GetString(name int) string {
	return strconv.Itoa(name)
}

func GetInt(name string) int {
	i, err := strconv.Atoi(name)
	if err != nil {
		return 0
	}
	return i
}

func OnlyNum(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func ProcessQueryPagination(payload *abstraction.QueryPagination) (int, int) {
	limit, _ := strconv.Atoi(payload.Limit)
	offset, _ := strconv.Atoi(payload.Offset)
	return limit, offset
}

func ProcessQueryOrder(payload *abstraction.QueryOrder) string {
	var order string
	if payload.Order == "" && payload.Direction == "" {
		order = "id"
	} else if payload.Order != "" && payload.Direction == "" {
		order = payload.Order
	} else {
		order = payload.Order + " " + payload.Direction
	}
	return order
}

func ProcessQueryFilter(payload *abstraction.QueryFilter) (string, error) {
	var whereArr []string
	var conditions []abstraction.Condition
	if payload.Conditions == "" {
		return "", nil
	}
	strJson, err := url.QueryUnescape(payload.Conditions)
	if err != nil {
		logrus.Error("Error unescaping query: ", err.Error())
		return "", res.CustomErrorBuilder(http.StatusBadRequest, err.Error(), "Failed filter with your conditions")
	}
	err = json.Unmarshal([]byte(strJson), &conditions)
	if err != nil {
		logrus.Error("Error unmarshal string to json: ", err.Error())
		return "", res.CustomErrorBuilder(http.StatusBadRequest, err.Error(), "Failed filter with your conditions")
	}
	for _, condition := range conditions {
		var column, comparation, value string

		if condition.Column == "name" || condition.Column == "supervisor" || condition.Column == "email" {
			column = condition.Column
			comparation = condition.Comparation
			value = "'" + condition.Value + "'"
			if condition.Comparation == "%" {
				comparation = "ILIKE"
				value = "'%" + condition.Value + "%'"
			} else if condition.Comparation != "=" && condition.Comparation != "%" {
				return "", res.CustomErrorBuilder(http.StatusBadRequest, "filter for that column and comparation cannot be processed", "Failed filter with your conditions")
			}
		} else if condition.Column == "companyid" || condition.Column == "divisionid" || condition.Column == "isactive" {
			column = condition.Column
			comparation = condition.Comparation
			value = condition.Value
			if condition.Comparation == "%" {
				column = "CAST(" + condition.Column + " AS VARCHAR)"
				comparation = "ILIKE"
				value = "'%" + condition.Value + "%'"
			} else if condition.Comparation != "=" && condition.Comparation != "%" {
				return "", res.CustomErrorBuilder(http.StatusBadRequest, "filter for that column and comparation cannot be processed", "Failed filter with your conditions")
			}
		} else if condition.Column == "joindate" {
			column = condition.Column
			comparation = condition.Comparation
			if strings.ToLower(comparation) != "between" {
				return "", res.CustomErrorBuilder(http.StatusBadRequest, "filter for that column and comparation cannot be processed", "Failed filter with your conditions")
			}
			valueDate := strings.Split(condition.Value, "_")
			value = fmt.Sprintf("%s AND %s", "'"+valueDate[0]+"'", "'"+valueDate[1]+"'")
		} else {
			return "", res.CustomErrorBuilder(http.StatusBadRequest, "filter for that column is not available", "Failed filter with your conditions")
		}

		where := column + " " + comparation + " " + value
		whereArr = append(whereArr, where)
	}
	return strings.Join(whereArr, " AND "), nil
}

func StringInSlice(text string, data []string) bool {
	for _, row := range data {
		if row == text {
			return true
		}
	}
	return false
}

func SaveFileEmployeePhoto(payload dto.EmployeePhoto) (*string, error) {
	extImge := []string{".jpg", ".jpeg", ".png", ".gif", ".svg"}
	extension := filepath.Ext(payload.Name)
	nameFile := strings.TrimSuffix(payload.Name, extension)

	base64, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		return nil, res.CustomErrorBuilder(http.StatusUnprocessableEntity, err, "error decode string base64")
	}
	if !strings.Contains(payload.Type, "image/") || !StringInSlice(strings.ToLower(extension), extImge) {
		return nil, res.CustomErrorBuilder(http.StatusNotAcceptable, "status not acceptable", "extension not allowed")
	}
	if payload.Size > 2000000 {
		return nil, res.CustomErrorBuilder(http.StatusNotAcceptable, "status not acceptable", "file size is too large")
	}

	fileName := fmt.Sprintf("%s%s", strings.ReplaceAll(nameFile, " ", "")+"_"+time.Now().Format("20060102150405"), extension)
	destinationPath := path.Join("../employeephoto", fileName)
	reader := bytes.NewReader(base64)

	dst, err := os.Create(destinationPath)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.UploadFileCreateError, err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, reader); err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.UploadFileDestError, err)
	}

	return &fileName, err
}
