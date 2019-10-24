package acfun

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	CoverTokenURL       = domain + "/v2/user/content/upToken"
	CoverUploadURL      = "https://upload.qiniup.com/"
	OSSTokenURL         = domain + "/rest/pc-direct/upload/ali/getToken"
	UploadFinishURL     = domain + "/rest/pc-direct/upload/ali/uploadFinish"
	CreateVideoURL      = domain + "/rest/pc-direct/contribute/createVideo"
	SubmitVideosURL     = domain + "/contribute/video"
	CertifiedCookieName = "stochastic" //投稿的certified值来源于cookies里
	VideoPendingURL     = domain + "/rest/pc-direct/contribute/getVideoPendings"
)

//图片token
type vdata struct {
	Uptoken string `json:"uptoken"`
	URL     string `json:"url"`
}

type ossTokenJson struct {
	baseInfo
	IsOversea     bool   `json:"isOversea"`
	fileSize      int64  //存放文件大小,获取上传token起见会写入
	filePath      string //存放文件路径
	UploadAddress string `json:"uploadAddress"`
	UploadAuth    string `json:"uploadAuth"`
	RequestID     string `json:"requestId"`
	VideoID       string `json:"videoId"` //用来请求接口后转换成视频的短数字
	HostName      string `json:"host-name"`
	UploadConfig  struct {
		Parallel             int8  `json:"parallel"` //并行
		RetryCount           int8  `json:"retryCount"`
		RetryDurationSeconds int8  `json:"retryDurationSeconds"`
		PartSize             int64 `json:"partSize"`
	} `json:"uploadConfig"`
}

type uploadInfo struct {
	Retry   bool   `json:"retry"`
	IsImage bool   `json:"isImage"`
	Loaded  int8   `json:"loaded"`
	State   string `json:"state"`

	File struct {
	} `json:"file"`
	Endpoint_     interface{} `json:"_endpoint"`
	Bucket_       interface{} `json:"_bucket"`
	Object_       interface{} `json:"_object"`
	localFIlename string
	VideoInfo     struct {
		StorageLocation string `json:"StorageLocation"`
		UserData        struct {
		} `json:"UserData"`
	} `json:"videoInfo"`
	UserData string `json:"userData"`
	Ri       string `json:"ri"`
	VideoID  string `json:"videoId"`
	Endpoint string `json:"endpoint"`
	Bucket   string `json:"bucket"`
	Object   string `json:"object"`

	Checkpoint Checkpoint `json:"checkpoint"`
}

type Checkpoint struct {
	File struct {
	} `json:"file"`
	Name string `json:"name"`
	*partsFinishInfo
	FileSize int64 `json:"fileSize"`
	PartSize int64 `json:"partSize"`
}

type donePart struct {
	Number int    `json:"number"`
	Etag   string `json:"etag"`
}

type VedioPending struct {
	Status       int8   `json:"status"`
	SourceStatus int8   `json:"sourceStatus"`
	CreateTime   string `json:"createTime"`
	VideoKey     string `json:"videoKey"`
	VideoID      string `json:"videoId"`
	FileName     string `json:"fileName"`
}

type SubmitVedioInfo struct {
	CreationType int      `json:"creationType"` // 3原创，1转载
	ChannelID    int      `json:"channelId"`    //86生活娱乐
	Videos       []string //上传视频的文件路径
	VideoIDs     []int64  //上传但是未投稿的视频id
	PicFile      string
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	TagNames     [MaxTags]string `json:"tagNames"`
}

func (ac *Acfun) getCoverToken() (*vdata, error) {
	resp, err := ac.R().SetHeaders(map[string]string{
		"sec-fetch-site": "same-origin",
		"devicetype":     "7",
	}).Get(CoverTokenURL)
	if err != nil {
		return nil, err
	}
	var data = struct {
		baseInfo
		Vdata vdata `json:"vdata"`
	}{}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", err, resp)
	}
	if data.ErrorMsg != "" {
		return nil, errors.New(data.ErrorMsg)
	}
	str, err := base64.StdEncoding.DecodeString(data.Vdata.Uptoken)
	if err != nil {
		return nil, fmt.Errorf("deocde %s", err)
	}
	data.Vdata.Uptoken = strings.SplitN(string(str), ":", 2)[1]
	return &data.Vdata, nil
}

