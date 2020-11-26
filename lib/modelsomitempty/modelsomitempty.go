package modelsomitempty

// Modeling the Messages received from proofpoint logs
// ProofpointMessage is top level object
type ProofpointMessage struct {
	Connection ConnectionObject `json:"connection,omitempty"`
	Envelope   EnvelopeObject   `json:"envelope,omitempty"`
	Filter     FilterObject     `json:"filter,omitempty"`
	Guid       string           `json:"guid,omitempty"`
	Msg        MsgObject        `json:"msg,omitempty"`
	MsgParts   []MsgPartsObject `json:"msgParts,omitempty"`
	Pps        PpsObject        `json:"pps,omitempty"`
	Ts         string           `json:"ts,omitempty"`
}
type MsgObject struct {
	Header           MsgHeaderObject        `json:"header,omitempty"`
	Lang             string                 `json:"lang,omitempty"`
	NormalizedHeader NormalizedHeaderObject `json:"normalizedHeader,omitempty"`
	ParsedAddresses  ParsedAddressesObject  `json:"parsedAddresses,omitempty"`
	SizeBytes        int                    `json:"sizeBytes,omitempty"`
}
type MsgHeaderObject struct {
	Cc         []string `json:"cc,omitempty"`
	From       []string `json:"from,omitempty"`
	MessageId  []string `json:"message-id,omitempty"`
	ReplyTo    []string `json:"reply-to,omitempty"`
	ReturnPath []string `json:"return-path,omitempty"`
	Subject    []string `json:"subject,omitempty"`
	To         []string `json:"to,omitempty"`
}
type PpsObject struct {
	Agent   string `json:"agent,omitempty"`
	Cid     string `json:"cid,omitempty"`
	Version string `json:"version,omitempty"`
}
type ConnectionObject struct {
	Sid           string    `json:"sid,omitempty"`
	Country       string    `json:"country,omitempty"`
	Helo          string    `json:"string,omitempty"`
	Host          string    `json:"host,omitempty"`
	Ip            string    `json:"ip,omitempty"`
	Protocol      string    `json:protocol,omitempty"`
	ResolveStatus string    `json:"resolveStatus,omitempty"`
	Tls           TlsObject `json:"tls,omitempty"`
}

type EnvelopeObject struct {
	From        string   `json:"from,omitempty"`
	FromHashed  string   `json:"fromhashed,omitempty"`
	Rcpts       []string `json:"rcpts,omitempty"`
	RcptsHashed []string `json:"rcptsHashed,omitempty"`
}

/*
Docs say there is smime but no evidence of it in the logs
- possibly because encryption not in heavy use???
*/
type FilterObject struct {
	Actions         []FilterActionsObject `json:"actions,omitempty"`
	Disposition     string                `json:"string,omitempty"`
	Pe              PeObject              `json:"pe,omitempty"`
	DurationSecs    float64               `json:"durationSecs,omitempty"`
	CurrentFolder   string                `json:"currentFolder,omitempty"`
	IsMsgEncrypted  bool                  `json:"isMsgEncrypted,omitempty"`
	IsMsgReinjected bool                  `json:"isMsgReinjected,omitempty"`
	Mid             int64                 `json:"mid,omitempty"`
	Modules         ModulesObject         `json:"modules,omitempty"`
	MsgSizeBytes    int64                 `json:"msgSizeBytes,omitempty"`
	OrigGuid        string                `json:origGuid,omitempty"`
	Qid             string                `json:qid,omitempty"`
	Quarantine      QuarantineObject      `json:"quarantine,omitempty"`
	Routes          []string              `json:"routes,omitempty"`
	RouteDirection  string                `json:routeDirection,omitempty"`
	StartTime       string                `json:"startTime,omitempty"`
	SubOrgs         SubOrgsObject         `json:"suborgs,omitempty"`
}

type FilterActionsObject struct {
	Module  string `json:"module,omitempty"`
	Rule    string `json:"rule,omitempty"`
	IsFinal bool   `json:"isFinal,omitempty"`
	Action  string `json:"action,omitempty"`
}
type SubOrgsObject struct {
	Sender string   `json:"sender,omitempty"`
	Rcpts  []string `json:"rcpts,omitempty"`
}
type ModulesObject struct {
	Dkimv      []DkimvObject    `json:"dkimv,omitempty"`
	Dmarc      DmarcObject      `json:"dmarc,omitempty"`
	Pdr        PdrObject        `json:"pdr,omitempty"`
	Spam       SpamObject       `json:"spam,omitempty"`
	Spf        SpfObject        `json:"spf,omitempty"`
	UrlDefense UrlDefenseObject `json:"urlDefense,omitempty"`
	ZeroHour   ZeroHourObject   `json:"zeroHour,omitempty"`
}
type SpfObject struct {
	Result string `json:"result,omitempty"`
	Domain string `json:"domain,omitempty"`
}
type ZeroHourObject struct {
	Score string `json:score:,omitempty"`
}
type UrlDefenseObject struct {
	Version UrlDefenseVersion `json:"version,omitempty"`
	Counts  CountsObject      `json:"counts,omitempty"`
}
type UrlDefenseVersion struct {
	Engine string `json:"string,omitempty"`
}

