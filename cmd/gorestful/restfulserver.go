package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

func main() {
	// webservice
	ws := new(restful.WebService)
	ws.Path("/test").Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	// route
	ws.Route(ws.GET("/testhandler").To(testHandler))
		//// docs
		//Doc("get all users").
		//Metadata(restfulspec.KeyOpenAPITags, nil).
		//Writes([]SampleResp{}).
		//Returns(200, "OK", []SampleResp{}))

	// containner
	container := restful.NewContainer().Add(ws)
	http.ListenAndServe(":8081", container)
}


func testHandler(request *restful.Request, response *restful.Response){
	resp := SampleResp{
		Status: "0",
		Msg:    "接收请求成功"}
	fmt.Println(resp)
	response.WriteEntity(resp)
}

// 该结构体需要是公有的，让外部去json化
type SampleResp struct {
	Status string `json:"sTatus" description:"identifier of the SampleResp"` // ``因为内部有双引号""，所以这边是``
	Msg	string `json:"MSg" description:"identifier of the SampleResp"`		 // 这个是 struct的tag，类似于java的注解，可以通过反射获取
																			 // json化的时候，会去读取json字段的值作为key
}

