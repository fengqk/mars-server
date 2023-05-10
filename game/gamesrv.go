package game

import (
	"log"

	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/cluster"
	"github.com/fengqk/mars-base/common"
	"github.com/fengqk/mars-base/db"
	"github.com/fengqk/mars-base/network"
	"github.com/fengqk/mars-base/rpc"
	"github.com/fengqk/mars-server/message"
	"github.com/golang/protobuf/proto"
)

type (
	IServerMgr interface {
		Init() bool
		InitDB() bool
		GetServer() *network.ServerSocket
	}

	ServerMgr struct {
		service   *network.ServerSocket
		snowFlake *cluster.Snowflake
	}

	Config struct {
		common.Server    `yaml:"game"`
		common.Mysql     `yaml:"mysql"`
		common.Etcd      `yaml:"etcd"`
		common.SnowFlake `yaml:"snowflake"`
		common.Raft      `yaml:"raft"`
		common.Nats      `yaml:"nats"`
		common.Stub      `yaml:"stub"`
	}
)

var (
	CONF   Config
	SERVER ServerMgr
	RdID   int
)

func (s *ServerMgr) Init() bool {
	//初始配置文件
	base.ReadConf("mars.yaml", &CONF)

	ShowMessage := func() {
		base.LOG.Println("**********************************************************")
		base.LOG.Printf("\tGAME IP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
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

	//snowflake
	s.snowFlake = cluster.NewSnowflake(CONF.SnowFlake.Endpoints)

	packetEventProcess := new(EventProcess)
	packetEventProcess.Init()

	//本身game集群管理
	cluster.MGR.InitCluster(&common.ClusterInfo{Type: rpc.SERVICE_GAME, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)},
		CONF.Etcd.Endpoints, CONF.Nats.Endpoints, cluster.WithMailBoxEtcd(CONF.Raft.Endpoints), cluster.WithStubMailBoxEtcd(CONF.Raft.Endpoints, &CONF.Stub))
	cluster.MGR.BindPacketFunc(actor.MGR.PacketFunc)

	return true
}

func (s *ServerMgr) InitDB() bool {
	return db.OpenDB(CONF.Mysql) != nil
}

func (s *ServerMgr) GetServer() *network.ServerSocket {
	return s.service
}

func SendToClient(clusterId uint32, packet proto.Message) {
	pakcetHead := packet.(message.Packet).GetPacketHead()
	if pakcetHead != nil {
		cluster.MGR.SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_GATE, ClusterId: clusterId, Id: pakcetHead.Id}, "", proto.MessageName(packet), packet)
	}
}

func SendToZone(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_ZONE, SrcClusterId: cluster.MGR.Id()}
	cluster.MGR.SendMsg(head, funcName, params...)
}
