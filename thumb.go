// Capture 1 frame from video
//
//		./ffmpeg -y -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /home/geek/Videos/out_000111.mp4
//	 ../bin/ffmpeg-amd64-static/ffmpeg -y -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /home/geek/Videos/out_001.mp4
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

	// ffmpeg -y -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /home/geek/Videos/out_00011.mp4
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

// ./bin/ffmpeg-amd64-static/ffmpeg -y -ss 00:01 -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /home/geek/Videos/out_004.mp4
func Ins(inVid string, inImg, outVid, time string) error {
	app := "./bin/ffmpeg-amd64-static/ffmpeg"
	// ./ffmpeg -y -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /var/tmp/out_00011.mp4
	arg01 := "-y"
	//arg02 := "-ss"
	//arg03 := time

	// ./bin/ffmpeg-amd64-static/ffmpeg -y -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /home/geek/Videos/out_009.mp4

	arg04 := "-i"
	arg05 := inVid
	arg06 := "-i"
	arg07 := inImg
	arg08 := "-map"
	arg09 := "1"
	arg10 := "-map"
	arg11 := "0"
	arg12 := "-c"
	arg13 := "copy"
	arg14 := "-disposition:0"
	arg15 := "attached_pic"
	arg16 := outVid

	// arg02, arg03
	cmd := exec.Command(app, arg01, arg04, arg05, arg06, arg07, arg08, arg09, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
	err := cmd.Run()
	return err
}
