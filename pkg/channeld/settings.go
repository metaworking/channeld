package channeld

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/channeldorg/channeld/pkg/channeldpb"
	"github.com/channeldorg/channeld/pkg/common"
	"github.com/pkg/profile"
)

type GlobalSettingsType struct {
	Development   bool
	LogLevel      *NullableInt // zapcore.Level
	LogFile       *NullableString
	ProfileOption func(*profile.Profile)
	ProfilePath   string

	ServerNetwork              string
	ServerAddress              string
	ServerReadBufferSize       int
	ServerWriteBufferSize      int
	ServerFSM                  string
	ServerBypassAuth           bool
	ServerConnRecoverable      bool
	ServerConnRecoverTimeoutMs int64

	ClientNetworkWaitMasterServer bool
	ClientNetwork                 string
	ClientAddress                 string
	ClientReadBufferSize          int
	ClientWriteBufferSize         int
	ClientFSM                     string

	CompressionType channeldpb.CompressionType

	MaxConnectionIdBits uint8

	ConnectionAuthTimeoutMs int64
	MaxFailedAuthAttempts   int
	MaxFsmDisallowed        int

	SpatialControllerConfig NullableString
	SpatialChannelIdStart   common.ChannelId
	EntityChannelIdStart    common.ChannelId

	ChannelSettings map[channeldpb.ChannelType]ChannelSettingsType

	EnableRecordPacket bool

	ReplaySessionPersistenceDir string
}

type ACLSettingsType struct {
	Sub    ChannelAccessLevel
	Unsub  ChannelAccessLevel
	Remove ChannelAccessLevel
}

type ChannelSettingsType struct {
	TickIntervalMs                 uint
	DefaultFanOutIntervalMs        uint32
	DefaultFanOutDelayMs           int32
	RemoveChannelAfterOwnerRemoved bool
	// Should the channel send ChannelOwnerLost and ChannelOwnerRecovered message to its subscribers when the owner is lost and recovered?
	SendOwnerLostAndRecovered bool
	ACLSettings               ACLSettingsType
	// Optinal. The full name of the Protobuf message type for the channel data (including the package name)
	DataMsgFullName string
}

var GlobalSettings = GlobalSettingsType{
	LogLevel:                   &NullableInt{},
	LogFile:                    &NullableString{},
	ServerReadBufferSize:       0x0001ffff,
	ServerWriteBufferSize:      256,
	ServerFSM:                  "config/server_authoratative_fsm.json",
	ServerBypassAuth:           true,
	ServerConnRecoverable:      false,
	ServerConnRecoverTimeoutMs: 0,
	ClientReadBufferSize:       0x0001ffff,
	ClientWriteBufferSize:      512,
	ClientFSM:                  "config/client_non_authoratative_fsm.json",
	CompressionType:            channeldpb.CompressionType_NO_COMPRESSION,
	// Mirror uses int32 as the connId
	MaxConnectionIdBits:     31,
	ConnectionAuthTimeoutMs: 5000,
	MaxFailedAuthAttempts:   5,
	MaxFsmDisallowed:        10,
	SpatialChannelIdStart:   0x00010000,
	EntityChannelIdStart:    0x00080000,
	ChannelSettings: map[channeldpb.ChannelType]ChannelSettingsType{
		channeldpb.ChannelType_GLOBAL: {
			TickIntervalMs:                 10,
			DefaultFanOutIntervalMs:        20,
			DefaultFanOutDelayMs:           0,
			RemoveChannelAfterOwnerRemoved: false,
			SendOwnerLostAndRecovered:      false,
		},
	},
}

type NullableInt struct {
	Value    int
	HasValue bool
}

func (i NullableInt) String() string {
	if i.HasValue {
		return strconv.Itoa(i.Value)
	} else {
		return ""
	}
}

func (i *NullableInt) Set(s string) error {
	val, err := strconv.Atoi(s)
	if err == nil {
		i.Value = val
		i.HasValue = true
	}
	return err
}

type NullableString struct {
	Value    string
	HasValue bool
}

func (ns NullableString) String() string {
	return ns.Value
}

func (ns *NullableString) Set(s string) error {
	ns.Value = s
	ns.HasValue = true
	return nil
}

