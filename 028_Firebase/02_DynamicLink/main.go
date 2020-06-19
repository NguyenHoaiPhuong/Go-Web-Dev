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

func testGetDynamicLinkStats() {
	conf := config.GetConfig()
	ctx := context.Background()
	opt := option.WithCredentialsFile(conf.GoogleCredentials)
	firebasedynamiclinksService, err := firebasedynamiclinks.NewService(ctx, opt)
	if err != nil {
		panic(err)
	}

	// link := "https://tokoin.co/wallet/user_id=136&email=pqkhanh88@gmail.com&referral_code=fpFaJc"
	// link := "https://tokoin.co/wallet/R56w"
	link := "https://itvietnam.page/wallet/T8GLQtZPAZa2Cuvx5"
	call := firebasedynamiclinksService.V1.GetLinkStats(link)
	call = call.DurationDays(7)
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	for _, eventStat := range resp.LinkEventStats {
		fmt.Println(eventStat.Event)
	}
}

func main() {
	// testCreateDynamicLink()

	testGetDynamicLinkStats()
}
