package engine

import (
	"KworkTasksNotifier/src/models"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

// GetCategoryData создает строку с Multi Part Form Data нагрузкой
// для запроса
func GetCategoryData(categoryId int) *strings.Reader {
	data := strings.NewReader("-----------------------------128415455034179316042866893723\r\nContent-Disposition: form-data; name=\"fc\"\r\n\r\n" + string(categoryId) + "\r\n-----------------------------128415455034179316042866893723--\r\n")

	return data
}

// CreateHttpRequest создает http запрос для последующего вызова
func CreateHttpRequest(categoryId int) (req *http.Request, err error) {
	data := GetCategoryData(categoryId)

	req, err = http.NewRequest("POST", "https://kwork.ru/projects", data)
	if err != nil {
		return nil, errors.New("Ошибка при создании запроса")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=---------------------------137124826736163243852560029715")
	req.Header.Set("Origin", "https://kwork.ru")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Referer", "https://kwork.ru/projects?fc=41")
	req.Header.Set("Cookie", "_kmid=005f61a013de158554ac161f34c50793; _kmfvt=1678994428; site_version=desktop; _gcl_au=1.1.1211017322.1678994429; _ym_uid=1678994429792213627; _ym_d=1678994429; _sp_id.b695=97a3e069-d646-4cc7-bcb4-c765736d9f9b.1678994429.51.1681051832.1681043070.9019918d-f441-41ba-b850-490475aea128; _ga=GA1.2.122855717.1678994429; uad=1140418064136c1f084f2096000157; userId=11404180; slrememberme=11404180_%242y%2410%24TOKSm2ZUES010.%2F2SNRZ2ukcRdnq7t2LSDjmdeC%2FqhPQNOUjiMPg.; _kmwl=1; list_type_sdisplay=table; show_mobile_app_banner_date=1; RORSSQIHEK=4ca90df0bad9a2c5516512e71265b848; yandex_client_id=1678994429792213627; google_client_id=122855717.1678994429; _ubtcuid=clg9iucrj000020fpt79b2dk5; _ym_isad=2; _gid=GA1.2.656302312.1681026905; _sp_ses.b695=*; _dc_gtm_UA-68703836-1=1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Accept-Charset", "utf-8")

	return req, nil
}

// GetHttpResponse выполняет http запрос и возвращает ответ
func GetHttpResponse(request *http.Request) (bodyText []byte, err error) {
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyText, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyText, nil
}

// DeserializeResponseBody десериализирует json в массив обьектов KworkResponseModel
func DeserializeResponseBody(bodyText []byte) (objects []models.KworkResponseModel, err error) {
	orders := gjson.Get(string(bodyText), "data.wants")

	objects = make([]models.KworkResponseModel, 0)

	err = json.Unmarshal([]byte(orders.String()), &objects)

	if err != nil {
		return nil, err
	}
	return objects, nil
}

func GetData(categoryId int) (objects []models.KworkResponseModel, err error) {
	request, err := CreateHttpRequest(categoryId)
	if err != nil {
		return nil, err
	}

	response, err := GetHttpResponse(request)
	if err != nil {
		return nil, err
	}

	objects, err = DeserializeResponseBody(response)
	if err != nil {
		return nil, err
	}

	return objects, nil
}
