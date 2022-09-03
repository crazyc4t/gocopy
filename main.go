package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "gocopy",
		Usage: "cp implementation made with go!",
		Action: func(cCtx *cli.Context) error {
			copy(cCtx.Args().Get(0), cCtx.Args().Get(1))
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func copy(src, dst string) int64 {
	fileSrc, err := os.Open(src)
	if err != nil {
		fmt.Println("Couldn't open file")
		log.Fatal(err)
	}
	defer fileSrc.Close()

	fileDst, err := os.Create(dst)
	if err != nil {
		fmt.Println("Couldn't create the destination file, check the path name")
		log.Fatal(err)
	}
	defer fileDst.Close()
	nBytes, err := io.Copy(fileDst, fileSrc)
	fmt.Println("Copied!")
	return nBytes
}
