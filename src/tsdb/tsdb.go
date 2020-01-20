package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
 * @Package Name: tsdb
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-4-25 下午4:21
 * @Description:
 */



const (
	Conn_timeout    time.Duration = time.Second * 3
	Default_timeout time.Duration = time.Second * 3
)

type Connection struct {
	conn   net.Conn
	lock   *sync.RWMutex
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewConn(address string) (*Connection, error) {
	_, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTimeout("tcp", address, Conn_timeout)
	if err != nil {
		return nil, err
	}

	return &Connection{
		conn:   conn,
		lock:   new(sync.RWMutex),
		Reader: bufio.NewReader(conn),
		Writer: bufio.NewWriter(conn),
	}, nil
}

func (conn *Connection) Send(data []byte) error {
	//conn.lock.Lock()
	conn.conn.SetWriteDeadline(time.Now().Add(Default_timeout))
	_, err := conn.Writer.Write(data)
	if err != nil {
		//conn.lock.Unlock()
		return err
	}

	err = conn.Writer.Flush()

	//conn.lock.Unlock()
	return err
}

func (conn *Connection) Close() error {
	//conn.lock.Lock()
	//defer conn.lock.Unlock()
	return conn.conn.Close()
}

type TsdbItem struct {
	Metric    string            `json:"metric"`
	Tags      map[string]string `json:"tags"`
	Value     float64           `json:"value"`
	Timestamp int64             `json:"timestamp"`
}

func (this *TsdbItem) String() string {
	return fmt.Sprintf(
		"<Metric:%s, Tags:%v, Value:%v, TS:%d>",
		this.Metric,
		this.Tags,
		this.Value,
		this.Timestamp,
	)
}

func (this *TsdbItem) TsdbString() (s string) {
	s = fmt.Sprintf("put %s %d %.3f ", this.Metric, this.Timestamp, this.Value)

	for k, v := range this.Tags {
		key := strings.ToLower(strings.Replace(k, " ", "_", -1))
		value := strings.Replace(v, " ", "_", -1)
		s += key + "=" + value + " "
	}

	return s
}

type MetaData struct {
	Metric      string            `json:"metric"`
	Endpoint    string            `json:"endpoint"`
	Timestamp   int64             `json:"timestamp"`
	Step        int64             `json:"step"`
	Value       float64           `json:"value"`
	CounterType string            `json:"counterType"`
	Tags        map[string]string `json:"tags"`
}

func convert2TsdbItem(d *MetaData) *TsdbItem {
	t := TsdbItem{Tags: make(map[string]string)}

	for k, v := range d.Tags {
		t.Tags[k] = v
	}
	t.Tags["endpoint"] = d.Endpoint
	t.Metric = d.Metric
	t.Timestamp = d.Timestamp
	t.Value = d.Value
	return &t
}

func NewMetaData() *MetaData {
	return &MetaData{
		Metric:      "tsdb.status",
		Endpoint:    "opentsdb",
		Tags:        make(map[string]string),
		Step:        60,
		Value:       10,
		Timestamp:   time.Now().Unix(),
		CounterType: "GUAGE",
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

//put data  to opentsdb
func PutDataToTsdb() {
	//c, err := NewConn("10.100.101.70:4242")
	//c, err := NewConn("10.114.113.22:9091")
	c, err := NewConn("10.114.113.22:4242")
	if err != nil {
		log.Println(err)
		return
	}
	var closeCh chan struct{} = make(chan struct{}, 1)

	go func(closeCh chan struct{}, c *Connection) {
		ticker := time.NewTicker(Default_timeout)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				d := convert2TsdbItem(NewMetaData())
				var tsdbBuffer bytes.Buffer
				tsdbBuffer.WriteString(d.TsdbString())
				tsdbBuffer.WriteString("\n")
				err := c.Send(tsdbBuffer.Bytes())
				if err != nil {
					log.Println(err)
				}

				log.Println("send tsdb message success, msg content:", d.String())

			case <-closeCh:
				return
			}
		}
	}(closeCh, c)

	/*
			//windows not support that.
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			log.Println(os.Getpid(), "register signal notify")
			<-sigs
			close(closeCh)
			time.Sleep(Default_timeout)
	*/

	select {}
}

type Query struct {
	Aggregator string            `json:"aggregator"`
	Metric     string            `json:"metric"`
	Rate       bool              `json:"rate,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
}

type QueryParams struct {
	Start             interface{} `json:"start"`
	End               interface{} `json:"end,omitempty"`
	Queries           []Query     `json:"queries,omitempty"`
	NoAnnotations     bool        `json:"no_annotations,omitempty"`
	GlobalAnnotations bool        `json:"global_annotations,omitempty"`
	MsResolution      bool        `json:"ms,omitempty"`
	ShowTSUIDs        bool        `json:"show_tsuids,omitempty"`
	ShowSummary       bool        `json:"show_summary,omitempty"`
	ShowQuery         bool        `json:"show_query,omitempty"`
	Delete            bool        `json:"delete,omitempty"`
}

type QueryResponse struct {
	Metric        string             `json:"metric"`
	Tags          map[string]string  `json:"tags"`
	AggregateTags []string           `json:"aggregateTags"`
	Dps           map[string]float64 `json:"dps"`
}

func NewQueryParams() (*QueryParams, error) {
	return &QueryParams{}, nil
}

//query data from opentsdb by metric and time
func QueryByMetricAndTimestamp() {
	//    item := &Query{Tags: make(map[string]string)}
	//    item.Tags["endpoint"] = "hadoop1"
	//    item.Aggregator = "none"
	//    item.Metric = "cpu.idle"
	//format time for q.start and q.end

	now := time.Now()
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(70 * m)
	m2 := now.Add(60 * m)
	start := m1.Format("2019/01/02-15:04:05")
	end := m2.Format("2020/01/02-15:04:05")
	fmt.Println(start)
	fmt.Println(end)
	q, _ := NewQueryParams()
	q.Start = start
	q.End = end
	//    q.Start = "2018/11/07-01:30:00"
	//    q.End = "2018/11/07-01:40:00"
	//    q.Start = "1h-ago"
	//    q.Queries = append(q.Queries, *item)
	//q.Queries = append(q.Queries, Query{Aggregator: "none", Metric: "cpu.idle", Tags: map[string]string{"endpoint": "hadoop1"}})
	q.Queries = append(q.Queries, Query{Aggregator: "none", Metric: "tsdb.status", Tags: map[string]string{"endpoint": "opentsdb"}})
	data, err := json.Marshal(*q)
	if err != nil {
		fmt.Println(err)
	}

	//req, err := http.NewRequest("POST", "http://10.100.101.70:4242/api/query", bytes.NewBuffer(data))
	req, err := http.NewRequest("POST", "http://10.114.113.22:4242/api/query", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	//    fmt.Println("new request :", req)
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("client do response :", resp)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tsbody :", string(body))
	var tsdb []QueryResponse
	err = json.Unmarshal(body, &tsdb)
	if err != nil {
		log.Fatalln("parse fail:", err)
	}
	for k, _ := range tsdb {
		        fmt.Println("Metric:", tsdb[k].Metric)
		        fmt.Println("Tags:", tsdb[k].Tags)
		        fmt.Println("AggregateTags:", tsdb[k].AggregateTags)
		        fmt.Println("Dps", tsdb[k].Dps)
		        fmt.Println("Dps len", len(tsdb[k].Dps))
		t := time.Now()
		sli := make([]int, 0)
		var intstr int
		for tk, vv := range tsdb[k].Dps {
			fmt.Println(tk, ":", vv)
			intstr, _ = strconv.Atoi(tk)
			sli = append(sli, intstr)
		}
		sort.Ints(sli[:])
		for _, sv := range sli {
			fmt.Println(sv, ":", tsdb[k].Dps[strconv.Itoa(sv)])
		}
		fmt.Println("slice创建速度："+time.Now().Sub(t).String(), sli)
	}

}

func main() {

	//QueryByMetricAndTimestamp()
	PutDataToTsdb()

}