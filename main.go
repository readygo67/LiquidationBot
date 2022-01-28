package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo67/LiquidationBot/config"
	"github.com/readygo67/LiquidationBot/db"
	"github.com/readygo67/LiquidationBot/syncer"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "f", "config.yml", "Path to config.yml")
}

var wg sync.WaitGroup

func main() {
	conf, err := config.New(configFile)
	if err != nil {
		panic(err)
	}

	exist := fileExist(conf.DB)
	if !exist {
		log.Info("DB file does not exist, initiate it")
		dbm.InitDB(conf.DB)
	}
	db, err := leveldb.OpenFile(conf.DB, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c, err := ethclient.Dial(conf.RPCURL)
	if err != nil {
		panic(err)
	}

	log.Info("client dial success", "rpc", conf.RPCURL)
	fmt.Printf("client dial success, rpc:%v\n", conf.RPCURL)
	s := syncer.NewSyncer(c, db, conf)
	s.Start()

	waitExit()

	s.Stop()

}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func waitExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
	log.Info("Received interrupt, shutting down...")
}
