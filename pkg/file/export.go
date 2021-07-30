package file

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
)

// DownloadExcel 公共下载execl方法
func DownloadExcel(c *gin.Context, file *excelize.File) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+gotool.IdUtils.IdUUIDToRan(false)+".xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("FileName", gotool.IdUtils.IdUUIDToRan(false)+".xlsx")
	file.Write(c.Writer)
}
