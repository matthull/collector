package state

import (
	"database/sql"
	"time"

	"github.com/pganalyze/collector/config"
)

// PersistedState - State thats kept across collector runs to be used for diffs
type PersistedState struct {
	CollectedAt time.Time

	// Databases we connected to and fetched local catalog data (e.g. schema)
	DatabaseOidsWithLocalCatalog []Oid

	StatementStats PostgresStatementStatsMap
	RelationStats  PostgresRelationStatsMap
	IndexStats     PostgresIndexStatsMap
	FunctionStats  PostgresFunctionStatsMap

	Roles     []PostgresRole
	Databases []PostgresDatabase
	Backends  []PostgresBackend
	Relations []PostgresRelation
	Settings  []PostgresSetting
	Functions []PostgresFunction
	Version   PostgresVersion
	Logs      []LogLine
	Explains  []PostgresExplain

	DataDirectory string
	System        SystemState

	CollectorStats CollectorStats
}

// TransientState - State thats only used within a collector run (and not needed for diffs)
type TransientState struct {
	Statements PostgresStatementMap
}

// DiffState - Result of diff-ing two persistent state structs
type DiffState struct {
	StatementStats DiffedPostgresStatementStatsMap
	RelationStats  DiffedPostgresRelationStatsMap
	IndexStats     DiffedPostgresIndexStatsMap
	FunctionStats  DiffedPostgresFunctionStatsMap

	SystemCPUStats     DiffedSystemCPUStatsMap
	SystemNetworkStats DiffedNetworkStatsMap
	SystemDiskStats    DiffedDiskStatsMap

	CollectorStats DiffedCollectorStats
}

// StateOnDiskFormatVersion - Increment this when an old state preserved to disk should be ignored
const StateOnDiskFormatVersion = 1

type StateOnDisk struct {
	FormatVersion uint

	PrevStateByAPIKey map[string]PersistedState
}

type CollectionOpts struct {
	CollectPostgresRelations bool
	CollectPostgresSettings  bool
	CollectPostgresLocks     bool
	CollectPostgresFunctions bool
	CollectPostgresBloat     bool
	CollectPostgresViews     bool

	CollectLogs              bool
	CollectExplain           bool
	CollectSystemInformation bool

	CollectorApplicationName string
	StatementTimeoutMs       int32 // Statement timeout for all SQL statements sent to the database

	DiffStatements bool

	SubmitCollectedData bool
	TestRun             bool

	StateFilename    string
	WriteStateUpdate bool
}

type GrantConfig struct {
	// Here be dragons
}

type Grant struct {
	Valid    bool
	Config   GrantConfig       `json:"config"`
	S3URL    string            `json:"s3_url"`
	S3Fields map[string]string `json:"s3_fields"`
}

type Server struct {
	Config           config.ServerConfig
	Connection       *sql.DB
	PrevState        PersistedState
	RequestedSslMode string
	Grant            Grant
}