/*
Watch NoRewriteIsEmail
*/
type CountsObject struct {
	Total                        int `json:"total,omitempty"`
	Unique                       int `json:"unique,omitempty"`
	Rewritten                    int `json:"rewritten,omitempty"`
	NoRewriteIsEmail             int `json:"noRewriteIsEmail,omitempty"`
	NoRewriteIsExcludedDomain    int `json:"noRewriteIsExcludedDomain,omitempty"`
	NoRewriteIsUnsupportedScheme int `json:"noRewriteIsUnsupportedScheme,omitempty"`
	NoReWriteIsSchemeles         int `json:"noRewriteIsSchemeless,omitempty"`
	NoRewriteIsMaxLengthExceeded int `json:"noRewriteIsMaxLengthExceeded,omitempty"`
	NoRewriteIsContentTypeText   int `json:"noRewriteIsContentTypeText,omitempty"`
}
type SpamObject struct {
	Charsets []string      `json:"charsets,omitempty"`
	Langs    []string      `json:"langs,omitempty"`
	Scores   ScoresObject  `json:"scores,omitempty"`
	Version  VersionObject `json:"version,omitempty"`
}
type ScoresObject struct {
	Engine      int               `json:"engine,omitempty"`
	Overall     int               `json:"overall,omitempty"`
	Classifiers ClassifiersObject `json:"classifiers,omitempty"`
}

// These weren't in given docs but def in logs returned
type ClassifiersObject struct {
	Adult       int `json:"adult,omitempty"`
	Bulk        int `json:"bulk,omitempty"`
	Impostor    int `json:"impostor,omitempty"`
	LowPriority int `json:"lowPriority,omitempty"`
	Malware     int `json:"malware,omitempty"`
	Mlx         int `json:"mlx,omitempty"`
	Mlxlog      int `json:"mlxlog,omitempty"`
	Phish       int `json:"phish,omitempty"`
	Spam        int `json:"spam,omitempty"`
	Suspect     int `json:"suspect,omitempty"`
}

type VersionObject struct {
	Engine      string `json:"engine,omitempty"`
	Definitions string `json:"definitions,omitempty"`
}

