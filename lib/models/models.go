package models

// Modeling the Messages received from proofpoint POD logs; May 2019 Revsion B
// ProofpointMessage is the top level object
type ProofpointMessage struct {
	Connection ConnectionObject `json:"connection"`
	Envelope   EnvelopeObject   `json:"envelope"`
	Filter     FilterObject     `json:"filter"`
	Guid       string           `json:"guid"`
	Msg        MsgObject        `json:"msg"`
	MsgParts   []MsgPartsObject `json:"msgParts"`
	Pps        PpsObject        `json:"pps"`
	Ts         string           `json:"ts"`
}
type MsgObject struct {
	Header           MsgHeaderObject        `json:"header"`
	Lang             string                 `json:"lang"`
	NormalizedHeader NormalizedHeaderObject `json:"normalizedHeader"`
	ParsedAddresses  ParsedAddressesObject  `json:"parsedAddresses"`
	SizeBytes        int                    `json:"sizeBytes"`
}
type MsgHeaderObject struct {
	Cc         []string `json:"cc"`
	From       []string `json:"from"`
	MessageId  []string `json:"message-id"`
	ReplyTo    []string `json:"reply-to"`
	ReturnPath []string `json:"return-path"`
	Subject    []string `json:"subject"`
	To         []string `json:"to"`
}
type PpsObject struct {
	Agent   string `json:"agent"`
	Cid     string `json:"cid"`
	Version string `json:"version"`
}
type ConnectionObject struct {
	Sid           string    `json:"sid"`
	Country       string    `json:"country"`
	Helo          string    `json:"helo"`
	Host          string    `json:"host"`
	Ip            string    `json:"ip"`
	Protocol      string    `json:protocol"`
	ResolveStatus string    `json:"resolveStatus"`
	Tls           TlsObject `json:"tls"`
}

type EnvelopeObject struct {
	From        string   `json:"from"`
	FromHashed  string   `json:"fromhashed"`
	Rcpts       []string `json:"rcpts"`
	RcptsHashed []string `json:"rcptsHashed"`
}

/**
Docs say there is smime but no evidence of it in the logs
- probably should add it regardless
*/
type FilterObject struct {
	Actions         []FilterActionsObject `json:"actions"`
	Disposition     string                `json:"disposition"`
	Pe              PeObject              `json:"pe"`
	DurationSecs    float64               `json:"durationSecs"`
	CurrentFolder   string                `json:"currentFolder"`
	IsMsgEncrypted  bool                  `json:"isMsgEncrypted"`
	IsMsgReinjected bool                  `json:"isMsgReinjected"`
	Mid             int64                 `json:"mid"`
	Modules         ModulesObject         `json:"modules"`
	MsgSizeBytes    int64                 `json:"msgSizeBytes"`
	OrigGuid        string                `json:origGuid"`
	Qid             string                `json:qid"`
	Quarantine      QuarantineObject      `json:"quarantine"`
	Routes          []string              `json:"routes"`
	RouteDirection  string                `json:routeDirection"`
	StartTime       string                `json:"startTime"`
	SubOrgs         SubOrgsObject         `json:"suborgs"`
}

type FilterActionsObject struct {
	Module  string `json:"module"`
	Rule    string `json:"rule"`
	IsFinal bool   `json:"isFinal"`
	Action  string `json:"action"`
}
type SubOrgsObject struct {
	Sender string   `json:"sender"`
	Rcpts  []string `json:"rcpts"`
}

type ModulesObject struct {
	Dkimv      []DkimvObject    `json:"dkimv"`
	Dmarc      DmarcObject      `json:"dmarc"`
	Pdr        PdrObject        `json:"pdr"`
	Spam       SpamObject       `json:"spam"`
	Spf        SpfObject        `json:"spf"`
	UrlDefense UrlDefenseObject `json:"urlDefense"`
	ZeroHour   ZeroHourObject   `json:"zeroHour"`
}
type SpfObject struct {
	Result string `json:"result"`
	Domain string `json:"domain"`
}
type ZeroHourObject struct {
	Score string `json:score:"`
}
type UrlDefenseObject struct {
	Version UrlDefenseVersion `json:"version"`
	Counts  CountsObject      `json:"counts"`
}
type UrlDefenseVersion struct {
	Engine string `json:"Engine"`
}

