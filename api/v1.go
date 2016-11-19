package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

type V1 struct {
	Whitelist []*regexp.Regexp
	*Resource
}

type authData struct {
	ClientSecret string `json:"client_secret"`
	ClientId     string `json:"client_id"`
	GrantType    string `json:"grant_type"`
	Redirect_Uri string `json:"redirect_uri"`
	Code         string `json:"code"`
}

type tokenData struct {
	AccessToken string `json:"access_token"`
}

//LogInWithInstagram Redirect to instagram authorize API
func (rc *V1) LogInWithInstagram(w http.ResponseWriter, r *http.Request) {
	requestURL := rc.Config.Instagram.AuthUrl
	requestURL += fmt.Sprintf("client_id=%s&", rc.Config.Instagram.ClientId)
	requestURL += fmt.Sprintf("redirect_uri=%s&", rc.Config.Instagram.RedirectUri)
	requestURL += fmt.Sprintf("response_type=%s&", rc.Config.Instagram.ResponseType)
	requestURL += fmt.Sprintf("scope=%s", rc.Config.Instagram.Scope)

	type RedirectUri struct {
		Url string `json:"redirect_uri"`
	}
	url := RedirectUri{
		Url: requestURL,
	}

	jsonData, err := json.Marshal(url)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func requestToken(rc *V1, code string) (interface{}, error) {

	form := url.Values{}
	form.Add("client_id", rc.Config.Instagram.ClientId)
	form.Add("client_secret", rc.Config.Instagram.ClientSecret)
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", rc.Config.Instagram.RedirectUri)

	res, err := http.PostForm(rc.Config.Instagram.TokenUrl, form)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", data)
	defer res.Body.Close()
	return data, nil
}

//Authorize Redirected from to instagram authorize api
func (rc *V1) Authorize(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	instaError := r.URL.Query().Get("error")

	if len(code) == 0 || len(instaError) != 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	retData, err := requestToken(rc, code)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jsonData, err := json.Marshal(retData)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

//Index ...
func (rc *V1) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//SearchMedia given lat,log,radius pull out media
func (rc *V1) SearchMedia(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lng")
	distance := r.URL.Query().Get("distance")
	if len(distance) == 0 {
		distance = "5000"
	}

	if len(lat) == 0 || len(lng) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url := "https://api.instagram.com/v1/media/search"
	url += "?lat=" + lat + "?lng=" + lng + "?distance=" + distance
	url += "?access_token=" + rc.AccessToken

	res, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	//construct json data model
	//convert to data stream
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

//SearchTrendingHashtag ...
func (rc *V1) SearchTrendingHashtag(w http.ResponseWriter, r *http.Request) {

}

//GetTagCounts get hashtag counts in a given hash tags
func (rc *V1) GetTagCounts(w http.ResponseWriter, r *http.Request) {

}

//AddWatchingHashTag add watch hashtag and periodically run api (daily)
//and store result
func (rc *V1) AddWatchingHashTag(w http.ResponseWriter, r *http.Request) {

}

//RemoveWatchingHashtag ...
func (rc *V1) RemoveWatchingHashtag(w http.ResponseWriter, r *http.Request) {

}
