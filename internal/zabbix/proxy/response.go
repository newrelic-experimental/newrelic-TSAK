package proxy

import (
	"bytes"
	"io/ioutil"
	"compress/zlib"
	"encoding/json"
	"fmt"
)

// ProxyResponse class.
type ProxyResponse struct {
	Data     []uint8 `json:"data"`
	Response string  `json:"response"`
	Info     string  `json:"info"`
	JSON     string  `json:"json"`
}

// ProxyResponse class constructor.
func NewProxyResponse(data []uint8) (r *ProxyResponse, err error) {
	if len(data) < 13 {
		err = fmt.Errorf("NewProxyResponse Input data to short = %d", len(data))
		return
	}
	jsonData := data[13:]

	if data[4] == 1 {
		r = &ProxyResponse{Data: data, JSON: string(jsonData)}
	} else if data[4] == 3 {
		buf := bytes.NewReader(data[13:])
		z, err := zlib.NewReader(buf)
		defer z.Close()
		if err != nil {
			return nil, err
		}
		jsonData, err := ioutil.ReadAll(z)
		if err != nil {
			err = fmt.Errorf("Error decompressing response: %v", err)
		}
		r = &ProxyResponse{Data: data, JSON: string(jsonData)}
	} else {
		r = &ProxyResponse{Data: data, JSON: string(jsonData)}
	}
	err = json.Unmarshal([]byte(r.JSON), r)
	if err != nil {
		err = fmt.Errorf("Error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			err = fmt.Errorf("%s ; Syntax error at byte offset %d", err, e.Offset)
		}
		return
	}
	return
}

// ProxyConfigResponse class.
type ProxyConfigResponse struct {
	Globalmacro struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"globalmacro"`
	Hosts struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"hosts"`
	Interface struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"interface"`
	HostsTemplates struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"hosts_templates"`
	Hostmacro struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"hostmacro"`
	Items struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"items"`
	Drules struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"drules"`
	Dchecks struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"dchecks"`
	Regexps struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"regexps"`
	Expressions struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"expressions"`
	Groups struct {
		Fields []string `json:"fields"`
		Data   [][]int  `json:"data"`
	} `json:"groups"`
	Config struct {
		Fields []string `json:"fields"`
		Data   [][]int  `json:"data"`
	} `json:"config"`
	Httptest struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httptest"`
	Httptestitem struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httptestitem"`
	Httpstep struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httpstep"`
	Httpstepitem struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httpstepitem"`
	Data     []uint8 `json:"data"`
	Response string  `json:"response"`
	Info     string  `json:"info"`
}

// ProxyConfigReponse class constructor.
func NewProxyConfigResponse(data []uint8) (r *ProxyConfigResponse, err error) {
	jsonData := data[13:]

	if data[4] == 1 {
		r = &ProxyConfigResponse{Data: data}
		err = json.Unmarshal(jsonData, r)
	} else if data[4] == 3 {
		buf := bytes.NewReader(data[13:])
		z, err := zlib.NewReader(buf)
		defer z.Close()
		if err != nil {
			return nil, err
		}
		jsonData, err := ioutil.ReadAll(z)
		if err != nil {
			err = fmt.Errorf("Error decompressing response: %v", err)
		}
		// fmt.Println("CFG",string(jsonData))
		r = &ProxyConfigResponse{Data: data}
		err = json.Unmarshal(jsonData, r)
	} else {
		r = &ProxyConfigResponse{Data: data}
		err = json.Unmarshal(jsonData, r)
	}
	if err != nil {
		err = fmt.Errorf("Error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			err = fmt.Errorf("%s ; Syntax error at byte offset %d", err, e.Offset)
		}
		return
	}
	return
}

func (response *ProxyConfigResponse) GetHosts() (hosts map[float64]Host) {

	items := make(map[float64][]Item)
	for _, i := range response.Items.Data {
		item := i.([]interface{})
		if len(item) >= 32 {
			items[item[3].(float64)] = append(items[item[3].(float64)], Item{
				Itemid:                item[0].(float64),
				Type:                  item[1].(float64),
				Key_:                  item[4].(string),
			})
		}
	}
	hosts = make(map[float64]Host)
	for _, d := range response.Hosts.Data {
		if len(d) >= 14 {
			if d[2].(float64) != 3 {
				hostid := d[0].(float64)
				hosts[hostid] = Host{
					Hostid:           hostid,
					Host:             d[1].(string),
					Status:           d[2].(float64),
					Name:             d[7].(string),
					Items:            items[hostid],
				}
			}
		}
	}
	return hosts
}

func (response *ProxyConfigResponse) DiscoverHosts() (hosts map[uint64][]interface{}, items map[uint64][]interface{}) {
	hosts = make(map[uint64][]interface{})
	items = make(map[uint64][]interface{})
	for _, i := range response.Items.Data {
		item := i.([]interface{})
		items[uint64(item[0].(float64))] = item
	}
	for _, d := range response.Hosts.Data {
		hosts[uint64(d[0].(float64))] = d
	}
	return
}
