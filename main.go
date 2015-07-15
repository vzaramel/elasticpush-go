package main

import (
    "github.com/Elasticpush/elasticpush-go/client"
    "fmt"
)

func main(){

    token := "9751998fa08dea907624a815270d294118e9fc5c85be1453bd4f338acf467e01";
    tokenSecret := "d7cb880a62ec0a6910cb449d5c6ca60cad8e233f171ed53f2b91ff88bf99cff3";
    c, err := client.New( token + ":" + tokenSecret , "11")
    if  err != nil {
        panic("Crashed");
    }
    c.SetClientId("06376146296970546")

    data := make(map[string]string)
    data["message"] = "Teste"

    httpResp, err := c.Dispatch("channel-0124", "event-test", data);
    if  err != nil {
        panic("Crashed Hard");
    }
    fmt.Printf("\n%s\n", httpResp.GetBody())
}