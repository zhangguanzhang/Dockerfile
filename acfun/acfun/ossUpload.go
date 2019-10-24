package acfun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"time"
)

type uploadAddress struct {
	Endpoint string `json:"Endpoint"`
	Bucket   string `json:"Bucket"`
	FileName string `json:"FileName"`
	partSize int64  //存放文件大小
	fileSize int64
}

type uploadAuth struct {
	SecurityToken   string    `json:"SecurityToken"`
	AccessKeyID     string    `json:"AccessKeyId"`
	ExpireUTCTime   time.Time `json:"ExpireUTCTime"`
	AccessKeySecret string    `json:"AccessKeySecret"`
	Expiration      string    `json:"Expiration"`
	Region          string    `json:"Region"`
}

type partsFinishInfo struct {
	UploadID  string     `json:"uploadId"`
	DoneParts []donePart `json:"doneParts"`
}

// 定义进度条监听器。
type OssProgressListener struct {
	partIndex          int
	partNum            int
	totalSize          int64
	ConsumedTotalBytes int64
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		//fmt.Printf("Part %d: transfer Started, ConsumedBytes: %d, TotalBytes %d, %d part in total.\n",
		//	listener.partIndex, event.ConsumedBytes, event.TotalBytes, listener.partNum)
	case oss.TransferDataEvent:
		fmt.Printf("\rParts: %d/%d, CurrentPartTotalBytes %d, ConsumedTotalBytes: %d, TotalBytes %d %d%%.",
			listener.partIndex, listener.partNum, event.TotalBytes, listener.ConsumedTotalBytes+event.ConsumedBytes,
			listener.totalSize, (listener.ConsumedTotalBytes+event.ConsumedBytes)*100/listener.totalSize)
	case oss.TransferCompletedEvent:
		listener.ConsumedTotalBytes += event.ConsumedBytes
		//fmt.Printf("\ntransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
		//	event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

//https://help.aliyun.com/document_detail/88604.html?spm=a2c4g.11186623.2.14.46972d74XdyyyU#title-6lv-fsj-fey
func ossUpload(address *uploadAddress, auth *uploadAuth, filePath string) (*partsFinishInfo, error) {
	// 创建OSSClient实例。
	client, err := oss.New(address.Endpoint, auth.AccessKeyID, auth.AccessKeySecret,
		oss.SecurityToken(auth.SecurityToken))
	if err != nil {
		return nil, err
	}
	var FinishInfo = partsFinishInfo{
		DoneParts: make([]donePart, 0),
	}
	// 获取存储空间。
	bucket, err := client.Bucket(address.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//chunks, err := oss.SplitFileByPartNum(filePath, n)
	chunks, err := oss.SplitFileByPartSize(filePath, address.partSize)
	fd, _ := os.Open(filePath)
	defer fd.Close()
	// 初始化一个分片上传事件。
	imur, err := bucket.InitiateMultipartUpload(address.FileName)
	if err != nil {
		return nil, err
	}
	FinishInfo.UploadID = imur.UploadID
	// 上传分片。
	var parts []oss.UploadPart
	OssProgressListener := &OssProgressListener{
		partNum:   len(chunks),
		totalSize: address.fileSize,
	}
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		OssProgressListener.partIndex++
		// 对每个分片调用UploadPart方法上传。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number, oss.Progress(OssProgressListener))
		if err != nil {
			return &FinishInfo, err
		}
		parts = append(parts, part)
		FinishInfo.DoneParts = append(FinishInfo.DoneParts, donePart{Number: part.PartNumber, Etag: part.ETag})
	}
	fmt.Println("")
	_, err = bucket.CompleteMultipartUpload(imur, parts)
	if err != nil {
		return nil, err
	}

	return &FinishInfo, nil
}

