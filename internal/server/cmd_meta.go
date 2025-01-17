package server

import (
	"github.com/dicedb/dice/internal/cmd"
	"github.com/dicedb/dice/internal/comm"
	"github.com/dicedb/dice/internal/eval"
	"github.com/dicedb/dice/internal/shard"
)

// CmdType defines the type of DiceDB command based on how it interacts with shards.
// It uses an integer value to represent different command types.
type CmdType int

// Enum values for CmdType using iota for auto-increment.
// Global commands don't interact with shards, SingleShard commands interact with one shard,
// MultiShard commands interact with multiple shards, and Custom commands require a direct client connection.
const (
	Global      CmdType = iota // Global commands don't need to interact with shards.
	SingleShard                // Single-shard commands interact with only one shard.
	MultiShard                 // MultiShard commands interact with multiple shards using scatter-gather logic.
	Custom                     // Custom commands involve direct client communication.
)

// CmdsMeta stores metadata about DiceDB commands, including how they are processed across shards.
// CmdType indicates how the command should be handled, while Breakup and Gather provide logic
// for breaking up multishard commands and gathering their responses.
type CmdsMeta struct {
	Cmd          string                                                                                  // Command name.
	Breakup      func(mgr *shard.ShardManager, DiceDBCmd *cmd.DiceDBCmd, c *comm.Client) []cmd.DiceDBCmd // Function to break up multishard commands.
	Gather       func(responses ...eval.EvalResponse) []byte                                             // Function to gather responses from shards.
	RespNoShards func(args []string) []byte                                                              // Function for commands that don't interact with shards.
	CmdType                                                                                              // Enum indicating the command type.
}

