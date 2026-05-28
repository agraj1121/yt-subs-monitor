package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Kind string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind string `json:"kind"`
	Id string `json:"id"`
	Stats Stats `json:"statistics"`
}

type Stats struct {
	Views string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

func GetSubscribers() (Item, error) {
	request, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels/", nil)

	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	q := request.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_API_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status : ", resp.Status)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return Item{}, err
	}

	if len(response.Items) == 0 {
		return Item{}, fmt.Errorf("no channel found")
	}

	return response.Items[0], nil

}