func (ac *Acfun) PostCover(filePath string) (string, error) {
	fh, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fh.Close()
	fi, _ := os.Stat(filePath)

	token, err := ac.getCoverToken()
	if err != nil {
		return "", fmt.Errorf("cannot get token:%s", err)
	}

	resp, err := ac.R().SetFile("file", filePath).SetFormData(
		map[string]string{
			"token":            token.Uptoken,
			"id":               "WU_FILE_0",
			"name":             filepath.Base(filePath),
			"type":             "image/jpeg",
			"lastModifiedDate": fi.ModTime().Format("Mon Jan 02 2006 15:04:05 GMT-0700 (中国标准时间)"),
			"size":             string(fi.Size())}).
		Post(CoverUploadURL)
	if err != nil {
		return "", err
	}
	var data = struct {
		baseInfo
		Hash string `json:"hash"`
		Key  string `json:"key"`
	}{}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return "", fmt.Errorf("%s, %s", err, resp)
	}
	if data.Key == "" {
		return "", errors.New(resp.String())
	}

	return fmt.Sprintf("%s/%s", token.URL, data.Key), nil
}

//文件名不在视频文件后缀内的话接口会报错
func (ac *Acfun) getVedioToken(filePath string) (*ossTokenJson, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	resp, err := ac.R().SetFormData(
		map[string]string{
			"name":     filepath.Base(filePath),
			"size":     strconv.FormatInt(fi.Size(), 10),
			"template": "1",
		}).Post(OSSTokenURL)
	if err != nil {
		return nil, err
	}
	var data = ossTokenJson{}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", err, resp)
	}
	if data.UploadConfig.PartSize == 0 {
		return nil, errors.New(resp.String())
	}
	data.filePath = filePath
	data.fileSize = fi.Size()
	return &data, nil
}

//返回videoId
func (ac *Acfun) UploadVedio(filePath string) (int64, error) {
	tokenInfo, err := ac.getVedioToken(filePath)
	if err != nil {
		return 0, fmt.Errorf("cannot get oss token %s", err)
	}
	var (
		uploadAdress uploadAddress
		uploadAuth   uploadAuth
	)
	result, _ := base64.StdEncoding.DecodeString(tokenInfo.UploadAddress)
	_ = json.Unmarshal(result, &uploadAdress)
	uploadAdress.partSize = tokenInfo.UploadConfig.PartSize
	result, _ = base64.StdEncoding.DecodeString(tokenInfo.UploadAuth)
	_ = json.Unmarshal(result, &uploadAuth)

	uploadAdress.fileSize = tokenInfo.fileSize

	partsFinishInfo, err := ossUpload(&uploadAdress, &uploadAuth, filePath)
	if err != nil {
		return 0, err
	}
	ri, _ := uuid.NewRandom()
	var vedioUploadInfo = &uploadInfo{
		State:         "Success",
		UserData:      "eyJWb2QiOnsiU3RvcmFnZUxvY2F0aW9uIjoiIiwiVXNlckRhdGEiOnt9fX0=",
		localFIlename: filepath.Base(filePath),
		Ri:            ri.String(),
		VideoID:       tokenInfo.VideoID,
		Endpoint:      uploadAdress.Endpoint,
		Bucket:        uploadAdress.Bucket,
		Object:        uploadAdress.FileName,
		Loaded:        1,
		Checkpoint: Checkpoint{
			File:            struct{}{},
			Name:            uploadAdress.FileName,
			FileSize:        tokenInfo.fileSize,
			PartSize:        tokenInfo.UploadConfig.PartSize,
			partsFinishInfo: partsFinishInfo,
		},
	}
	vedioID, err := ac.getVideoID(vedioUploadInfo)
	if err != nil {
		return 0, fmt.Errorf("get vedio id %s", err)
	}
	return vedioID, nil
}

func (ac *Acfun) getVideoID(uploadInfo *uploadInfo) (int64, error) {
	uploadStr, _ := json.Marshal(uploadInfo)

	resp, err := ac.R().SetFormData(map[string]string{
		"uploadInfo": string(uploadStr),
		"errorCode":  "",
		"errorMsg":   "",
	}).SetHeader("Sec-Fetch-Site", "same-origin").Post(UploadFinishURL)
	if err != nil {
		return 0, err
	}
	var data = struct {
		baseInfo
		Result   int    `json:"result"`
		SourceID int64  `json:"sourceId"`
		VideoID  int64  `json:"videoId"`
		HostName string `json:"host-name"`
	}{}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return 0, fmt.Errorf("%s, %s", err, resp)
	}

	if data.Result != 0 {
		return 0, errors.New(data.ErrorMsg)
	}

	resp, err = ac.R().SetFormData(map[string]string{
		"videoKey": uploadInfo.VideoID,
		"fileName": uploadInfo.localFIlename,
	}).SetHeader("Sec-Fetch-Site", "same-origin").Post(CreateVideoURL)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return 0, fmt.Errorf("%s, %s", err, resp)
	}
	if data.Result != 0 {
		return 0, errors.New(data.ErrorMsg)
	}
	return data.VideoID, nil

}