// WorkerCmdsMeta is a map that associates command names with their corresponding metadata.
var (
	WorkerCmdsMeta = map[string]CmdsMeta{}

	// Metadata for global commands that don't interact with shards.
	// PING is an example of global command.
	pingCmdMeta = CmdsMeta{
		Cmd:          "PING",
		CmdType:      Global,
		RespNoShards: eval.RespPING,
	}

	// Metadata for single-shard commands that only interact with one shard.
	// These commands don't require breakup and gather logic.
	setCmdMeta = CmdsMeta{
		Cmd:     "SET",
		CmdType: SingleShard,
	}
	getCmdMeta = CmdsMeta{
		Cmd:     "GET",
		CmdType: SingleShard,
	}
	getsetCmdMeta = CmdsMeta{
		Cmd:     "GETSET",
		CmdType: SingleShard,
	}
	setexCmdMeta = CmdsMeta{
		Cmd:     "SETEX",
		CmdType: SingleShard,
	}
	jsonArrAppendCmdMeta = CmdsMeta{
		Cmd:     "JSON.ARRAPPEND",
		CmdType: SingleShard,
	}
	jsonArrLenCmdMeta = CmdsMeta{
		Cmd:     "JSON.ARRLEN",
		CmdType: SingleShard,
	}
	jsonArrPopCmdMeta = CmdsMeta{
		Cmd:     "JSON.ARRPOP",
		CmdType: SingleShard,
	}
	getrangeCmdMeta = CmdsMeta{
		Cmd:     "GETRANGE",
		CmdType: SingleShard,
	}
	zaddCmdMeta = CmdsMeta{
		Cmd:     "ZADD",
		CmdType: SingleShard,
	}
	zcountCmdMeta = CmdsMeta{
		Cmd:     "ZCOUNT",
		CmdType: SingleShard,
	}
	zrangeCmdMeta = CmdsMeta{
		Cmd:     "ZRANGE",
		CmdType: SingleShard,
	}
	appendCmdMeta = CmdsMeta{
		Cmd:     "APPEND",
		CmdType: SingleShard,
	}
	zpopminCmdMeta = CmdsMeta{
		Cmd:     "ZPOPMIN",
		CmdType: SingleShard,
	}
	zrankCmdMeta = CmdsMeta{
		Cmd:     "ZRANK",
		CmdType: SingleShard,
	}
	pfaddCmdMeta = CmdsMeta{
		Cmd:     "PFADD",
		CmdType: SingleShard,
	}
	pfcountCmdMeta = CmdsMeta{
		Cmd:     "PFCOUNT",
		CmdType: SingleShard,
	}
	pfmergeCmdMeta = CmdsMeta{
		Cmd:     "PFMERGE",
		CmdType: SingleShard,
	}

	jsonclearCmdMeta = CmdsMeta{
		Cmd:     "JSON.CLEAR",
		CmdType: SingleShard,
	}

	jsonstrlenCmdMeta = CmdsMeta{
		Cmd:     "JSON.STRLEN",
		CmdType: SingleShard,
	}

	jsonobjlenCmdMeta = CmdsMeta{
		Cmd:     "JSON.OBJLEN",
		CmdType: SingleShard,
	}
	hlenCmdMeta = CmdsMeta{
		Cmd:     "HLEN",
		CmdType: SingleShard,
	}
	hstrlenCmdMeta = CmdsMeta{
		Cmd:     "HSTRLEN",
		CmdType: SingleShard,
	}
	hscanCmdMeta = CmdsMeta{
		Cmd:     "HSCAN",
		CmdType: SingleShard,
	}

	jsonarrinsertCmdMeta = CmdsMeta{
		Cmd:     "JSON.ARRINSERT",
		CmdType: SingleShard,
	}

	jsonarrtrimCmdMeta = CmdsMeta{
		Cmd:     "JSON.ARRTRIM",
		CmdType: SingleShard,
	}

	jsonobjkeystCmdMeta = CmdsMeta{
		Cmd:     "JSON.OBJKEYS",
		CmdType: SingleShard,
	}

	incrCmdMeta = CmdsMeta{
		Cmd:     "INCR",
		CmdType: SingleShard,
	}
	incrByCmdMeta = CmdsMeta{
		Cmd:     "INCRBY",
		CmdType: SingleShard,
	}
	decrCmdMeta = CmdsMeta{
		Cmd:     "DECR",
		CmdType: SingleShard,
	}
	decrByCmdMeta = CmdsMeta{
		Cmd:     "DECRBY",
		CmdType: SingleShard,
	}
	incrByFloatCmdMeta = CmdsMeta{
		Cmd:     "INCRBYFLOAT",
		CmdType: SingleShard,
	}
	hincrbyCmdMeta = CmdsMeta{
		Cmd:     "HINCRBY",
		CmdType: SingleShard,
	}
	hincrbyfloatCmdMeta = CmdsMeta{
		Cmd:     "HINCRBYFLOAT",
		CmdType: SingleShard,
	}
	hrandfieldCmdMeta = CmdsMeta{
		Cmd:     "HRANDFIELD",
		CmdType: SingleShard,
	}
	zpopmaxCmdMeta = CmdsMeta{
		Cmd:     "ZPOPMAX",
		CmdType: SingleShard,
	}
	bfaddCmdMeta = CmdsMeta{
		Cmd:     "BF.ADD",
		CmdType: SingleShard,
	}
	bfreserveCmdMeta = CmdsMeta{
		Cmd:     "BF.RESERVE",
		CmdType: SingleShard,
	}
	bfexistsCmdMeta = CmdsMeta{
		Cmd:     "BF.EXISTS",
		CmdType: SingleShard,
	}
	bfinfoCmdMeta = CmdsMeta{
		Cmd:     "BF.INFO",
		CmdType: SingleShard,
	}

	cmsInitByDimCmdMeta = CmdsMeta{
		Cmd:     "CMS.INITBYDIM",
		CmdType: SingleShard,
	}

	cmsInitByProbCmdMeta = CmdsMeta{
		Cmd:     "CMS.INITBYPROB",
		CmdType: SingleShard,
	}

	cmsInfoCmdMeta = CmdsMeta{
		Cmd:     "CMS.INFO",
		CmdType: SingleShard,
	}

	cmsIncrByCmdMeta = CmdsMeta{
		Cmd:     "CMS.INCRBY",
		CmdType: SingleShard,
	}

	cmsQueryCmdMeta = CmdsMeta{
		Cmd:     "CMS.QUERY",
		CmdType: SingleShard,
	}

	cmsMergeCmdMeta = CmdsMeta{
		Cmd:     "CMS.MERGE",
		CmdType: SingleShard,
	}

	// Metadata for multishard commands would go here.
	// These commands require both breakup and gather logic.

	// Metadata for custom commands requiring specific client-side logic would go here.
)

