// package cs
// @Author cuisi
// @Date 2024/6/18 15:06:00
// @Desc
package cs

import "github.com/cuisi521/go-cs-core/net/client"

// Client 客户端对象，用于http client的封装
func Client() *client.Client {
	return client.New()
}
