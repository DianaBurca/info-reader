package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func doRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ReadHandler ...
func ReadHandler(c *gin.Context) {

	params := c.Request.URL.Query()
	cityName := params["city_name"][0]

	findQry := "SELECT payload, COUNT(timestamp) AS exists FROM info WHERE city_name=? ORDER BY timestamp DESC LIMIT 1"
	result := map[string]interface{}{}
	err := CassandraSession.Query(findQry, cityName).MapScan(result)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	if result["exists"].(int64) != 0 {
		rPaylod := []byte(result["payload"].(string))
		payload := map[string]interface{}{}
		json.Unmarshal(rPaylod, &payload)
		c.JSON(http.StatusOK, payload)
	} else {
		resp, err := doRequest(fmt.Sprintf("http://fetcher/fetch?city=%s", cityName))
		if err == nil {
			var bodyBytes []byte
			if resp.StatusCode == http.StatusOK {
				bodyBytes, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				payload := map[string]interface{}{}
				json.Unmarshal(bodyBytes, &payload)
				c.JSON(http.StatusOK, payload)
			}
		} else {
			fmt.Println("err: ", err)
		}
	}

}