// Need to conform NoRewriteIsEmail;
type CountsObject struct {
	Total                        int `json:"total"`
	Unique                       int `json:"unique"`
	Rewritten                    int `json:"rewritten"`
	NoRewriteIsEmail             int `json:"noRewriteIsEmail"`
	NoRewriteIsExcludedDomain    int `json:"noRewriteIsExcludedDomain"`
	NoRewriteIsUnsupportedScheme int `json:"noRewriteIsUnsupportedScheme"`
	NoReWriteIsSchemeles         int `json:"noRewriteIsSchemeless"`
	NoRewriteIsMaxLengthExceeded int `json:"noRewriteIsMaxLengthExceeded"`
	NoRewriteIsContentTypeText   int `json:"noRewriteIsContentTypeText"`
}
type SpamObject struct {
	Charsets []string      `json:"charsets"`
	Langs    []string      `json:"langs"`
	Scores   ScoresObject  `json:"scores"`
	Version  VersionObject `json:"version"`
}
type ScoresObject struct {
	Engine      int               `json:"engine"`
	Overall     int               `json:"overall"`
	Classifiers ClassifiersObject `json:"classifiers"`
}

// These werent't in the given doc but they are def in the logs
type ClassifiersObject struct {
	Adult       int `json:"adult"`
	Bulk        int `json:"bulk"`
	Impostor    int `json:"impostor"`
	LowPriority int `json:"lowPriority"`
	Malware     int `json:"malware"`
	Mlx         int `json:"mlx"`
	Mlxlog      int `json:"mlxlog"`
	Phish       int `json:"phish"`
	Spam        int `json:"spam"`
	Suspect     int `json:"suspect"`
}

type VersionObject struct {
	Engine      string `json:"engine"`
	Definitions string `json:"definitions"`
}

/* Not positive this is in current version; i see no evidence of it
type SandboxObject struct {
		ErrorStatus string `json:"errorStatus"`
}
*/
type PdrObject struct {
	V2 PdrV2Object `json:"v2"`
}
type PdrV2Object struct {
	Response string `json:"response"`
	Rscore   int    `json:"rscore"`
}
type DkimvObject struct {
	Domain   string `json:"domain"`
	Selector string `json:"selector"`
	Result   string `json:"result"`
}
type DmarcObject struct {
	// note "filterd" [sic] documentation
	FilteredResult string              `json:"filterdResult"`
	AuthResults    []AuthResultsObject `json:"authResults"`
	Records        DmarcRecordsObject  `json:"records"`
	Srvid          string              `json:"srvid"`
	Alignment      []AlignmentObject   `json:"alignment"`
}
type DmarcRecordsObject struct {
	Query  string `json:"query"`
	Record string `json:"record"`
}
type AlignmentObject struct {
	FromDomain string                   `json:"fromDomain"`
	Results    []AlignmentResultsObject `json:"results"`
}
type AlignmentResultsObject struct {
	Result      string `json:"result"`
	Method      string `json:"method"`
	Identity    string `json:"identity"`
	IdentityOrg string `json:"identityOrg"`
}
type AuthResultsObject struct {
	EmailIdentities EmailIdentitiesObject `json:"emailIdentitities"`
	Method          string                `json:"method"`
	Propspec        PropspecObject        `json:"propspec"`
	Reason          string                `json:"reason"`
}

