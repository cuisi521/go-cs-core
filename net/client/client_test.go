package client

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

// 测试http get请求
// 请求方式：go test -v -run="TestHttpGet"
func TestHttpGet(t *testing.T) {

	c := New()
	c.SetHeader("Authorization", "test")
	c.SetContentType(httpHeaderContentTypeForm)
	resp, err := c.Get("http://localhost:8001/api/v1/dc/dataCenter/get", nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		result := resp.ReadAllString()
		fmt.Println("result:", result)
	}
	defer func() {
		resp.Close()
	}()

}

func TestHttpGetImgUrl(t *testing.T) {
	c := New()
	c.SetHeader("Authorization", "test")
	c.SetContentType(httpHeaderContentTypeForm)
	resp, err := c.Get("图片url地址", nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		// result := resp.ReadAll()
		file, err := os.Create("/Users/cuisi/Documents/work/study/go-cs-core/cmd/test.jpg")
		if err != nil {
			fmt.Println("Error creating the file:", err)
			return
		}
		defer file.Close()

		// 将响应体的内容写入文件
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println("Error writing the image to the file:", err)
			return
		}

		fmt.Println("Image downloaded successfully and saved as 'downloaded_image.jpg'")
	}
	defer func() {
		resp.Close()
	}()
}

func TestHttpPost(t *testing.T) {
	for j := 0; j < 3; j++ {
		c := New()
		c.SetHeader("Authorization", fmt.Sprint(j))
		c.SetContentType(httpHeaderContentTypeForm)
		data := map[string]string{"t": "111", "t2": "222"}
		byteData, _ := json.Marshal(data)
		resp, err := c.Post("http://localhost:8001/api/v1/dc/dataCenter/post", byteData)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			result := resp.ReadAllString()
			fmt.Println("result:", result)
		}
		defer func() {
			resp.Close()
		}()
	}
}

func TestHttpPut(t *testing.T) {
	for j := 0; j < 3; j++ {
		c := New()
		c.SetHeader("Authorization", fmt.Sprint(j))
		c.SetContentType(httpHeaderContentTypeForm)
		data := map[string]string{"t": fmt.Sprintf("s%v", j), "t2": fmt.Sprintf("v%v", j)}
		byteData, _ := json.Marshal(data)
		resp, err := c.Put("http://localhost:8001/api/v1/dc/dataCenter/put", byteData)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			result := resp.ReadAllString()
			fmt.Println("result:", result)
		}
		defer func() {
			resp.Close()
		}()
	}
}

func TestHttpPatch(t *testing.T) {
	for j := 0; j < 3; j++ {
		c := New()
		c.SetHeader("Authorization", fmt.Sprint(j))
		c.SetContentType(httpHeaderContentTypeForm)
		data := map[string]string{"t": fmt.Sprintf("s%v", j), "t2": fmt.Sprintf("v%v", j)}
		byteData, _ := json.Marshal(data)
		resp, err := c.PATCH("http://localhost:8001/api/v1/dc/dataCenter/pathc", byteData)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			result := resp.ReadAllString()
			fmt.Println("result:", result)
		}
		defer func() {
			resp.Close()
		}()
	}
}

func TestHttpDelete(t *testing.T) {
	for j := 0; j < 3; j++ {
		c := New()
		c.SetHeader("Authorization", fmt.Sprint(j))
		c.SetContentType(httpHeaderContentTypeForm)
		data := map[string]string{"t": fmt.Sprintf("s%v", j), "t2": fmt.Sprintf("v%v", j)}
		byteData, _ := json.Marshal(data)
		resp, err := c.Delete("http://localhost:8001/api/v1/dc/dataCenter/delete", byteData)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			result := resp.ReadAllString()
			fmt.Println("result:", result)
		}
		defer func() {
			resp.Close()
		}()
	}
}

func TestProxy(t *testing.T) {
	err := cfResp()
	if err != nil {
		cfResp()
	}

}

func cfResp() error {
	c := New()
	// c.SetHeader("Authorization", "test")
	// c.SetContentType(httpHeaderContentTypeForm)
	resp, err := c.Get("http://123.207.199.207:9006/api/v1/dc/dataCenter/get", nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		result := resp.ReadAllString()
		fmt.Println("result:", result)
	}
	defer func() {
		resp.Close()
	}()
	return err
}
