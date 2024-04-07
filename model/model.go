package model

type PostBody struct {
    Title string `json:"title"`
    StartAt string `json:"startAt"`
    EndAt string `json:"endAt"`
    Conditions []struct{
        AgeStart int `json:"ageStart"`
        AgeEnd int `json:"ageEnd"`
        Country []string `json:"country"`
        Gender string `json:"gender"`
        Platform []string `json:"platform"`
    } `json:"conditions"`
}
type ResultBody struct {
    Items []ResultItem `json:"items"`
}

type ResultItem struct {
    Title string `json:"title"`
    StartAt string `json:"startAt"`
    EndAt string `json:"endAt"`
    
}
