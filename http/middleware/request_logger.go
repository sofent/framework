package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/totoval/framework/helpers"
	"io/ioutil"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request

		// print request data
		requestData, err := c.GetRawData()
		if err != nil{
			fmt.Println(err.Error())
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData)) // 关键点
		helpers.Dump(string(requestData))

		// collect response data
		responseWriter := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = responseWriter


		c.Next()

		// after request
		// print response data
		helpers.Dump(responseWriter.body.String())

		// access the status we are sending
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}