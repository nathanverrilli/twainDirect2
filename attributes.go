package twain_direct

const MICRONS_PER_INCH int64 = 25400

// TwainDirectAttributes recognized Twain Direct attributes
func TwainDirectAttributes() (Attributes []string) {
	return []string{
		"alarms",
		"alarmVolume",
		"automaticDeskew",
		"automaticSize",
		"barcodes",
		"bitDepthReduction",
		"brightness",
		"compression",
		"continuousScan",
		"contrast",
		"cropping",
		"discardBlankImages",
		"doubleFeedDetection",
		"doubleFeedDetectionLength",
		"doubleFeedDetectionResponse",
		"doubleFeedDetectionSensitivity",
		"flipRotation",
		"height",
		"imageMerge",
		"imageMergeHeightThreshold",
		"invert",
		"jpegQuality",
		"micr",
		"mirror",
		"noiseFilter",
		"overScan",
		"numberOfSheets",
		"offsetX",
		"offsetY",
		"patchCodes",
		"resolution",
		"rotation",
		"sheetHandling",
		"sheetSize",
		"threshold",
		"uncalibratedImage",
		"width",
	}
}

type SheetSize struct {
	Name   string
	Width  int64
	Height int64
}

// TwainDirectSheetSizes iso and din sizes in microns
func TwainDirectSheetSizes() (SheetSizes []SheetSize) {
	return []SheetSize{
		{"din4A0", 1682000, 2378000},
		{"din2A0", 1189000, 1682000},
		{"isoA0", 841000, 118900},
		{"isoA1", 594000, 841000},
		{"isoA2", 420000, 594000},
		{"isoA3", 297000, 420000},
		{"isoA4", 210000, 297000},
		{"isoA5", 148000, 210000},
		{"isoA6", 105000, 148000},
		{"isoA7", 74000, 105000},
		{"isoA8", 52000, 7400},
		{"isoA9", 37000, 52000},
		{"isoA10", 26000, 37000},
		{"isoB0", 1000000, 1414000},
		{"isoB1", 707000, 1000000},
		{"isoB2", 500000, 707000},
		{"isoB3", 353000, 500000},
		{"isoB4", 250000, 353000},
		{"isoB5", 176000, 250000},
		{"isoB6", 125000, 176000},
		{"isoB7", 88000, 125000},
		{"isoB8", 62000, 88000},
		{"isoB9", 44000, 62000},
		{"isoB10", 31000, 44000},
		{"isoC0", 917000, 1297000},
		{"isoC1", 648000, 917000},
		{"isoC2", 458000, 648000},
		{"isoC3", 324000, 458000},
		{"isoC4", 229000, 324000},
		{"isoC5", 162000, 229000},
		{"isoC6", 114000, 162000},
		{"isoC7", 81000, 114000},
	}

}
