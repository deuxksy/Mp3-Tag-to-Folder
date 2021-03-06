package main

import (
	"fmt"
	"github.com/dhowden/tag"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func example(path string) {
	f, _ := os.Open(path)
	m, _ := tag.ReadFrom(io.ReadSeeker(f))

	var di, _ = m.Disc()
	var ti, _ = m.Track()
	log.Printf("%s-%s-%d-%d-%s", m.Artist(), m.Album(), di, ti, m.Title())
}

func example2(oldPathFile string) (string, string) {
	f, _ := os.Open(oldPathFile)
	m, _ := tag.ReadFrom(io.ReadSeeker(f))
	var file, newPath string
	if m != nil {
		var di, _ = m.Disc()
		//var ti, _ = m.Track()
		_, file = filepath.Split(oldPathFile)
		//log.Printf("%s-%s-%d-%s", strings.TrimSpace(m.AlbumArtist()), strings.TrimSpace(m.Album()), di, strings.TrimSpace(m.Title()))
		newPath = fmt.Sprintf("D:/Users/crom/Music/ts/%s/%s/%d",
			strings.Replace(
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								strings.Replace(
									strings.Replace(
										strings.Replace(
											strings.Replace(
												strings.TrimSpace(m.AlbumArtist()),
												":", "：", -1),
											"/", "／", -1),
										"?", "？", -1),
									"*", "＊", -1),
								">", "＞", -1),
							"<", "＜", -1),
						"\"", "＂", -1),
					"|", "｜", -1),
				"...", "…", -1),
			strings.Replace(
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								strings.Replace(
									strings.Replace(
										strings.Replace(
											strings.Replace(
												strings.Replace(
													strings.TrimSpace(m.Album()),
													":", "：", -1),
												"/", "／", -1),
											"?", "？", -1),
										"*", "＊", -1),
									">", "＞", -1),
								"<", "＜", -1),
							"\"", "＂", -1),
						"|", "｜", -1),
					"...", "…", -1),
				"..", "…", -1),
			di)
		err := os.MkdirAll(newPath, 600)
		if err != nil {
			log.Println(err)
		}
	}
	return oldPathFile, fmt.Sprintf("%s/%s", newPath, file)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func copy(src, dst string) (int64, error) {
	if !fileExists(dst) {
		sourceFileStat, err := os.Stat(src)
		if err != nil {
			return 0, err
		}

		if !sourceFileStat.Mode().IsRegular() {
			return 0, fmt.Errorf("%s is not a regular file", src)
		}

		source, err := os.Open(src)
		if err != nil {
			return 0, err
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			return 0, err
		}
		defer destination.Close()
		nBytes, err := io.Copy(destination, source)
		return nBytes, err
	}
	return 0, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.Println(strings.Replace(strings.TrimSpace("첼로를 좋아하세요?"), "?", "？", -1))

	_ = filepath.Walk("D:/Users/crom/Music/Tracks", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		info, _ = os.Stat(path)
		if !info.IsDir() {
			_, newPathFile := example2(path)
			nBytes, err := copy(path, newPathFile)
			if err != nil {
				log.Println(err)
			} else {
				if nBytes > 0 {
					fmt.Printf("%d,%s\n", nBytes, path)
				}
			}
		}
		return nil
	})
}
