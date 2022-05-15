package interface_adapter

import "github.com/fyk7/code-snippets-app/app/domain/model"

type TagPostReq struct {
	TagName string `json:"tag_name"`
}

func (spr *TagPostReq) ConvertToModel() model.Tag {
	return model.Tag{
		TagName: spr.TagName,
	}
}

type TagPutReq struct {
	TagID   uint64 `json:"tag_id"`
	TagName string `json:"tag_name"`
}

func (spr *TagPutReq) ConvertToModel() model.Tag {
	return model.Tag{
		TagID:   uint64(spr.TagID),
		TagName: spr.TagName,
	}
}
