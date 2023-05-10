package db

import (
	"log"

	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/cluster"
	"github.com/fengqk/mars-base/common"
	"github.com/fengqk/mars-base/db"
	"github.com/fengqk/mars-base/network"
	"github.com/fengqk/mars-base/rpc"
)

type (
	IServerMgr interface {
		Init(string) bool
		InitDB() bool
		GetServer() *network.ServerSocket
	}

	ServerMgr struct {
		service *network.ServerSocket
	}

	Config struct {
		common.Server `yaml:"db"`
		common.Mysql  `yaml:"mysql"`
		common.Etcd   `yaml:"etcd"`
		common.Nats   `yaml:"nats"`
		common.Raft   `yaml:"raft"`
		common.Stub   `yaml:"stub"`
	}
)

var (
	CONF   Config
	SERVER ServerMgr
)

func (s *ServerMgr) Init() bool {
	//初始配置文件
	base.ReadConf("mars.yaml", &CONF)

	ShowMessage := func() {
		base.LOG.Println("**********************************************************")
		base.LOG.Printf("\tDB IP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		base.LOG.Println("**********************************************************")
	}
	ShowMessage()

	base.LOG.Println("正在初始化数据库连接...")
	if s.InitDB() {
		base.LOG.Printf("[%s]数据库连接是失败...", CONF.Mysql.Name)
		log.Fatalf("[%s]数据库连接是失败...", CONF.Mysql.Name)
		return false
	}
	base.LOG.Printf("[%s]数据库初始化成功!", CONF.Mysql.Name)

	//初始化socket
	s.service = new(network.ServerSocket)
	s.service.Init(CONF.Server.Ip, CONF.Server.Port)
	s.service.Start()

	//本身db集群管理
	cluster.MGR.InitCluster(&common.ClusterInfo{Type: rpc.SERVICE_DB, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints,
		cluster.WithStubMailBoxEtcd(CONF.Raft.Endpoints, &CONF.Stub))
	cluster.MGR.BindPacketFunc(actor.MGR.PacketFunc)

	PLAYERSAVEMGR.Init()

	return true
}

func (s *ServerMgr) InitDB() bool {
	return db.OpenDB(CONF.Mysql) != nil
}

func (s *ServerMgr) GetServer() *network.ServerSocket {
	return s.service
}
