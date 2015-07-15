package client

import (
    "net/http"
    "strings"
    "errors"
    "encoding/json"
    // "fmt"
    "io/ioutil"
    "bytes"
)

const API_VERSION = "v1";
const HOST = "https://api.elasticpush.com/"

type Client struct{
    token string
    tokenSecret string
    apiId string
    url string
    clientId string
}

type HttpResponse struct{
    Body []byte
    Code int
}

func (r *HttpResponse) GetBody() []byte{
    return r.Body
}

func (r *HttpResponse) GetCode() int{
    return r.Code
}

func New( token, apiId string) (*Client, error){
    url := HOST + API_VERSION + "/apps/" + apiId
    splittedToken := strings.Split(token, ":")
    if len(splittedToken) != 2 {
        return nil, errors.New("Invalid token format")
    }
    c := &Client{ splittedToken[0], splittedToken[1], apiId, url, ""}
    return c, nil
}

func (c *Client) SetClientId( id string) {
    c.clientId = id
}

func (c *Client) Dispatch( channel, event string, data interface{}) (*HttpResponse, error){

    requestBody := make(map[string]interface{})
    requestBody["channel"] = channel;
    requestBody["event"] = event;
    requestBody["data"] = data;
    if c.clientId != ""{
        requestBody["identifier"] = c.clientId
    }

    requestBodyJson, err := json.Marshal(requestBody)
    if err != nil{
        return nil, err
    }
    
    req, err := http.NewRequest("POST", c.url +"/events" , bytes.NewBuffer(requestBodyJson))
    if err != nil {
        return nil, err
    }

    req.Header.Set("X-Token", c.token)
    req.Header.Set("X-Secret-Token", c.tokenSecret)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    // Send the request via a client
    client := &http.Client{}

    // fmt.Printf("%+v", req)

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if body, err := ioutil.ReadAll(resp.Body); err == nil {
        return &HttpResponse{body,resp.StatusCode}, nil
    }
    return nil, err
}