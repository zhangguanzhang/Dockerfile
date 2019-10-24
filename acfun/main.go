package main

import (
	"acfun/acfun"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"os"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	var (
		username, password string
		tagNames           []string
		submitInfo         = acfun.SubmitVedioInfo{TagNames: [acfun.MaxTags]string{}}
	)
	flag.StringVarP(&username, "username", "u", "", "login name, shuold better use phone number")
	flag.StringVarP(&password, "password", "p", "", "login password")
	flag.StringVarP(&submitInfo.Title, "name", "n", "", "title name")
	flag.StringVarP(&submitInfo.Description, "description", "d", "", "Description info")
	flag.IntVarP(&submitInfo.ChannelID, "channel-id", "c", 86, "channelID")
	flag.IntVarP(&submitInfo.CreationType, "creation-type", "", 1, "creationType")
	flag.Int64SliceVarP(&submitInfo.VideoIDs, "id", "i", nil, "The id of the uploaded but unsubmitted video")
	flag.StringSliceVarP(&tagNames, "tags", "t", nil, "tags, count must <= 4")
	flag.StringVarP(&submitInfo.PicFile, "pic", "", "", "cover pic")
	flag.Parse()

	if password == "" {
		log.Fatal("password cannot be empty!")
	}
	// arg为1的话必须带上 上传但是未投稿的视频id, 第一个arg为封面图
	if flag.NArg() == 0 && len(submitInfo.VideoIDs) == 0 {
		log.Fatal("At least one video needs to be submitted")
	}

	for _, v := range flag.Args() {
		submitInfo.Videos = append(submitInfo.Videos, v)
	}
	for i, v := range tagNames {
		if i >= 4 {
			break
		}
		submitInfo.TagNames[i] = v
	}

	ac, err := acfun.NewAcfun(username, password, os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Login failed: %s", err)
	}
	log.Infoln("Has successfully logged in")

	err = ac.SubmitVideos(&submitInfo)
	if err != nil {
		log.Fatalf("submit  failed: %s", err)
	}
	log.Infoln("Successfully submit!")
}

