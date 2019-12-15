package modules

import (
	"crypto/sha1"
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/oauth"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type WeChat struct {
	wc *wechat.Wechat
}

func NewWeChat() *WeChat {
	return &WeChat{
		wc: wechat.NewWechat(&wechat.Config{
			AppID: "wx351c57a5ab698412",
			AppSecret: "0fe0fc7311fa355b4a3ecd91e651a361",
			Token: "wYu33WX12Account29",
			EncodingAESKey: "lgQ87CZxlwD2V9exfVw7hNfGlloDppwFS38Vd4ynIxf",
			Cache: NewWxCache(),
		}),
	}
}

/**
 * redirectURL, scope, state := "http://go.ywycloud.com/wx_redirect", "snsapi_userinfo", "STATE"
 * wc := modules.NewWeChat()
 * wc.AuthRequest( ... )
 * wc.RedirectURI( c.Query("code") )
 */
func (we *WeChat) AuthRequest(writer http.ResponseWriter, req *http.Request, redirectURL string, scope string, state string) (err error) {
	auth := we.wc.GetOauth()
	err = auth.Redirect(writer, req, redirectURL, scope, state)
	if err != nil {
		return
	}

	return
}

func (we *WeChat) RedirectURI(code string) (err error, resToken oauth.ResAccessToken, userInfo oauth.UserInfo) {
	if code == "" {
		return
	}

	auth := we.wc.GetOauth()
	resToken, err = auth.GetUserAccessToken(code)
	if err != nil {
		return
	}

	userInfo, err = auth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
	if err != nil {
		return
	}

	return
}

/**
 * adjust when update the wechat configuration
**/
func WechatIdentify(timestamp string, nonce string, signatureIn string, echostr string) string {
	token := "wYu33WX12Account29"

	/**
	 *	Todo: token identified
	 */
	str := []string{token, timestamp, nonce}
	sort.Strings(str)
	s := sha1.New()
	io.WriteString(s, strings.Join(str, ""))

	signatureGen := fmt.Sprintf("%x", s.Sum(nil))

	if signatureGen != signatureIn {
		fmt.Printf("signatureGen != signatureIn signatureGen=%s,signatureIn=%s\n", signatureGen, signatureIn)
		return ""
	} else {
		if echostr != "" {
			return echostr
		} else {
			fmt.Println("echostr is nil")
			return ""
		}
	}
}

/**
 * =====================================================================================================
 * Todo: Add Cache Interface{}
**/

type wxCache struct {
	r *rd
}

var (
	_ cache.Cache = &wxCache{}
)

func NewWxCache() *wxCache {
	return &wxCache{
		r: InstanceRedis().instance(),
	}
}

func (wxc *wxCache) Get(key string) interface{} {
	return wxc.r.r.Get(key).Val()
}

func (wxc *wxCache) Set(key string, val interface{}, timeout time.Duration) error {
	return wxc.r.r.Set(key, val, timeout).Err()
}

func (wxc *wxCache) IsExist(key string) bool {
	ok := wxc.r.r.Exists(key).Val()
	if ok != 1 {
		return false
	}
	return true
}

func (wxc *wxCache) Delete(key string) error {
	return wxc.r.r.Del(key).Err()
}
