package models

const (
	TableNameImages = "images"
)

type imagesModel struct{}

type ImagesList struct {
	Version_id        string `json:"version_id"`
	Docker_image_id   string `json:"docker_image_id"`
	Version_name      string `json:"version_name"`
	App_name          string `json:"app_name"`
	Repository        string `json:"repository"`
	Tag               string `json:"tag"`
	Repository_harbor string `json:"repository_harbor"`
	Create_date       string `json:"create_date"`
}
type ImagesCountNum struct {
	Count_num int64 `json:"count_num"`
}

func (*imagesModel) ListImages(pageNo string, pageSize string, versionName string) (countNum int64, imagesList []ImagesList, err error) {
	sql := "SELECT count(1) count_num FROM docker_image a, version b WHERE a.version_id = b.version_id AND b.version_name LIKE concat('%', ?, '%') limit 0,1000"
	var imagesCountNum ImagesCountNum
	err = OrmerImm().Raw(sql, versionName).QueryRow(&imagesCountNum)
	if err != nil {
		return 0, nil, err
	}
	countNum = imagesCountNum.Count_num
	sql = "SELECT concat(a.version_id, '') AS version_id , concat(a.docker_image_id, '') AS docker_image_id, b.version_name , b.app_name, a.repository, a.tag, a.repository_harbor , concat(a.create_date, '') AS create_date FROM docker_image a, version b WHERE a.version_id = b.version_id AND b.version_name LIKE concat('%', ?, '%') ORDER BY a.create_date DESC LIMIT ?, ?"
	values := [3]string{versionName, pageNo, pageSize}
	_, err = OrmerImm().Raw(sql, values).QueryRows(&imagesList)
	if err != nil {
		return 0, nil, err
	}
	return
}
