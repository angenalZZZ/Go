package go_opentsdb

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"angenalZZZ/go-program/go-opentsdb/client"
	"angenalZZZ/go-program/go-opentsdb/config"
)

var Db *client.Client
var op *config.OpenTSDBConfig

// 初始化Client
func Init() {
	if Db != nil {
		return
	}

	// config
	op = &config.OpenTSDBConfig{Host: "127.0.0.1:4242"}

	db, e := client.NewClient(*op)
	if e != nil {
		log.Fatal(e) // 中断程序时输出
	}

	// check
	if e := db.Ping(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
	Db = &db
}

// 数据库 OpenTSDB Client close
func ShutdownClient() {
	log.Println("时序数据库 OpenTSDB Client Drop caches..")
	if _, e := (*Db).Dropcaches(); e != nil {
		log.Fatal(e) // 中断程序时输出
	}
}

// 数据库OpenTSDB : go Test()
func Test() {
	Init()
	log.Printf("时序数据库 OpenTSDB Client: Test starting.. Addr: %s\n\n", (*op).Host)

	db := *Db

	//1. POST /api/put
	log.Println("Begin to test POST /api/put.")

	name := []string{"cpu", "disk", "net", "mem", "bytes"}
	cpuDatas := make([]client.DataPoint, 0)
	timestamp := time.Now().Unix()
	time.Sleep(time.Second)
	tags := make(map[string]string)
	tags["host"] = "tsdb-host"
	tags["try-name"] = "tsdb-sample"
	tags["demo-name"] = "opentsdb-test"

	for c := 3; c > 0; c-- {
		for i := 0; i < len(name); i++ {
			time.Sleep(time.Second)
			data := client.DataPoint{
				Metric:    name[i],
				Timestamp: time.Now().Unix(),
				Value:     rand.Float64(),
			}
			data.Tags = tags
			cpuDatas = append(cpuDatas, data)
			log.Printf("  %d.Prepare datapoint %s\n", i, data.String())
		}
	}

	if resp, err := db.Put(cpuDatas, "details"); err != nil {
		log.Printf("  Error occurs when putting datapoints: %v", err)
	} else {
		log.Printf("  %s", resp.String())
	}
	log.Println("Finish testing POST /api/put.\n ")

	//2.1 POST /api/query to query
	log.Println("Begin to test POST /api/query.")
	time.Sleep(2 * time.Second)
	queryParam := client.QueryParam{
		Start: timestamp,
		End:   time.Now().Unix(),
	}
	subqueries := make([]client.SubQuery, 0)
	for _, metric := range name {
		subQuery := client.SubQuery{
			Aggregator: "sum",
			Metric:     metric,
			Tags:       tags,
		}
		subqueries = append(subqueries, subQuery)
	}
	queryParam.Queries = subqueries
	if queryResp, err := db.Query(queryParam); err != nil {
		log.Printf("Error occurs when querying: %v", err)
	} else {
		log.Printf("%s", queryResp.String())
	}
	log.Println("Finish testing POST /api/query.\n ")

	//2.2 POST /api/query/last
	log.Println("Begin to test POST /api/query/last.")
	time.Sleep(1 * time.Second)
	subqueriesLast := make([]client.SubQueryLast, 0)
	for _, metric := range name {
		subQueryLast := client.SubQueryLast{
			Metric: metric,
			Tags:   tags,
		}
		subqueriesLast = append(subqueriesLast, subQueryLast)
	}
	queryLastParam := client.QueryLastParam{
		Queries:      subqueriesLast,
		ResolveNames: true,
		BackScan:     24,
	}
	if queryLastResp, err := db.QueryLast(queryLastParam); err != nil {
		log.Printf("Error occurs when querying last: %v", err)
	} else {
		log.Printf("%s", queryLastResp.String())
	}
	log.Println("Finish testing POST /api/query/last.\n ")

	//2.3 POST /api/query to delete
	log.Println("Begin to test POST /api/query to delete.")
	queryParam.Delete = true
	if queryResp, err := db.Query(queryParam); err != nil {
		log.Printf("Error occurs when deleting: %v", err)
	} else {
		log.Printf("%s", queryResp.String())
	}

	time.Sleep(5 * time.Second)
	log.Println("Query again which shoud return null.")
	queryParam.Delete = false
	if queryResp, err := db.Query(queryParam); err != nil {
		log.Printf("Error occurs when quering: %v", err)
	} else {
		log.Printf("%s", queryResp.String())
	}
	log.Println("Finish testing POST /api/query to delete.\n ")

	//3. GET /api/aggregators
	log.Println("Begin to test GET /api/aggregators.")
	aggreResp, err := db.Aggregators()
	if err != nil {
		log.Printf("Error occurs when acquiring aggregators: %v", err)
	} else {
		log.Printf("%s", aggreResp.String())
	}
	log.Println("Finish testing GET /api/aggregators.\n ")

	//4. GET /api/config
	log.Println("Begin to test GET /api/config.")
	configResp, err := db.Config()
	if err != nil {
		log.Printf("Error occurs when acquiring config info: %v", err)
	} else {
		log.Printf("%s", configResp.String())
	}
	log.Println("Finish testing GET /api/config.\n ")

	//5. Get /api/serializers
	log.Println("Begin to test GET /api/serializers.")
	serilResp, err := db.Serializers()
	if err != nil {
		log.Printf("Error occurs when acquiring serializers info: %v", err)
	} else {
		log.Printf("%s", serilResp.String())
	}
	log.Println("Finish testing GET /api/serializers.\n ")

	//6. Get /api/stats
	log.Println("Begin to test GET /api/stats.")
	statsResp, err := db.Stats()
	if err != nil {
		log.Printf("Error occurs when acquiring stats info: %v", err)
	} else {
		log.Printf("%s", statsResp.String())
	}
	log.Println("Finish testing GET /api/stats.\n ")

	//7. Get /api/suggest
	log.Println("Begin to test GET /api/suggest.")
	typeValues := []string{client.TypeMetrics, client.TypeTagk, client.TypeTagv}
	for _, typeItem := range typeValues {
		sugParam := client.SuggestParam{
			Type: typeItem,
		}
		log.Printf("  Send suggest param: %s", sugParam.String())
		sugResp, err := db.Suggest(sugParam)
		if err != nil {
			log.Printf("  Error occurs when acquiring suggest info: %v\n", err)
		} else {
			log.Printf("  Recevie response: %s\n", sugResp.String())
		}
	}
	log.Println("Finish testing GET /api/suggest.\n ")

	//8. Get /api/version
	log.Println("Begin to test GET /api/version.")
	versionResp, err := db.Version()
	if err != nil {
		log.Printf("Error occurs when acquiring version info: %v", err)
	} else {
		log.Printf("%s", versionResp.String())
	}
	log.Println("Finish testing GET /api/version.\n ")

	//9. Get /api/dropcaches
	log.Println("Begin to test GET /api/dropcaches.")
	dropResp, err := db.Dropcaches()
	if err != nil {
		log.Printf("Error occurs when acquiring dropcaches info: %v", err)
	} else {
		log.Printf("%s", dropResp.String())
	}
	log.Println("Finish testing GET /api/dropcaches.\n ")

	//10. POST /api/annotation
	log.Println("Begin to test POST /api/annotation.")
	custom := make(map[string]string, 0)
	custom["owner"] = "tsdb"
	custom["host"] = "tsdb-host"
	addedST := time.Now().Unix()
	addedTsuid := "000001000001000002"
	anno := client.Annotation{
		StartTime:   addedST,
		Tsuid:       addedTsuid,
		Description: "tsdb test annotation",
		Notes:       "These would be details about the event, the description is just a summary",
		Custom:      custom,
	}
	if queryAnnoResp, err := db.UpdateAnnotation(anno); err != nil {
		log.Printf("Error occurs when posting annotation info: %v", err)
	} else {
		log.Printf("%s", queryAnnoResp.String())
	}
	log.Println("Finish testing POST /api/annotation.\n ")

	//11. GET /api/annotation
	log.Println("Begin to test GET /api/annotation.")
	queryAnnoMap := make(map[string]interface{}, 0)
	queryAnnoMap[client.AnQueryStartTime] = addedST
	queryAnnoMap[client.AnQueryTSUid] = addedTsuid
	if queryAnnoResp, err := db.QueryAnnotation(queryAnnoMap); err != nil {
		log.Printf("Error occurs when acquiring annotation info: %v", err)
	} else {
		log.Printf("%s", queryAnnoResp.String())
	}
	log.Println("Finish testing GET /api/annotation.\n ")

	//12. GET /api/annotation
	log.Println("Begin to test DELETE /api/annotation.")
	if queryAnnoResp, err := db.DeleteAnnotation(anno); err != nil {
		log.Printf("Error occurs when deleting annotation info: %v", err)
	} else {
		log.Printf("%s", queryAnnoResp.String())
	}
	log.Println("Finish testing DELETE /api/annotation.\n ")

	//13. POST /api/annotation/bulk
	log.Println("Begin to test POST /api/annotation/bulk.")
	bulkAnnNum := 4
	anns := make([]client.Annotation, 0)
	bulkAddBeginST := time.Now().Unix()
	addedTsuids := make([]string, bulkAnnNum)
	for i := 0; i < bulkAnnNum-1; i++ {
		addedST := time.Now().Unix()
		addedTsuid := fmt.Sprintf("%s%d", "00000100000100000", i)
		addedTsuids = append(addedTsuids, addedTsuid)
		anno := client.Annotation{
			StartTime:   addedST,
			Tsuid:       addedTsuid,
			Description: "tsdb test annotation",
			Notes:       "These would be details about the event, the description is just a summary",
		}
		anns = append(anns, anno)
	}
	if bulkAnnoResp, err := db.BulkUpdateAnnotations(anns); err != nil {
		log.Printf("Error occurs when posting bulk annotation info: %v", err)
	} else {
		log.Printf("%s", bulkAnnoResp.String())
	}
	log.Println("Finish testing POST /api/annotation/bulk.\n ")

	//14. DELETE /api/annotation/bulk
	log.Println("Begin to test DELETE /api/annotation/bulk.")
	bulkAnnoDelete := client.BulkAnnoDeleteInfo{
		StartTime: bulkAddBeginST,
		Tsuids:    addedTsuids,
		Global:    false,
	}
	if bulkAnnoResp, err := db.BulkDeleteAnnotations(bulkAnnoDelete); err != nil {
		log.Printf("Error occurs when deleting bulk annotation info: %v", err)
	} else {
		log.Printf("%s", bulkAnnoResp.String())
	}
	log.Println("Finish testing DELETE /api/annotation/bulk.\n ")

	//15. GET /api/uid/uidmeta
	log.Println("Begin to test GET /api/uid/uidmeta.")
	metaQueryParam := make(map[string]string, 0)
	metaQueryParam["type"] = client.TypeMetrics
	metaQueryParam["uid"] = "00003A"
	if resp, err := db.QueryUIDMetaData(metaQueryParam); err != nil {
		log.Printf("Error occurs when querying uidmetadata info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}

	log.Println("Finish testing GET /api/uid/uidmeta.\n ")

	//16. POST /api/uid/uidmeta
	log.Println("Begin to test POST /api/uid/uidmeta.")
	uidMetaData := client.UIDMetaData{
		Uid:         "00002A",
		Type:        "metric",
		DisplayName: "System CPU Time",
	}
	if resp, err := db.UpdateUIDMetaData(uidMetaData); err != nil {
		log.Printf("Error occurs when posting uidmetadata info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}
	log.Println("Finish testing POST /api/uid/uidmeta.\n ")

	//17. DELETE /api/uid/uidmeta
	log.Println("Begin to test DELETE /api/uid/uidmeta.")
	uidMetaData = client.UIDMetaData{
		Uid:  "00003A",
		Type: "metric",
	}
	if resp, err := db.DeleteUIDMetaData(uidMetaData); err != nil {
		log.Printf("Error occurs when deleting uidmetadata info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}
	log.Println("Finish testing DELETE /api/uid/uidmeta.\n ")

	//18. POST /api/uid/assign
	log.Println("Begin to test POST /api/uid/assign.")
	metrics := []string{"sys.cpu.0", "sys.cpu.1", "illegal!character"}
	tagk := []string{"host"}
	tagv := []string{"web01", "web02", "web03"}
	assignParam := client.UIDAssignParam{
		Metric: metrics,
		Tagk:   tagk,
		Tagv:   tagv,
	}
	if resp, err := db.AssignUID(assignParam); err != nil {
		log.Printf("Error occurs when assgining uid info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}
	log.Println("Finish testing POST /api/uid/assign.\n ")

	//19. GET /api/uid/tsmeta
	log.Println("Begin to test GET /api/uid/tsmeta.")
	if resp, err := db.QueryTSMetaData("000001000001000001"); err != nil {
		log.Printf("Error occurs when querying tsmetadata info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}
	log.Println("Finish testing GET /api/uid/tsmeta.\n ")

	//20. POST /api/uid/tsmeta
	log.Println("Begin to test POST /api/uid/tsmeta.")
	custom = make(map[string]string, 0)
	custom["owner"] = "tsdb"
	custom["department"] = "paas dep"
	tsMetaData := client.TSMetaData{
		Tsuid:       "000001000001000001",
		DisplayName: "System CPU Time for Webserver 01",
		Custom:      custom,
	}
	if resp, err := db.UpdateTSMetaData(tsMetaData); err != nil {
		log.Printf("Error occurs when posting tsmetadata info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}
	log.Println("Finish testing POST /api/uid/tsmeta.\n ")

	//21. DELETE /api/uid/tsmeta
	log.Println("Begin to test DELETE /api/uid/tsmeta.")
	tsMetaData = client.TSMetaData{
		Tsuid: "000001000001000001",
	}
	if resp, err := db.DeleteTSMetaData(tsMetaData); err != nil {
		log.Printf("Error occurs when deleting tsmetadata info: %v", err)
	} else {
		log.Printf("%s", resp.String())
	}
	log.Println("Finish testing DELETE /api/uid/tsmeta.\n ")
}
