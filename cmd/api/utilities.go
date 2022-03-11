package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter,status int,data interface{},wrap string) error {
	//we dont know what data will come in so above=> data interface{}
	wrapper := make(map[string]interface{})  //key is string and any kind of content
	wrapper[wrap]=data
	js,err:=json.Marshal(wrapper)
	/* the json will look like below
	{
	    "moviewrap": {
	        "id": 123,
	        "title": "mission impossible",
	        "description": "some description",
	        "year": 2021,
	        "updated_at": "2022-03-11T23:51:14.663321+05:45"
	    }
	}
	*/
	if err!= nil{
		return err
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
