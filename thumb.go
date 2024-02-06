// Capture 1 frame from video
//
//	ffmpeg -i ./in.mp4 -i ./out1.png -map 0 -map 1 -c copy -c:v:1 png -disposition:v:1 attached_pic out.mp4
package thumb

import (
	"log"
	"os"
	"os/exec"
)

func Gen(vidName string, outImg, time string) error {
	f, err := os.OpenFile(vidName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 755)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	app := "./bin/ffmpeg-amd64-static/ffmpeg"
	arg1 := "-y"
	arg2 := "-ss"
	arg3 := time
	arg4 := "-i"
	arg5 := vidName
	arg6 := "-frames:v"
	arg7 := "1"
	arg8 := "-q:v"
	arg9 := "2"
	arg10 := outImg

	cmd := exec.Command(app, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Ins(inVid string, inImg, outVid, time string) error {
	app := "./bin/ffmpeg-amd64-static/ffmpeg"
	// ffmpeg -i in.mp4 -i IMAGE -map 0 -map 1 -c copy -c:v:1 png -disposition:v:1 attached_pic out.mp4

	// ffmpeg -y -i in.mp4 -i ./out1.png -map 0 -map 1 -c copy -c:v:1 png -disposition:v:1 attached_pic out.mp4
	arg01 := "-y"
	arg02 := "-i"
	arg03 := inVid
	arg04 := "-i"
	arg05 := inImg
	arg06 := "-map"
	arg07 := "0"
	arg08 := "-map"
	arg09 := "1"
	arg10 := "-c"
	arg11 := "copy"
	arg12 := "-c:v:1"
	arg13 := "png"
	arg14 := "-disposition:v:1"
	arg15 := "attached_pic"
	arg16 := outVid

	cmd := exec.Command(app, arg01, arg02, arg03, arg04, arg05, arg06, arg07, arg08, arg09, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
	err := cmd.Run()
	return err
}