// this is screwd up, the actual fieldname has period in it which is annoying
type PropspecObject struct {
	HeaderS string `json:"header.s"`
	HeaderD string `json:"header.d"`
}
type EmailIdentitiesObject struct {
	Header EIDHeaderObject `json:"header"`
	Smtp   EIDSmtpObject   `json:"smtp"`
}
type EIDHeaderObject struct {
	From string `json:"from"`
}
type EIDSmtpObject struct {
	Helo     string `json:"helo"`
	MailFrom string `json:"mailfrom"`
}
type QuarantineObject struct {
	Folder string `json:"folder"`
	Rule   string `json:"rule"`
}

// proofpoint encryption object
type PeObject struct {
	Rcpts []string `json:"rcpts"`
}
type MessageObject struct {
	Header           HeaderObject           `json:"header"`
	Lang             string                 `json:"lang"`
	NormalizedHeader NormalizedHeaderObject `json:"normalizedHeader"`
	ParsedAddresses  ParsedAddressesObject  `json:"parsedAddresses"`
	SizeBytes        int                    `json:"sizeBytes"`
}
type MsgPartsObject struct {
	DataBase64        string         `json:"dataBase64"`
	DetectedCharset   string         `json:"detectedCharset"`
	DetectedExt       string         `json:"detectedExt"`
	DetectedMime      string         `json:"detectedMime"`
	DetectedName      string         `json:"detectedName"`
	DetectedSizeBytes int            `json:"detectedSizeBytes"`
	Disposition       string         `json:"disposition"`
	MD5               string         `json:"md5"`
	SHA256            string         `json:"sha256"`
	IsArchive         bool           `json:"isArchive"`
	IsCorrupted       bool           `json:"isCorrupted"`
	IsDeleted         bool           `json:"isDeleted"`
	IsProtected       bool           `json:"isProtected"`
	IsTimedOut        bool           `json:"isTimedOut"`
	IsVirtual         bool           `json:"isVirtual"`
	LabeledCharset    string         `json:"labeledCharset"`
	LabeledExt        string         `json:"labeledExt"`
	LabeledMime       string         `json:"labeledMime"`
	LabeledName       string         `json:"labeledName"`
	Metadata          MetadataObject `json:"metadata"`
	SandboxStatus     string         `json:"sandboxStatus"`
	SizeDecodedBytes  int            `json:"sizeDecodedBytes"`
	StructureId       string         `json:"structureId"`
	Urls              []UrlsObject   `json:"urls"`
}

// Urls is an array of objects that occurs in MsgPartsObject; probably should be singular; whatever
type UrlsObject struct {
	Url                string   `json:"url"`
	IsRewritten        bool     `json:"isRewritten"`
	NotRewrittenReason string   `json:notRewrittenReason"`
	Src                []string `json:"src"`
}
type MetadataObject struct {
	Cloud     string `json:"cloud"`
	Viewport  string `json:"viewport"`
	Messageid string `json:"messageid"`
	Title     string `json:"title"`
	Generator string `json:"generator"`
}
type ParsedAddressesObject struct {
	Cc   []string `json:"cc"`
	From []string `json:"from"`
	To   []string `json:"to"`
}
type HeaderObject struct {
	Cc         []string `json:"cc"`
	From       []string `json:"from"`
	MessageId  []string `json:"message-id"`
	ReplyTo    []string `json:"reply-to"`
	ReturnPath []string `json:"return-path"`
	Subject    []string `json:"subject"`
}
type NormalizedHeaderObject struct {
	Cc         []string `json:"cc"`
	From       []string `json:"from"`
	MessageId  []string `json:"message-id"`
	ReplyTo    []string `json:"reply-to"`
	ReturnPath []string `json:"return-path"`
	Subject    []string `json:"subject"`
}
type TlsObject struct {
	Inbound InboundObject `json:"inbound"`
}
type InboundObject struct {
	Cipher     string `json:"cipher"`
	CipherBits int    `json:"cipherbits"`
	Policy     string `json:"policy"`
	Version    string `json:"version"`
}
