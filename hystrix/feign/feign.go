////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package feign

import (
	"bytes"
	"crypto/tls"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var defaultHeader = http.Header{"Content-Type": []string{"application/json"}}

func PostJSON(url string, body []byte) ([]byte, error) {
	_, data, err := Request(defaultHeader, http.MethodPost, url, bytes.NewReader(body))
	return data, err
}

func Request(header http.Header, method, url string, body io.Reader) (http.Header, []byte, error) {

	// 忽略证书校验
	var cli http.Client
	cli.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	var req *http.Request
	var res *http.Response
	var err error

	// 请求初始化
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	req.Header = header

	// 发起请求
	res, err = cli.Do(req)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	} else {
		defer func() { _ = res.Body.Close() }()
	}

	// 读取数据
	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}
	return res.Header, data, nil
}

func Json(data *[]byte, path ...interface{}) jsoniter.Any {
	return jsoniter.Get(*data, path...)
}
