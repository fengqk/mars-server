package gate

import (
	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/cluster"
	"github.com/fengqk/mars-base/common"
	"github.com/fengqk/mars-base/network"
	"github.com/fengqk/mars-base/rpc"
	"github.com/fengqk/mars-server/gate/player"
)

type (
	IServerMgr interface {
		Init() bool
		GetServer() *network.ServerSocket
		GetCluster() *cluster.Cluster
		GetPlayerMgr() *player.PlayerMgr
		OnServerStart()
	}

	ServerMgr struct {
		service   *network.ServerSocket
		isInited  bool
		playerMgr *player.PlayerMgr
		cluster   *cluster.Cluster
	}

	Config struct {
		common.Server `yaml:"gate"`
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
	if s.isInited {
		return true
	}

	base.ReadConf("mars.yaml", &CONF)

	ShowMessage := func() {
		base.LOG.Println("**********************************************************")
		base.LOG.Printf("\tGATE IP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		base.LOG.Println("**********************************************************")
	}
	ShowMessage()

	s.service = new(network.ServerSocket)
	s.service.Init(CONF.Server.Ip, CONF.Server.Port)
	s.service.SetMaxPacketLen(base.MAX_CLIENT_PACKET)
	s.service.SetConnectType(network.CLIENT_CONNECT)

	packetUserProcess := new(UserProcess)
	packetUserProcess.Init()
	s.service.BindPacketFunc(packetUserProcess.PacketFunc)
	s.service.Start()

	packetEventProcess := new(EventProcess)
	packetEventProcess.Init()

	s.cluster = new(cluster.Cluster)
	s.cluster.InitCluster(&common.ClusterInfo{Type: rpc.SERVICE_GATE, Ip: CONF.Server.Ip, Port: CONF.Server.Port},
		CONF.Etcd.Endpoints, CONF.Nats.Endpoints, cluster.WithStubMailBoxEtcd(CONF.Raft.Endpoints, &CONF.Stub))
	s.cluster.BindPacketFunc(actor.MGR.PacketFunc)
	s.cluster.BindPacketFunc(DispatchPacket)

	s.playerMgr = new(player.PlayerMgr)
	s.playerMgr.Init()
	return false

}

func (s *ServerMgr) GetServer() *network.ServerSocket {
	return s.service
}

func (s *ServerMgr) GetCluster() *cluster.Cluster {
	return s.cluster
}

func (s *ServerMgr) GetPlayerMgr() *player.PlayerMgr {
	return s.playerMgr
}

func (s *ServerMgr) OnServerStart() {
	s.service.Start()
}
