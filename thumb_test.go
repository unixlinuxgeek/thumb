// Work!!!
//  ./bin/ffmpeg-amd64-static/ffmpeg -y -ss 00:01 -i /home/geek/Videos/in.mp4 -i /home/geek/Videos/in.jpg -map 1 -map 0 -c copy -disposition:0 attached_pic /home/geek/Videos/out_004.mp4

package thumb

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	var tn = time.Now()
	var u = tn.Unix()
	var stamp = strconv.FormatInt(u, 10)

	err := Gen("./test_data/in.mp4", "./test_data/out_"+stamp+".png", "00:00:10")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("./test_data/out_" + stamp + ".png"); errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
}

func TestIns(t *testing.T) {
	var tn = time.Now()
	var u = tn.Unix()
	var stamp = strconv.FormatInt(u, 10)

	err := Ins("./test_data/in.mp4", "./test_data/out_"+stamp+".png", "./test_data/out_video_"+stamp+".mp4", "00:00:10")
	if err != nil {
		t.Fatal(err)
	}
}
