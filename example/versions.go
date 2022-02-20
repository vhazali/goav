package main

import (
	"log"

	"github.com/vhazali/goav/avcodec"
	"github.com/vhazali/goav/avdevice"
	"github.com/vhazali/goav/avfilter"
	"github.com/vhazali/goav/avformat"
	"github.com/vhazali/goav/avutil"
	"github.com/vhazali/goav/swresample"
	"github.com/vhazali/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
