package twain_direct

type Metadata struct {
	Status    Status    `json:"status"`
	Address   Address   `json:"address"`
	Image     Image     `json:"image"`
	Barcodes  []Barcode `json:"barcodes"`
	Micr      Micr      `json:"micr"`
	PatchCode PatchCode `json:"patchCode"`
	Vendors   []Vendor  `json:"vendors"`
}

type Address struct {
	ImageNumber int64  `json:"imageNumber"`
	ImagePart   int64  `json:"imagePart"`
	MoreParts   string `json:"moreParts"`
	// "lastPartInFile", "lastPartInFileMorePartsPending",
	// "morePartsPending"
	PixelFormatName string `json:"pixelFormatName"`
	SheetNumber     int64  `json:"sheetNumber"`
	Source          string `json:"source"`
	// "feederFront", "feederRear", "flatbed",
	// "planetary", "storage"
	SourceName string `json:"sourceName"`
	StreamName string `json:"streamName"`
}

type Barcode struct {
	Base64Data   string `json:"base64Data"`
	PixelOffsetX int64  `json:"pixelOffsetX"`
	PixelOffsetY int64  `json:"pixelOffsetY"`
	Type         string `json:"type"`
}

type Image struct {
	Compression string `json:"compression"`
	// "group4", "jpeg", "none"
	ImageMerged string `json:"imageMerged"`
	// "imageMerged", "notMerged"
	PixelFormat string `json:"pixelFormat"`
	// "bw", "gray8", "gray16", "rgb24", "rgb48"
	// other values possible?
	PixelHeight  int64 `json:"pixelHeight"`
	PixelOffsetX int64 `json:"pixelOffsetX"`
	PixelOffsetY int64 `json:"pixelOffsetY"`
	PixelWidth   int64 `json:"PixelWidth"`
	Resolution   int64 `json:"resolution"`
}

type Micr struct {
	Base64Data string `json:"base64Data"`
	Type       string `json:"type"`
	// "invalid", "micr", "raw"
}

type PatchCode struct {
	Type string `json:"type"`
}

type Status struct {
	Detected string `json:"detected"`
	// "coverOpen", "foldedCorner", "imageError",
	// "misfeed", "multifeed", "paperJam",
	// "staple"
	Success bool `json:"success"`
}

type Vendor struct {
	Vendor string `json:"vendor"`
}