/* Not positive this is in current version; i see no evidence of it
type SandboxObject struct {
		ErrorStatus string `json:"errorStatus,omitempty"`
}
*/
type PdrObject struct {
	V2 PdrV2Object `json:"v2,omitempty"`
}
type PdrV2Object struct {
	Response string `json:"response,omitempty"`
	Rscore   int    `json:"rscore,omitempty"`
}
type DkimvObject struct {
	Domain   string `json:"domain,omitempty"`
	Selector string `json:"selector,omitempty"`
	Result   string `json:"result,omitempty"`
}
type DmarcObject struct {
	// note "filterd" [sic] documentation
	FilteredResult string              `json:"filterdResult,omitempty"`
	AuthResults    []AuthResultsObject `json:"authResults,omitempty"`
	Records        DmarcRecordsObject  `json:"records,omitempty"`
	Srvid          string              `json:"srvid,omitempty"`
	Alignment      []AlignmentObject   `json:"alignment,omitempty"`
}
type DmarcRecordsObject struct {
	Query  string `json:"string,omitempty"`
	Record string `json:"record,omitempty"`
}
type AlignmentObject struct {
	FromDomain string                   `json:"fromDomain,omitempty"`
	Results    []AlignmentResultsObject `json:"results,omitempty"`
}
type AlignmentResultsObject struct {
	Result      string `json:"result,omitempty"`
	Method      string `json:"method,omitempty"`
	Identity    string `json:"identity,omitempty"`
	IdentityOrg string `json:"identityOrg,omitempty"`
}
type AuthResultsObject struct {
	EmailIdentities EmailIdentitiesObject `json:"emailIdentitities,omitempty"`
	Method          string                `json:"method,omitempty"`
	Propspec        PropspecObject        `json:"propspec,omitempty"`
	Reason          string                `json:"reason,omitempty"`
}
type PropspecObject struct {
	HeaderS string `json:"header.s,omitempty"`
	HeaderD string `json:"header.d,omitempty"`
}
type EmailIdentitiesObject struct {
	Header EIDHeaderObject `json:"header,omitempty"`
	Smtp   EIDSmtpObject   `json:"smtp,omitempty"`
}
type EIDHeaderObject struct {
	From string `json:"from,omitempty"`
}
type EIDSmtpObject struct {
	Helo     string `json:"helo,omitempty"`
	MailFrom string `json:"mailfrom,omitempty"`
}
type QuarantineObject struct {
	Folder string `json:"folder,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

// proofpoint encryption object
type PeObject struct {
	Rcpts []string `json:"rcpts,omitempty"`
}
type MessageObject struct {
	Header           HeaderObject           `json:"header,omitempty"`
	Lang             string                 `json:"lang,omitempty"`
	NormalizedHeader NormalizedHeaderObject `json:"normalizedHeader,omitempty"`
	ParsedAddresses  ParsedAddressesObject  `json:"parsedAddresses,omitempty"`
	SizeBytes        int                    `json:"sizeBytes,omitempty"`
}
type MsgPartsObject struct {
	DataBase64        string         `json:"dataBase64,omitempty"`
	DetectedCharset   string         `json:"detectedCharset,omitempty"`
	DetectedExt       string         `json:"detectedExt,omitempty"`
	DetectedMime      string         `json:"detectedMime,omitempty"`
	DetectedName      string         `json:"detectedName,omitempty"`
	DetectedSizeBytes int            `json:"detectedSizeBytes,omitempty"`
	Disposition       string         `json:"disposition,omitempty"`
	MD5               string         `json:"md5,omitempty"`
	SHA256            string         `json:"sha256,omitempty"`
	IsArchive         bool           `json:"isArchive,omitempty"`
	IsCorrupted       bool           `json:"isCorrupted,omitempty"`
	IsDeleted         bool           `json:"isDeleted,omitempty"`
	IsProtected       bool           `json:"isProtected,omitempty"`
	IsTimedOut        bool           `json:"isTimedOut,omitempty"`
	IsVirtual         bool           `json:"isVirtual,omitempty"`
	LabeledCharset    string         `json:"labeledCharset,omitempty"`
	LabeledExt        string         `json:"labeledExt,omitempty"`
	LabeledMime       string         `json:"labeledMime,omitempty"`
	LabeledName       string         `json:"labeledName,omitempty"`
	Metadata          MetadataObject `json:"metadata,omitempty"`
	SandboxStatus     string         `json:"sandboxStatus,omitempty"`
	SizeDecodedBytes  int            `json:"sizeDecodedBytes,omitempty"`
	StructureId       string         `json:"structureId,omitempty"`
	Urls              []UrlsObject   `json:"urls,omitempty"`
}

// Urls is an array of objects that occurs in MsgPartsObject
type UrlsObject struct {
	Url                string   `json:"url,omitempty"`
	IsRewritten        bool     `json:"isRewritten,omitempty"`
	NotRewrittenReason string   `json:notRewrittenReason,omitempty"`
	Src                []string `json:"src,omitempty"`
}
type MetadataObject struct {
	Cloud     string `json:"cloud,omitempty"`
	Viewport  string `json:"viewport,omitempty"`
	Messageid string `json:"messageid,omitempty"`
	Title     string `json:"title,omitempty"`
	Generator string `json:"generator,omitempty"`
}
type ParsedAddressesObject struct {
	Cc   []string `json:"cc,omitempty"`
	From []string `json:"from,omitempty"`
	To   []string `json:"to,omitempty"`
}
type HeaderObject struct {
	Cc         []string `json:"cc,omitempty"`
	From       []string `json:"from,omitempty"`
	MessageId  []string `json:"message-id,omitempty"`
	ReplyTo    []string `json:"reply-to,omitempty"`
	ReturnPath []string `json:"return-path,omitempty"`
	Subject    []string `json:"subject,omitempty"`
}
type NormalizedHeaderObject struct {
	Cc         []string `json:"cc,omitempty"`
	From       []string `json:"from,omitempty"`
	MessageId  []string `json:"message-id,omitempty"`
	ReplyTo    []string `json:"reply-to,omitempty"`
	ReturnPath []string `json:"return-path,omitempty"`
	Subject    []string `json:"subject,omitempty"`
}
type TlsObject struct {
	Inbound InboundObject `json:"inbound,omitempty"`
}
type InboundObject struct {
	Cipher     string `json:"cipher,omitempty"`
	CipherBits int    `json:"cipherbits,omitempty"`
	Policy     string `json:"policy,omitempty"`
	Version    string `json:"version,omitempty"`
}
