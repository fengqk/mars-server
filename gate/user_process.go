package gate

import (
	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/network"
	"github.com/fengqk/mars-base/rpc"
	"github.com/fengqk/mars-server/message"
)

// 消息防火墙
var (
	s_clientCheckFilters    map[string]bool //use for no check playerid
	s_clientCheckFilterInit bool
)

type (
	IUserProcess interface {
		actor.IActor

		CheckClientEx(uint32, string, rpc.RpcHead) bool
		CheckClient(uint32, string, rpc.RpcHead) *Player
		SwitchSendToGame(uint32, string, rpc.RpcHead, []byte)
		SwitchSendToZone(uint32, string, rpc.RpcHead, []byte)

		addKey(uint32, *base.Dh)
		delKey(uint32)
	}

	UserProcess struct {
		actor.Actor
		keyMap map[uint32]*base.Dh
	}
)

func IsCheckClient(msg string) bool {
	if !s_clientCheckFilterInit {
		s_clientCheckFilters = make(map[string]bool)
		s_clientCheckFilters["LoginAccountRequest"] = true
		s_clientCheckFilters["CreatePlayerRequest"] = true
		s_clientCheckFilters["LoginPlayerRequset"] = true
		s_clientCheckFilterInit = true
	}

	_, exist := s_clientCheckFilters[msg]
	return exist
}

func (u *UserProcess) Init() {
	u.Actor.Init()
	u.keyMap = map[uint32]*base.Dh{}
	actor.MGR.RegisterActor(u)
	u.Actor.Start()
}

func (u *UserProcess) PacketFunc(pkg rpc.Packet) bool {
	buff := pkg.Buff
	socketid := pkg.Id
	packetId, data := message.Decode(buff)
	packetRoute := message.GetPakcetRoute(packetId)
	if packetRoute == nil {
		//客户端主动断开
		if packetId == network.DISCONNECTINT {
			stream := base.NewBitStream(buff, len(buff))
			stream.ReadInt(32)
			SERVER.GetPlayerMgr().SendMsg(rpc.RpcHead{}, "DEL_ACCOUNT", uint32(stream.ReadInt(32)))
			u.SendMsg(rpc.RpcHead{}, "DISCONNECT", socketid)
		} else if packetId == network.HEART_PACKET {
			//心跳netsocket做处理，这里不处理
		} else {
			base.LOG.Printf("包解析错误1  socket=%d", socketid)
		}
		return true
	}

	//获取配置的路由地址
	packet := packetRoute.Func()
	err := message.UnmarshalText(packet, data)
	if err != nil {
		base.LOG.Printf("包解析错误2  socket=%d", socketid)
		return true
	}

	packetHead := packet.(message.Packet).GetPacketHead()
	if packetHead == nil || packetHead.Ckx != message.Default_Ipacket_Ckx || packetHead.Stx != message.Default_Ipacket_Stx {
		base.LOG.Printf("(A)致命的越界包,已经被忽略 socket=%d", socketid)
		return true
	}

	packetName := packetRoute.FuncName
	head := rpc.RpcHead{Id: packetHead.Id, SrcClusterId: SERVER.GetCluster().Id()}
	rpcPacket := rpc.Marshal(&head, &packetName, packet)
	//解析整个包
	if head.DestServerType == rpc.SERVICE_GAME {
		u.SwitchSendToGame(socketid, packetName, head, rpcPacket)
	} else if head.DestServerType == rpc.SERVICE_ZONE {
		u.SwitchSendToZone(socketid, packetName, head, rpcPacket)
	} else {
		actor.MGR.PacketFunc(rpc.Packet{Id: socketid, Buff: rpcPacket.Buff})
	}
	return true
}

func (u *UserProcess) CheckClient(sockId uint32, packetName string, head rpc.RpcHead) *Player {
	pPlayer := SERVER.GetPlayerMgr().GetPlayer(sockId)
	if pPlayer != nil && (pPlayer.PlayerId <= 0 || pPlayer.PlayerId != head.Id) {
		base.LOG.Fatalf("Old socket communication or viciousness[%d].", sockId)
		return nil
	}
	return pPlayer
}

func (u *UserProcess) CheckClientEx(sockId uint32, packetName string, head rpc.RpcHead) bool {
	if IsCheckClient(packetName) {
		return true
	}

	playerId := SERVER.GetPlayerMgr().GetPlayerId(sockId)
	if playerId <= 0 || playerId != head.Id {
		base.LOG.Fatalf("Old socket communication or viciousness[%d].", sockId)
		return false
	}
	return true
}

func (u *UserProcess) SwitchSendToGame(socketId uint32, packetName string, head rpc.RpcHead, packet rpc.Packet) {
	pPlayer := u.CheckClient(socketId, packetName, head)
	if pPlayer != nil {
		head.ClusterId = pPlayer.GClusterId
		head.DestServerType = rpc.SERVICE_GAME
		SERVER.GetCluster().Send(head, packet)
	}
}

func (u *UserProcess) SwitchSendToZone(socketId uint32, packetName string, head rpc.RpcHead, packet rpc.Packet) {
	pPlayer := u.CheckClient(socketId, packetName, head)
	if pPlayer != nil {
		head.ClusterId = pPlayer.ZClusterId
		head.DestServerType = rpc.SERVICE_ZONE
		SERVER.GetCluster().Send(head, packet)
	}
}

func (u *UserProcess) addKey(SocketId uint32, dh *base.Dh) {
	u.keyMap[SocketId] = dh
}

func (u *UserProcess) delKey(SocketId uint32) {
	delete(u.keyMap, SocketId)
}