func (ac *Acfun) SubmitVideos(submitInfo *SubmitVedioInfo) error {
	if len(submitInfo.Videos) == 0 && len(submitInfo.VideoIDs) == 0  {
		//fmt.Println(submitInfo.Videos, submitInfo.VideoIDs)
		return errors.New("need a minimum of one video")
	}
	type videoInfo struct {
		IsDeleted int8   `json:"isDeleted"`
		VideoID   int64  `json:"videoId"`
		Title     string `json:"title"`
	}
	videoInfos := make([]*videoInfo, 0)

	var (
		picUrl    string
		err       error
		partIndex int8
	)
	//获取上传但为投稿的视频id
	if len(submitInfo.VideoIDs) > 0 {
		VedioPendingList, err := ac.GetVedioPeindings()
		if err != nil {
			return err
		}
		for _, v := range submitInfo.VideoIDs {
			flag := false
			for _, pvedio := range VedioPendingList {
				if strconv.FormatInt(v, 10) == pvedio.VideoID {
					flag = true
					break
				}
			}
			if flag {
				videoInfos = append(videoInfos, &videoInfo{
					VideoID: v,
					Title:     fmt.Sprintf("Part%d", partIndex),
					IsDeleted: 0,
				})
				return nil
			} else {
				log.Errorf("video %d is not on the cloud", v)
			}
		}
	}

	//上传封面
	if submitInfo.PicFile != "" {
		//图片直链情况
		if strings.HasPrefix(submitInfo.PicFile, "http"){
			picUrl = submitInfo.PicFile
		} else {
			picUrl, err = ac.PostCover(submitInfo.PicFile)
			if err != nil {
				return fmt.Errorf("post cover %s", err)
			}
		}
	}

	for _, v := range submitInfo.Videos {
		partIndex++
		id, err := ac.UploadVedio(v)
		if err != nil {
			if len(videoInfos) != 0 { //上传出错打印已经成功上传的id
				log.Warnf("these successful uploads: %v", videoInfos)
			}
			return err
		}
		log.Infof("Part %d: %d has been successfully uploaded", partIndex, id)
		videoInfos = append(videoInfos, &videoInfo{VideoID: id,
			Title:     fmt.Sprintf("Part%d", partIndex),
			IsDeleted: 0,
		})
	}
	jsonBytes, _ := json.Marshal(&videoInfos)
	tagNames, _ := json.Marshal(submitInfo.TagNames)
	data := map[string]string{
		"title":           submitInfo.Title,
		"description":     submitInfo.Description,
		"tagNames":        string(tagNames),
		"creationType":    strconv.Itoa(submitInfo.CreationType),
		"channelId":       strconv.Itoa(submitInfo.ChannelID), //娱乐-生活娱乐
		"coverUrl":        picUrl,
		"videoInfos":      string(jsonBytes),
		"isJoinUpCollege": strconv.Itoa(0),
		"publishTime":     "",
		"certified":       ac.CertifiedValue,
	}

	err = ac.submitVedios(data)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Acfun) submitVedios(data map[string]string) error {

	resp, err := ac.R().SetFormData(data).
		SetHeader("Sec-Fetch-Site", "same-origin").
		SetHeader("x-requested-with", "XMLHttpRequest").
		Post(SubmitVideosURL)
	if err != nil {
		return err
	}
	var result = struct {
		Code   int    `json:"code"`
		Errno  int    `json:"errno"`
		Msg    string `json:"msg"`
		Errmsg string `json:"errmsg"`
		Data   struct {
			Result   int    `json:"result"`
			DougaID  int    `json:"dougaId"`
			HostName string `json:"host-name"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return fmt.Errorf("%s, %s", err, resp)
	}
	if result.Data.Result != 0 {
		return errors.New(result.Msg)
	}
	return nil
}

func (ac *Acfun) GetVedioPeindings() ([]VedioPending, error) {

	resp, err := ac.R().
		SetHeader("Sec-Fetch-Site", "same-origin").
		SetHeader("x-requested-with", "XMLHttpRequest").
		Post(VideoPendingURL)
	if err != nil {
		return nil, err
	}
	result := struct {
		Code     int            `json:"code"`
		Result   int            `json:"result"`
		Msg      string         `json:"msg"`
		Videos   []VedioPending `json:"videos"`
		HostName string         `json:"host-name"`
	}{}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", err, resp)
	}
	if result.Result != 0 {
		return nil, errors.New(result.Msg)
	}
	return result.Videos, nil
}

