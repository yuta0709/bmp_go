package bmp

import (
	"encoding/binary"
	"os"
)

type BitmapHeader struct {
	FileType       [2]byte
	FileSize       uint32
	Reserved1      uint16
	Reserved2      uint16
	HeaderSize     uint32
	InfoHeaderSize uint32
	PictWidth      uint32
	PictHeight     uint32
	PlaneNum       uint16
	BitPerPixel    uint16
	Compression    uint32
	ImageSize      uint32
	HResolution    uint32
	VResolution    uint32
	ColorUsed      uint32
	ColImportant   uint32
}

type RGB24 struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Bmp struct {
	Header     BitmapHeader
	ImageArray []RGB24
}

func Open(path string) (*Bmp, error) {
	file, err := os.Open(path)

	if err != nil{
		return nil, err
	}

	bmp := Bmp{}
	header := &bmp.Header
	err = binary.Read(file, binary.LittleEndian, header)

	if err != nil{
		return nil, err
	}

	imageSize := header.PictWidth * header.PictHeight
	bmp.ImageArray = make([]RGB24, imageSize)
	err = binary.Read(file, binary.LittleEndian, bmp.ImageArray)

	if err != nil{
		return nil, err
	}

	return &bmp, nil
}
