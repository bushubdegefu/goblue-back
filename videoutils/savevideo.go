package videoutils

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"semay.com/admin/database"
	"semay.com/admin/models"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) (string, error) {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[random.Int64()]
	}
	return string(b), nil
}

func CreateVideo(name string) {
	working_dir, _ := os.Getwd()
	video_name := strings.ReplaceAll(name, "Stream/", "")
	video_name = strings.ReplaceAll(video_name, ".mp4", "")
	command := fmt.Sprintf("MP4Box -dash 5000 -frag 5000 -rap -segment-name %s -subsegs-per-sidx 5 -url-template %s.mp4#video %s.mp4#audio", video_name, video_name, video_name)
	fmt.Println(command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = working_dir + "/Stream"

	if err := cmd.Run(); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	time.Sleep(2 * time.Second)
	os.Remove("Stream/" + video_name + ".mp4")
	defer SaveVideoToDatabse(name)

}

func SaveVideoToDatabse(name string) {
	pattern_name := strings.ReplaceAll(name, ".mp4", "")
	pattern := fmt.Sprintf(`.*%s.*`, pattern_name)
	regex := regexp.MustCompile(pattern)

	files, err := filepath.Glob("./Stream/*.mp4")
	if err != nil {
		log.Fatal(err)
	}

	frames, err := filepath.Glob("./Stream/*.m4s")
	if err != nil {
		log.Fatal(err)
	}

	files = append(files, frames...)

	manifest, err := filepath.Glob("./Stream/*.mpd")
	if err != nil {
		log.Fatal(err)
	}

	files = append(files, manifest...)

	filterdFiles := make([]string, 0)
	for _, str := range files {
		if regex.FindStringIndex(str) != nil {
			filterdFiles = append(filterdFiles, str)
		}
	}

	db := database.ReturnSession()
	for frame := range filterdFiles {
		vid_frame, oerr := os.Open(filterdFiles[frame])
		if oerr != nil {
			fmt.Println(oerr)
		}
		read_frame, _ := io.ReadAll(vid_frame)
		vid_name := strings.ReplaceAll(filterdFiles[frame], "Stream/", "")
		var blob = new(models.BlobVideo)
		blob.Name = vid_name
		blob.BlobVideo = read_frame
		tx := db.Begin()
		if err := tx.Create(&blob).Error; err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
		}
		tx.Commit()
		vid_frame.Close()
		time.Sleep(5 * time.Millisecond)
		os.Remove(filterdFiles[frame])
	}

}
