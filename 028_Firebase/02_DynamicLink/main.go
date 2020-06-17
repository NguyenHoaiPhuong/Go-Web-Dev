package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/028_Firebase/02_DynamicLink/config"
	"google.golang.org/api/firebasedynamiclinks/v1"
	"google.golang.org/api/option"
)

const (
	// DomainURIPrefix : domain uri
	DomainURIPrefix = "https://tokoin.co/wallet"
	// AndroidPackageName : android
	AndroidPackageName = "com.tokoin.wallet"
	// IosBundleID : IOS
	IosBundleID = "com.tokoin.wallet"
	// DesktopFallbackLink : desktop
	DesktopFallbackLink = "https://play.google.com/store/apps/details?id=com.tokoin.wallet"
	// OpenLink : the link that app will open
	// OpenLink = "https://apps.apple.com/us/app/tokoin-my-t-wallet/id1489276175"
	OpenLink = "https://tokoin.co/wallet?user_id=1234"
)

// CreateDynamiclink : create dynamic link
func CreateDynamiclink() (url string, code string) {
	conf := config.GetConfig()
	ctx := context.Background()
	opt := option.WithAPIKey(conf.GoogleAPIKey)
	firebasedynamiclinksService, err := firebasedynamiclinks.NewService(ctx, opt)
	if err != nil {
		panic(err)
	}

	call := firebasedynamiclinksService.ShortLinks.Create(&firebasedynamiclinks.CreateShortDynamicLinkRequest{
		Suffix: &firebasedynamiclinks.Suffix{
			Option: "UNGUESSABLE",
		},
		DynamicLinkInfo: &firebasedynamiclinks.DynamicLinkInfo{
			DomainUriPrefix: DomainURIPrefix,
			Link:            OpenLink,
			AndroidInfo: &firebasedynamiclinks.AndroidInfo{
				AndroidPackageName: AndroidPackageName,
			},
			IosInfo: &firebasedynamiclinks.IosInfo{
				IosBundleId: IosBundleID,
			},
			DesktopInfo: &firebasedynamiclinks.DesktopInfo{
				DesktopFallbackLink: DesktopFallbackLink,
			},
		},
	})
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}

	url = resp.ShortLink
	index := strings.LastIndex(url, "/")
	code = url[index+1:]

	return
}

func testCreateDynamicLink() {
	link, code := CreateDynamiclink()

	fmt.Println(link)
	fmt.Println(code)
}

func main() {
	// testCreateDynamicLink()

	conf := config.GetConfig()
	ctx := context.Background()
	/*opt := option.WithAPIKey(conf.GoogleAPIKey)
	firebasedynamiclinksService, err := firebasedynamiclinks.NewService(ctx, opt)*/
	opt := option.WithCredentialsFile(conf.GoogleCredentials)
	firebasedynamiclinksService, err := firebasedynamiclinks.NewService(ctx, opt)
	if err != nil {
		panic(err)
	}

	shortLink := "https://tokoin.co/wallet/nY4uAwqWetBTfyyb9"
	call := firebasedynamiclinksService.(shortLink)
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}
	bs, err := resp.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}
