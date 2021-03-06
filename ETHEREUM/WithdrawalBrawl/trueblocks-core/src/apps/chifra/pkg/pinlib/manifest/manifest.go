// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package manifest

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
)

type IpfsHash = string

type PinDescriptor struct {
	FileName  string   `json:"fileName"`
	BloomHash IpfsHash `json:"bloomHash"`
	IndexHash IpfsHash `json:"indexHash"`
}

type Manifest struct {
	// TODO: This structure will be updated with better data shortly
	FileName           string        `json:"fileName"`
	IndexFormat        IpfsHash      `json:"indexFormat"`
	BloomFormat        IpfsHash      `json:"bloomFormat"`
	CommitHash         string        `json:"commitHash"`
	PreviousHash       IpfsHash      `json:"prevHash"`
	NewBlockRange      ManifestRange `json:"newBlockRange"`
	PreviousBlockRange ManifestRange `json:"prevBlockRange"`
	NewPins            PinsList      `json:"newPins"`
	PreviousPins       PinsList      `json:"prevPins"`
}

type PinsList []PinDescriptor

// GetCsvOutput returns data for CSV and TSV formats
func (pl *PinsList) GetCsvOutput() *globals.CsvFormatted {
	data := &globals.CsvFormatted{
		Header: []string{
			"fileName", "bloomHash", "indexHash",
		},
	}

	for _, pin := range *pl {
		data.Content = append(data.Content, []string{
			pin.FileName, pin.BloomHash, pin.IndexHash,
		})
	}

	return data
}

// GetJsonOutput returns data for JSON format. In this case
// we just return the whole slice of PinDescriptor
func (pl *PinsList) GetJsonOutput() interface{} {
	return *pl
}
