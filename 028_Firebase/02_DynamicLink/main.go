package main

import (
	"context"
	"fmt"

	"google.golang.org/api/firebasedynamiclinks/v1"
	"google.golang.org/api/option"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/028_Firebase/02_DynamicLink/config"
)

func main() {
	// load config
	conf := config.GetConfig()

	ctx := context.Background()
	opt := option.WithCredentialsFile(conf.GoogleCredentials)
	firebasedynamiclinksService, err := firebasedynamiclinks.NewService(ctx, opt)
	if err != nil {
		panic(err)
	}

	// https://firebase.google.com/docs/reference/dynamic-links/link-shortener
	call := firebasedynamiclinksService.ShortLinks.Create(&firebasedynamiclinks.CreateShortDynamicLinkRequest{
		// LongDynamicLink: "https://play.google.com/store/apps/details?id=com.tokoin.wallet",
		Suffix: &firebasedynamiclinks.Suffix{
			Option: "SHORT",
		},
		DynamicLinkInfo: &firebasedynamiclinks.DynamicLinkInfo{
			// DynamicLinkDomain: "itvietnam.page",
			DomainUriPrefix: "https://itvietnam.page/wallet",
			Link:            "https://apps.apple.com/us/app/tokoin-my-t-wallet/id1489276175",
			AndroidInfo: &firebasedynamiclinks.AndroidInfo{
				// AndroidFallbackLink: "https://play.google.com/store/apps/details?id=com.tokoin.wallet",
				// AndroidLink:         "tokoin_wallet",
				AndroidPackageName: "com.tokoin.wallet",
			},
			IosInfo: &firebasedynamiclinks.IosInfo{
				// IosFallbackLink: "https://apps.apple.com/us/app/tokoin-my-t-wallet/id1489276175",
				IosBundleId: "com.tokoin.wallet",
			},
			DesktopInfo: &firebasedynamiclinks.DesktopInfo{
				DesktopFallbackLink: "https://play.google.com/store/apps/details?id=com.tokoin.wallet",
			},
		},
	})
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ShortLink)
}