// init initializes the WorkerCmdsMeta map by associating each command name with its corresponding metadata.
func init() {
	// Global commands.
	WorkerCmdsMeta["PING"] = pingCmdMeta

	// Single-shard commands.
	WorkerCmdsMeta["SET"] = setCmdMeta
	WorkerCmdsMeta["GET"] = getCmdMeta
	WorkerCmdsMeta["GETSET"] = getsetCmdMeta
	WorkerCmdsMeta["SETEX"] = setexCmdMeta
	WorkerCmdsMeta["JSON.ARRAPPEND"] = jsonArrAppendCmdMeta
	WorkerCmdsMeta["JSON.ARRLEN"] = jsonArrLenCmdMeta
	WorkerCmdsMeta["JSON.ARRPOP"] = jsonArrPopCmdMeta
	WorkerCmdsMeta["GETRANGE"] = getrangeCmdMeta
	WorkerCmdsMeta["APPEND"] = appendCmdMeta
	WorkerCmdsMeta["JSON.CLEAR"] = jsonclearCmdMeta
	WorkerCmdsMeta["JSON.STRLEN"] = jsonstrlenCmdMeta
	WorkerCmdsMeta["JSON.OBJLEN"] = jsonobjlenCmdMeta
	WorkerCmdsMeta["JSON.ARRINSERT"] = jsonarrinsertCmdMeta
	WorkerCmdsMeta["JSON.ARRTRIM"] = jsonarrtrimCmdMeta
	WorkerCmdsMeta["JSON.OBJKEYS"] = jsonobjkeystCmdMeta
	WorkerCmdsMeta["ZADD"] = zaddCmdMeta
	WorkerCmdsMeta["ZCOUNT"] = zcountCmdMeta
	WorkerCmdsMeta["ZRANGE"] = zrangeCmdMeta
	WorkerCmdsMeta["ZRANK"] = zrankCmdMeta
	WorkerCmdsMeta["PFADD"] = pfaddCmdMeta
	WorkerCmdsMeta["ZPOPMIN"] = zpopminCmdMeta
	WorkerCmdsMeta["PFCOUNT"] = pfcountCmdMeta
	WorkerCmdsMeta["PFMERGE"] = pfmergeCmdMeta
	WorkerCmdsMeta["HLEN"] = hlenCmdMeta
	WorkerCmdsMeta["HSTRLEN"] = hstrlenCmdMeta
	WorkerCmdsMeta["HSCAN"] = hscanCmdMeta
	WorkerCmdsMeta["INCR"] = incrCmdMeta
	WorkerCmdsMeta["INCRBY"] = incrByCmdMeta
	WorkerCmdsMeta["INCR"] = incrCmdMeta
	WorkerCmdsMeta["DECR"] = decrCmdMeta
	WorkerCmdsMeta["DECRBY"] = decrByCmdMeta
	WorkerCmdsMeta["INCRBYFLOAT"] = incrByFloatCmdMeta
	WorkerCmdsMeta["HINCRBY"] = hincrbyCmdMeta
	WorkerCmdsMeta["HINCRBYFLOAT"] = hincrbyfloatCmdMeta
	WorkerCmdsMeta["HRANDFIELD"] = hrandfieldCmdMeta
	WorkerCmdsMeta["PFADD"] = pfaddCmdMeta
	WorkerCmdsMeta["ZPOPMIN"] = zpopminCmdMeta
	WorkerCmdsMeta["PFCOUNT"] = pfcountCmdMeta
	WorkerCmdsMeta["PFMERGE"] = pfmergeCmdMeta
	WorkerCmdsMeta["HINCRBY"] = hincrbyCmdMeta
	WorkerCmdsMeta["HINCRBYFLOAT"] = hincrbyfloatCmdMeta
	WorkerCmdsMeta["HRANDFIELD"] = hrandfieldCmdMeta
	WorkerCmdsMeta["PFADD"] = pfaddCmdMeta
	WorkerCmdsMeta["PFCOUNT"] = pfcountCmdMeta
	WorkerCmdsMeta["PFMERGE"] = pfmergeCmdMeta
	WorkerCmdsMeta["HINCRBY"] = hincrbyCmdMeta
	WorkerCmdsMeta["HINCRBYFLOAT"] = hincrbyfloatCmdMeta
	WorkerCmdsMeta["HRANDFIELD"] = hrandfieldCmdMeta
	WorkerCmdsMeta["ZPOPMAX"] = zpopmaxCmdMeta
	WorkerCmdsMeta["BF.ADD"] = bfaddCmdMeta
	WorkerCmdsMeta["BF.RESERVE"] = bfreserveCmdMeta
	WorkerCmdsMeta["BF.EXISTS"] = bfexistsCmdMeta
	WorkerCmdsMeta["BF.INFO"] = bfinfoCmdMeta
	WorkerCmdsMeta["CMS.INITBYDIM"] = cmsInitByDimCmdMeta
	WorkerCmdsMeta["CMS.INITBYPROB"] = cmsInitByProbCmdMeta
	WorkerCmdsMeta["CMS.INFO"] = cmsInfoCmdMeta
	WorkerCmdsMeta["CMS.INCRBY"] = cmsIncrByCmdMeta
	WorkerCmdsMeta["CMS.QUERY"] = cmsQueryCmdMeta
	WorkerCmdsMeta["CMS.MERGE"] = cmsMergeCmdMeta
	// Additional commands (multishard, custom) can be added here as needed.
}
