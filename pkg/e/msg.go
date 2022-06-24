package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "fail",
	INVALID_PARAMS:        "Request parameter error",
	INVALID_TOKEN:         "Request parameter token is null",
	ERROR_EXIST_TAG:       "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:  "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:   "该标签不存在",
	ERROR_GET_TAGS_FAIL:   "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:  "统计标签失败",
	ERROR_ADD_TAG_FAIL:    "新增标签失败",
	ERROR_EDIT_TAG_FAIL:   "修改标签失败",
	ERROR_DELETE_TAG_FAIL: "删除标签失败",
	ERROR_EXPORT_TAG_FAIL: "导出标签失败",
	ERROR_IMPORT_TAG_FAIL: "导入标签失败",

	ERROR_EXIST_USER:            "The user exists",
	ERROR_EXIST_USER_FAIL:       "Get the existing user failure",
	ERROR_NOT_EXIST_USER:        "The user does not exist",
	ERROR_GET_USER_FAIL:         "Getting User failed",
	ERROR_GET_USERS_FAIL:        "Getting all User failed",
	ERROR_COUNT_USER_FAIL:       "统计User失败",
	ERROR_ADD_USER_FAIL:         "新增User失败",
	ERROR_EDIT_USER_FAIL:        "修改User失败",
	ERROR_DELETE_USER_FAIL:      "删除User失败",
	ERROR_CHECK_EXIST_USER_FAIL: "Check if the user exists failed",

	ERROR_NOT_EXIST_GRAPH:           "This Graph does not exist",
	ERROR_ADD_GRAPH_FAIL:            "新增Graph失败",
	ERROR_DELETE_GRAPH_FAIL:         "删除Graph失败",
	ERROR_CHECK_EXIST_GRAPH_FAIL:    "检查Graph是否存在失败",
	ERROR_EXIST_GRAPH:               "Graph is exist",
	ERROR_EDIT_GRAPH_FAIL:           "修改Graph失败",
	ERROR_COUNT_GRAPH_FAIL:          "统计Graph失败",
	ERROR_GET_GRAPHS_FAIL:           "获取多个Graph失败",
	ERROR_GET_GRAPH_FAIL:            "Obtain a single CID failure",
	ERROR_GEN_GRAPH_POSTER_FAIL:     "生成Graph海报失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
