package http_inl2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	WMS_URL string = "wms_url"
)

type GetMapCapabilitiesResponse struct {
	Code         uint8  `json:"result"`
	Desc         string `json:"desc"`
	Capabilities string `json:"capabilities"`
}

func CheckParamters(r *http.Request, keys ...string) bool {
	for _, key := range keys {
		value := r.Form.Get(key)
		if value == "" {
			return false
		}
	}

	return true
}

func GetMapCapabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_SERVER_FAILED, ""))
		}
	}()

	r.ParseForm()
	if !CheckParamters(r, WMS_URL) {
		fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_PARAMTER_ERR, ""))

	}

	wms_url := r.Form.Get(WMS_URL)
	log.Println(wms_url)

	http_client := &http.Client{}

	req, e := http.NewRequest("GET", wms_url, nil)
	if e != nil {
		log.Println(e.Error())
	}
	resp, err := http_client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_SERVER_FAILED, err.Error()))
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)

	response, _ := json.Marshal(&GetMapCapabilitiesResponse{
		Code:         0,
		Desc:         HTTP_RESULT[0],
		Capabilities: string(body),
	})

	fmt.Fprint(w, string(response))
}
