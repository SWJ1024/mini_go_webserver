package main

/*
var (
	EP string
	AK string
	SK string
)

func init() {
	EP = "oss-cn-beijing.aliyuncs.com"
	AK = "LTAI4G7oPAz2uLGHPLQUe6xw"
	SK = "as9y1kaOtoSs9DLDIcia6hAdfwCmqL"
	EP = config.GetOssAddr()
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service errer : %v", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket errer : %v", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(10))
	if err != nil {
		log.Printf("upload file error: %v", err)
		return false
	}
	return true
}

 */