func (s *GlobalSettingsType) ParseFlag() error {
	flag.BoolVar(&s.Development, "dev", false, "run in development mode?")
	flag.Var(s.LogLevel, "loglevel", "the log level, -1 = Debug, 0 = Info, 1= Warn, 2 = Error, 3 = Panic")
	//flag.Var(stringPtrFlag{s.LogFile, fmt.Sprintf("logs/%s.log", time.Now().Format("20060102150405"))}, "logfile", "file path to store the log")
	flag.Var(s.LogFile, "logfile", "file path to store the log")
	flag.Func("profile", "available options: cpu, mem, goroutine", func(str string) error {
		switch strings.ToLower(str) {
		case "cpu":
			s.ProfileOption = profile.CPUProfile
		case "mem":
			s.ProfileOption = profile.MemProfile
		case "goroutine":
			s.ProfileOption = profile.GoroutineProfile
		default:
			return fmt.Errorf("invalid profile type: %s", str)
		}
		return nil
	})
	flag.StringVar(&s.ProfilePath, "profilepath", "profiles", "the path to store the profile output files")

	flag.StringVar(&s.ServerNetwork, "sn", "tcp", "the network type for the server connections")
	flag.StringVar(&s.ServerAddress, "sa", ":11288", "the network address for the server connections")
	flag.IntVar(&s.ServerReadBufferSize, "srb", s.ServerReadBufferSize, "the read buffer size for the server connections")
	flag.IntVar(&s.ServerWriteBufferSize, "swb", s.ServerWriteBufferSize, "the write buffer size for the server connections")
	flag.StringVar(&s.ServerFSM, "sfsm", s.ServerFSM, "the path to the server FSM config")
	flag.BoolVar(&s.ServerBypassAuth, "sba", s.ServerBypassAuth, "should server bypasses the authentication?")
	flag.BoolVar(&s.ServerConnRecoverable, "scr", s.ServerConnRecoverable, "is the server connection recoverable?")
	flag.Int64Var(&s.ServerConnRecoverTimeoutMs, "scrt", s.ServerConnRecoverTimeoutMs, "the duration to wait for the server connection to recover. Default is 0 (infinite)")

	flag.BoolVar(&s.ClientNetworkWaitMasterServer, "cwm", true, "should the client network starts listening after the Global channel being possessed by the Master Server?")
	flag.StringVar(&s.ClientNetwork, "cn", "tcp", "the network type for the client connections")
	flag.StringVar(&s.ClientAddress, "ca", ":12108", "the network address for the client connections")
	flag.IntVar(&s.ClientReadBufferSize, "crb", s.ClientReadBufferSize, "the read buffer size for the client connections")
	flag.IntVar(&s.ClientWriteBufferSize, "cwb", s.ClientWriteBufferSize, "the write buffer size for the client connections")
	flag.StringVar(&s.ClientFSM, "cfsm", s.ClientFSM, "the path to the client FSM config")

	flag.BoolVar(&s.EnableRecordPacket, "erp", false, "enable record message packets send from clients")
	flag.StringVar(&s.ReplaySessionPersistenceDir, "rspd", "", "the path to write packet recording")

	// Use flag.Uint instead of flag.UintVar to avoid the default value being overwritten by the flag value
	ct := flag.Uint("ct", 0, "the compression type, 0 = No, 1 = Snappy")
	flag.Var(&s.SpatialControllerConfig, "scc", "the path to the spatial controller config file")
	scs := flag.Uint("scs", uint(s.SpatialChannelIdStart), "start ChannelId of spatial channels. Default is 0x00010000.")
	ecs := flag.Uint("ecs", uint(s.EntityChannelIdStart), "start ChannelId of entity channels. Default is 0x00080000.")
	mcb := flag.Uint("mcb", uint(s.MaxConnectionIdBits), "max bits of ConnectionId (e.g. 16 means max ConnectionId = 1<<16 - 1). Up to 32.")
	cat := flag.Uint("cat", uint(s.ConnectionAuthTimeoutMs), "the duration to allow a connection stay unauthenticated before closing it. Default is 5000. (0 = no limit)")
	mfaa := flag.Int("mfaa", s.MaxFailedAuthAttempts, "the max number of failed authentication attempts before closing the connection. Default is 5. (0 = no limit)")
	mfd := flag.Int("mfd", s.MaxFsmDisallowed, "the max number of disallowed FSM transitions before closing the connection. Default is 10. (0 = no limit)")

	chs := flag.String("chs", "config/channel_settings_hifi.json", "the path to the channel settings file")

	flag.Parse()

	if ct != nil {
		s.CompressionType = channeldpb.CompressionType(*ct)
	}

	if scs != nil {
		s.SpatialChannelIdStart = common.ChannelId(*scs)
	}

	if ecs != nil {
		s.EntityChannelIdStart = common.ChannelId(*ecs)
	}

	if mcb != nil {
		s.MaxConnectionIdBits = uint8(*mcb)
	}

	if cat != nil {
		s.ConnectionAuthTimeoutMs = int64(*cat)
	}

	if mfaa != nil {
		s.MaxFailedAuthAttempts = int(*mfaa)
	}

	if mfd != nil {
		s.MaxFsmDisallowed = int(*mfd)
	}

	chsData, err := os.ReadFile(*chs)
	if err == nil {
		if err := json.Unmarshal(chsData, &GlobalSettings.ChannelSettings); err != nil {
			return fmt.Errorf("failed to unmarshall channel settings: %v", err)
		}
	} else {
		return fmt.Errorf("failed to read channel settings: %v", err)
	}

	return nil
}

func (s GlobalSettingsType) GetChannelSettings(t channeldpb.ChannelType) ChannelSettingsType {
	settings, exists := s.ChannelSettings[t]
	if !exists {
		settings = s.ChannelSettings[channeldpb.ChannelType_GLOBAL]
	}
	return settings
}
