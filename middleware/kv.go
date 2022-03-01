package middleware

import (
	"context"
	"log"
	// "os"
	"strings"
	"time"

	"go.etcd.io/etcd/client/v3"
	// "google.golang.org/grpc/grpclog"
)

var (
	dialTimeout = 5 * time.Second
	endpoints   = []string{"172.18.0.71:2379", "172.18.0.72:2379", "172.18.0.73:2379", "172.18.0.74:2379", "172.18.0.75:2379"}
)

//KV is wrapper for KV Database
type KV struct {
	cli *clientv3.Client
	err error
}

//New returns singleton instance of KV
func (kv *KV) New() {
	// clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	kv.cli, kv.err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if kv.err != nil {
		log.Fatal(kv.err)
	}
}

//Put upserts data into KV store
func (kv *KV) Put(key, value string) {
	start := time.Now()
	_, kv.err = kv.cli.Put(context.TODO(), key, value)
	if kv.err != nil {
		log.Fatal(kv.err)
	} else {
		log.Printf("Successfully put %s:::::%s to etcd\n", key, value)
		elapsed := time.Since(start)
		log.Printf("Put to etcd took %s \n", elapsed)
	}

}

//Get fetches data from KV store
func (kv *KV) Get(key string) string {
	start := time.Now()
	getResp, err := kv.cli.Get(context.TODO(), key)
	if err != nil {
		log.Fatal(err)
	}

	if getResp.Count >= 1 {
		log.Printf("Successfully got : %s :: for key : %s :: \n", getResp.Kvs[0].Value, getResp.Kvs[0].Key)
		elapsed := time.Since(start)
		log.Printf("Get from etcd took %s \n", elapsed)
		return string(getResp.Kvs[0].Value)
	}
	return "{}"

}

//GetFromKey fetches data after a time range
func (kv *KV) GetFromKey(key string) string {
	start := time.Now()
	getResp, err := kv.cli.Get(context.TODO(), key, clientv3.WithFromKey(), clientv3.WithLimit(0))
	if err != nil {
		log.Fatal(err)
	}

	if getResp.Count >= 1 {
		elapsed := time.Since(start)
		log.Printf("Get from etcd took %s \n", elapsed)
		var rb1 strings.Builder
		rb1.WriteString("[")
		log.Printf("Successfully got : %d :: values for key : %s :: \n", getResp.Count, key)
		for _, ev := range getResp.Kvs {
			rb1.WriteString(string(ev.Value))
			rb1.WriteString(",")
			// log.Printf("Successfully got : %s :: for key : %s :: \n", ev.Value, ev.Key)
		}
		rb2 := strings.TrimSuffix(rb1.String(), ",")
		rb2 = rb2 + "]"
		return rb2
	}
	log.Printf("No Value found for key : %s :: \n", key)
	return "{}"

}

//GetFromKeyWithLimit fetches data after a time range with limit
func (kv *KV) GetFromKeyWithLimit(key string, limit int64) string {
	start := time.Now()
	getResp, err := kv.cli.Get(context.TODO(), key, clientv3.WithFromKey(), clientv3.WithLimit(limit))
	if err != nil {
		log.Fatal(err)
	}

	if getResp.Count >= 1 {
		elapsed := time.Since(start)
		log.Printf("Get from etcd took %s \n", elapsed)
		var rb1 strings.Builder
		rb1.WriteString("[")
		log.Printf("Successfully got : %d :: values for key : %s :: \n", getResp.Count, key)
		for _, ev := range getResp.Kvs {
			rb1.WriteString(string(ev.Value))
			rb1.WriteString(",")
			// log.Printf("Successfully got : %s :: for key : %s :: \n", ev.Value, ev.Key)
		}
		rb2 := strings.TrimSuffix(rb1.String(), ",")
		rb2 = rb2 + "]"
		return rb2
	}
	log.Printf("No Value found for key : %s :: \n", key)
	return "{}"

}
