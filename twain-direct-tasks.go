package twain_direct

type task struct {
	Actions Action `json:"actions"`
}

type Action struct {
	Action string `json:"action"`
	// "configure", "encryptionProfiles",
	// "encryptionPublicKeys",
	// "encryptionReport"
	Comment   string `json:"comment"`
	Exception string `json:"exception"`
	// "fail", "ignore"
	Vendor               string                `json:"vendor"`
	Streams              []Stream              `json:"streams"`
	EncryptionProfiles   []EncryptionProfile   `json:"encryptionProfiles"`
	EncryptionPublicKeys []EncryptionPublicKey `json:"encryptionPublicKeys"`
	EncryptionReport     EncryptionReport      `json:"encryptionReport"`
}

type Stream struct {
	Comment   string `json:"comment"`
	Vendor    string `json:"vendor"`
	Exception string `json:"exception"`
	// "fail", "ignore", "nextStream"
	Name    string   `json:"name"`
	Sources []Source `json:"sources"`
}

type Source struct {
	Comment   string `json:"comment"`
	Vendor    string `json:"vendor"`
	Exception string `json:"exception"`
	// "fail", "ignore", "nextStream"
	Name   string `json:"name"`
	Source string `json:"source"`
	// default: "any"
	// "feeder", "feederFront", "feederRear",
	// "flatbed", "planetary", "storage"
	PixelFormats []PixelFormat `json:"pixelFormats"`
}

type PixelFormat struct {
	Comment   string `json:"comment"`
	Vendor    string `json:"vendor"`
	Exception string `json:"exception"`
	// "fail", "ignore", "nextStream"
	Name string `json:"name"`
	// default: "any"
	// "feeder", "feederFront", "feederRear",
	// "flatbed", "planetary", "storage"
	PixelFormat string      `json:"pixelFormat"`
	Attributes  []Attribute `json:"attributes"`
}

type Attribute struct {
	Comment   string  `json:"comment"`
	Vendor    string  `json:"vendor"`
	Exception string  `json:"exception"`
	Attribute string  `json:"attribute"`
	Values    []Value `json:"values"`
}

type Value struct {
	Comment   string `json:"comment"`
	Vendor    string `json:"vendor"`
	Exception string `json:"exception"`
	Value     string `json:"value"`
}

type EncryptionProfile struct {
	Comment   string `json:"comment"`
	Vendor    string `json:"vendor"`
	Exception string `json:"exception,omitempty"`
	Profile   string `json:"profile"`
}

type EncryptionPublicKey struct {
	Comment         string `json:"comment"`
	Vendor          string `json:"vendor"`
	Exception       string `json:"exception,omitempty"`
	Base64PublicKey string `json:"base64PublicKey,omitempty"`
	PublicKeyType   string `json:"publicKeyType,omitempty"`
	// default: "pem"
}

type EncryptionReport struct {
	DigitalSignatures    []DigitalSignature    `json:"digitalSignature,omitempty"`
	EncryptionProfiles   []EncryptionProfile   `json:"encryptionProfiles,omitempty"`
	EncryptionPublicKeys []EncryptionPublicKey `json:"encryptionPublicKeys,omitempty"`
}

type DigitalSignature struct {
	Comment          string `json:"comment"`
	Vendor           string `json:"vendor"`
	DigitalSignature string `json:"digitalSignature"`
}
