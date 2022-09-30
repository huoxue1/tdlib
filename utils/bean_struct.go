package utils

type JxBean struct {
	BeanNum int `json:"beanNum"`
	List    []struct {
		Amount      int    `json:"amount"`
		CreateDate  string `json:"createDate"`
		VisibleInfo string `json:"visibleInfo"`
	} `json:"list"`
	WillExpireNum int `json:"willExpireNum"`
}

type Detail struct {
	Date         string `json:"date"`
	Amount       string `json:"amount"`
	EventMassage string `json:"eventMassage"`
}

type TotalBeanResp struct {
	Data struct {
		JdVvipCocoonInfo struct {
			JdVvipCocoonStatus string `json:"JdVvipCocoonStatus"`
		} `json:"JdVvipCocoonInfo"`
		JdVvipInfo struct {
			JdVvipStatus string `json:"jdVvipStatus"`
		} `json:"JdVvipInfo"`
		AssetInfo struct {
			AccountBalance string `json:"accountBalance"`
			BaitiaoInfo    struct {
				AvailableLimit     string `json:"availableLimit"`
				BaiTiaoStatus      string `json:"baiTiaoStatus"`
				Bill               string `json:"bill"`
				BillOverStatus     string `json:"billOverStatus"`
				Outstanding7Amount string `json:"outstanding7Amount"`
				OverDueAmount      string `json:"overDueAmount"`
				OverDueCount       string `json:"overDueCount"`
				UnpaidForAll       string `json:"unpaidForAll"`
				UnpaidForMonth     string `json:"unpaidForMonth"`
			} `json:"baitiaoInfo"`
			BeanNum   string `json:"beanNum"`
			BtFfkInfo struct {
				AppId       string `json:"appId"`
				LinkUrl     string `json:"linkUrl"`
				NumText     string `json:"numText"`
				NumUnitText string `json:"numUnitText"`
				Status      string `json:"status"`
				Subtitle    string `json:"subtitle"`
				Title       string `json:"title"`
			} `json:"btFfkInfo"`
			CouponNum  string `json:"couponNum"`
			CouponRed  string `json:"couponRed"`
			RedBalance string `json:"redBalance"`
		} `json:"assetInfo"`
		FavInfo struct {
			ContentNum  string `json:"contentNum"`
			FavDpNum    string `json:"favDpNum"`
			FavGoodsNum string `json:"favGoodsNum"`
			FavShopNum  string `json:"favShopNum"`
			FootNum     string `json:"footNum"`
			IsGoodsRed  string `json:"isGoodsRed"`
			IsShopRed   string `json:"isShopRed"`
		} `json:"favInfo"`
		GameBubbleList []struct {
			CarouselInfos []struct {
				Icon string `json:"icon"`
				Text string `json:"text"`
			} `json:"carouselInfos"`
			Key   string `json:"key"`
			Title string `json:"title"`
		} `json:"gameBubbleList"`
		GrowHelperCoupon struct {
			AddDays     int     `json:"addDays"`
			BatchId     int     `json:"batchId"`
			CouponKind  int     `json:"couponKind"`
			CouponModel int     `json:"couponModel"`
			CouponStyle int     `json:"couponStyle"`
			CouponType  int     `json:"couponType"`
			Discount    float64 `json:"discount"`
			LimitType   int     `json:"limitType"`
			MsgType     int     `json:"msgType"`
			Quota       float64 `json:"quota"`
			RoleId      int     `json:"roleId"`
			State       int     `json:"state"`
			Status      int     `json:"status"`
		} `json:"growHelperCoupon"`
		KplInfo struct {
			KplInfoStatus string `json:"kplInfoStatus"`
			Mopenbp17     string `json:"mopenbp17"`
			Mopenbp22     string `json:"mopenbp22"`
		} `json:"kplInfo"`
		OrderInfo struct {
			CommentCount     string        `json:"commentCount"`
			Logistics        []interface{} `json:"logistics"`
			OrderCountStatus string        `json:"orderCountStatus"`
			ReceiveCount     string        `json:"receiveCount"`
			WaitPayCount     string        `json:"waitPayCount"`
		} `json:"orderInfo"`
		PlusFloor struct {
			LeftTabs []struct {
				ContentType int    `json:"contentType"`
				ImageUrl    string `json:"imageUrl"`
				Link        string `json:"link"`
				SubTitle    string `json:"subTitle"`
				Title       string `json:"title"`
			} `json:"leftTabs"`
			MidTabs []struct {
				ContentType int    `json:"contentType"`
				ImageUrl    string `json:"imageUrl"`
				Link        string `json:"link"`
				SubTitle    string `json:"subTitle"`
				Title       string `json:"title"`
			} `json:"midTabs"`
			RightTabs []struct {
				ContentType int    `json:"contentType"`
				ImageUrl    string `json:"imageUrl"`
				Link        string `json:"link"`
				SubTitle    string `json:"subTitle"`
				Title       string `json:"title"`
			} `json:"rightTabs"`
		} `json:"plusFloor"`
		PlusPromotion struct {
			Status int `json:"status"`
		} `json:"plusPromotion"`
		TfAdvertInfo struct {
			Status string `json:"status"`
		} `json:"tfAdvertInfo"`
		UserInfo struct {
			BaseInfo struct {
				AccountType    string `json:"accountType"`
				Alias          string `json:"alias"`
				BaseInfoStatus string `json:"baseInfoStatus"`
				CurPin         string `json:"curPin"`
				DefinePin      string `json:"definePin"`
				HeadImageUrl   string `json:"headImageUrl"`
				LevelName      string `json:"levelName"`
				Nickname       string `json:"nickname"`
				UserLevel      string `json:"userLevel"`
			} `json:"baseInfo"`
			IsHideNavi     string `json:"isHideNavi"`
			IsHomeWhite    string `json:"isHomeWhite"`
			IsJTH          string `json:"isJTH"`
			IsKaiPu        string `json:"isKaiPu"`
			IsPlusVip      string `json:"isPlusVip"`
			IsQQFans       string `json:"isQQFans"`
			IsRealNameAuth string `json:"isRealNameAuth"`
			IsWxFans       string `json:"isWxFans"`
			Jvalue         string `json:"jvalue"`
			OrderFlag      string `json:"orderFlag"`
			PlusInfo       struct {
			} `json:"plusInfo"`
			TmpActWaitReceiveCount string `json:"tmpActWaitReceiveCount"`
			XbKeepLink             string `json:"xbKeepLink"`
			XbKeepOpenStatus       string `json:"xbKeepOpenStatus"`
			XbKeepScore            string `json:"xbKeepScore"`
			XbScore                string `json:"xbScore"`
		} `json:"userInfo"`
		UserLifeCycle struct {
			IdentityId      string `json:"identityId"`
			LifeCycleStatus string `json:"lifeCycleStatus"`
			TrackId         string `json:"trackId"`
		} `json:"userLifeCycle"`
	} `json:"data"`
	Msg       string `json:"msg"`
	Retcode   string `json:"retcode"`
	Timestamp int64  `json:"timestamp"`
}
