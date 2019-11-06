package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ncw/directio"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		in, err := directio.OpenFile("hexdump", os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println("Error on open: ", err)
		}

		block := directio.AlignedBlock(directio.BlockSize)
		n, err := io.ReadFull(in, block)
		if err != nil {
			fmt.Println("Error on read: ", err)
		}

		fmt.Println("read", n, "bytes from file")

		w.Header().Set("Content-Type", "text/html")
		w.Write(block)
	})

	http.ListenAndServe(":80", nil)
}
