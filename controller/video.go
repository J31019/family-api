package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MonkeyBuisness/family-api/model"
)

// VideoList returns video list.
func (s Service) VideoList(ctx context.Context) ([]model.Video, error) {
	list, err := loadVideoList()
	if err != nil {
		return nil, err
	}
	for i := range list {
		list[i].URL = nil
	}
	return list, nil
}

// VideoURL returns video URL by ID.
func (s Service) VideoURLByID(ctx context.Context, videoID string) (*model.VideoURL, error) {
	list, err := loadVideoList()
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].URL == nil {
			continue
		}
		if strings.EqualFold(videoID, list[i].ID) {
			return list[i].URL, nil
		}
	}
	return nil, fmt.Errorf("could not find video by ID %s", videoID)
}

func loadVideoList() ([]model.Video, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path.Join(wd, "video.json"))
	if err != nil {
		return nil, err
	}
	var video []model.Video
	if err := json.Unmarshal(data, &video); err != nil {
		return nil, err
	}
	return video, nil
}
