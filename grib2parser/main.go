package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/yakawa/go-libgrib2"
)

func main() {
	grib2File, err := os.Open("zenmodel_171205/Z__C_RJTD_20171205000000_GSM_GPV_Rgl_FD0006_grib2.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer grib2File.Close()

	data, err := ioutil.ReadAll(grib2File)
	if err != nil {
		log.Fatal(err)
	}

	libgrib2.Read(data)
}
