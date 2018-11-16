package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
)

func main() {
	// webservice
	ws := new(restful.WebService)
	ws.Path("/test").Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	// route
	ws.Route(ws.GET("/testhandler").To(testHandler))
	// containner
	restful.NewContainer().Add(ws)

}


func testHandler(request *restful.Request, response *restful.Response){
	fmt.Println(request)
	response.WriteEntity(sampleResp{
		status: "0",
		msg: "接收请求成功",
	})
}

type sampleResp struct {
	status string
	msg	string
}
