# Wechat MP Utils for Go

微信公众号SDK for Go

已集成

### 用户
 - 获取用户列表
 - 获取用户基本信息（包括UnionID机制）
 - 批量获取用户基本信息
 - 设置用户备注名

### 标签
 - 获取公众号已创建的标签
 - 创建标签
 - 编辑标签
 - 删除标签
 - 获取标签下粉丝列表
 - 批量为用户打标签
 - 批量为用户取消标签
 - 获取用户身上的标签列表

### 黑名单管理
 - 获取公众号的黑名单列表
 - 拉黑用户
 - 取消拉黑用户

### template
 - 设置所属行业
 - 获取设置的行业信息
 - 获得模板ID
 - 获取模板列表
 - 删除模板
 - 发送模板消息



```go
package wechat // import "github.com/limtech/wechat"

const TEMPLATE_SET_INDUSTRY_API string = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s" ...
const USER_LIST_API string = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s" ...
const USER_BLACKLIST_ALL_API string = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s" ...
const USER_TAG_LIST_API string = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s" ...
const ACCESS_TOKEN_API string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
const JSSDK_SIGNATURE_STRING string = "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
const TICKET_API string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
func GetAccessToken(appId string, appSecret string) (AccessTokenData, error)
func GetJSSDKSignature(ticket string, url string) (signature string, nonceStr string, timestamp int64)
func HttpGet(url string) ([]byte, error)
func HttpPost(url string, data url.Values, headers map[string]string) ([]byte, error)
func HttpPostJson(url string, data interface{}, header map[string]string) ([]byte, error)
func RandomString() string
type AccessTokenData struct{ ... }
type ErrStruct struct{ ... }
type JSSDK struct{ ... }
    func GetJSSDKConfig(appId string, ticket string, url string) JSSDK
type Message struct{ ... }
    func NewMessage(accessToken string) *Message
type MessageTemplate struct{ ... }
type MessageTemplateAll struct{ ... }
type MessageTemplateIndustry struct{ ... }
type MessageTemplateIndustryItem struct{ ... }
type MessageTemplateSendResult struct{ ... }
type Ticket struct{ ... }
    func NewTicket(accessToken string) *Ticket
type TicketData struct{ ... }
type User struct{ ... }
    func NewUser(accessToken string) *User
type UserInfo struct{ ... }
type UserInfoBatch struct{ ... }
type UserInfoData struct{ ... }
type UserList struct{ ... }
type UserListAll struct{ ... }
type UserTagData struct{ ... }
type UserTagIdList struct{ ... }
type UserTagItem struct{ ... }
type UserTagList struct{ ... }
type UserTagUserListData struct{ ... }

```
