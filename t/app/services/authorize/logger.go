package authorize

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/google/uuid"
//	"regexp"
//	"service-api/src/models/ent"
//	"service-api/src/models/repo/member"
//	"time"
//)
//
//type LoggerService struct {
//	repo member.LoggerRepo
//}
//
//func (a *LoggerService) Write(c *gin.Context, memberId uuid.UUID, token string) (uuid.UUID, error) {
//	agent := c.Request.Header.Get("User-Agent")
//	regex, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP|micromessenger|)")
//	device := regex.FindString(agent)
//
//	d, err := a.repo.Write(ent.MemberAuthorizeLog{
//		Channel:      device,
//		CreateTime:   time.Now(),
//		MemberID:     memberId,
//		Token:        token,
//		ClientIP:     c.ClientIP(),
//		RemoteIP:     c.RemoteIP(),
//		Device:       device,
//		DeviceDetail: agent,
//		Snapshot:     "",
//	})
//
//	if err != nil {
//		return uuid.UUID{}, err
//	}
//
//	return d.ID, nil
//}
