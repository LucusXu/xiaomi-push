package xmpush

type SendResult struct {
	Result      string 		`json:"result"`
	TraceId   	string 		`json:"trace_id"`
	Code        int64  		`json:"code"`
	Description string 		`json:"description,omitempty"`
	Info        string 		`json:"info,omitempty"`
	Data 		ResultData	`json:"data,omitempty"`
	Reason      string 		`json:"reason,omitempty"`
}

type ResultData struct {
	BadRegids	string `json:"bad_regids,omitempty"`
	Id			string `json:"id,omitempty"`
}

type Result struct {
	Result      string 		`json:"result"`
	TraceId   	string 		`json:"trace_id"`
	Code        int64  		`json:"code"`
	Description string 		`json:"description,omitempty"`
	Info        string 		`json:"info,omitempty"`
	Data 		ResultData	`json:"data,omitempty""`
	Reason      string 		`json:"reason,omitempty"`
}

type StatsResult struct {
	Result
	Data struct {
		Data []struct {
			Date                  string `json:"date"`
			AliasRecipients       int64  `json:"alias_recipients"`
			UserAccountRecipients int64  `json:"useraccount_recipients"`
			RegIDRecipients       int64  `json:"regid_recipients"`
			Received              int64  `json:"received"`
			BroadcastRecipients   int64  `json:"broadcast_recipients"`
			Click                 int64  `json:"click"`
			SingleRecipients      int64  `json:"single_recipients"`
		} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}

type SingleStatusResult struct {
	Result
	Data struct {
		Data struct {
			CreateTime      string `json:"create_time"`
			CreateTimestamp int64  `json:"create_timestamp"`
			TimeToLive      string `json:"time_to_live"`
			ClickRate       string `json:"click_rate"`
			MsgType         string `json:"msg_type"`
			DeliveryRate    string `json:"delivery_rate"`
			Delivered       int32  `json:"delivered"`
			ID              string `json:"id"`
			Click           int32  `json:"click"`
			Resolved        int32  `json:"resolved"`
		} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}

type BatchStatusResult struct {
	Result
	Data struct {
		Data []struct {
			CreateTime      string `json:"create_time"`
			CreateTimestamp int64  `json:"create_timestamp"`
			TimeToLive      string `json:"time_to_live"`
			ClickRate       string `json:"click_rate"`
			MsgType         string `json:"msg_type"`
			DeliveryRate    string `json:"delivery_rate"`
			Delivered       int32  `json:"delivered"`
			ID              string `json:"id"`
			Click           int32  `json:"click"`
			Resolved        int32  `json:"resolved"`
		} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}

type MultiStatsResult struct {
	Result 	string `json:"result"`
	TraceId string `json:"trace_id"`
	Code 	int `json:"code"`
	Data struct {
		Data map[string]MultiStatsResultDetail `json:"data"`
	} `json:"data,omitempty"`
}

type MultiStatsResultDetail struct {
	CreateTime      		string `json:"create_time"`
	InvalidTarget 			int `json:"invalid_target"`
	MsgSend 				int `json:"msg_send"`
	RawCounter				int `json:"raw_counter"`
	ClickRate       		string `json:"click_rate"`
	Delivered       		int  `json:"delivered"`
	Click           		int  `json:"click"`
	BarClosed				int32  `json:"bar_closed"`
	DeviceConditionUnmatch 	int32 `json:"device_condition_unmatch"`
	CreateTimestamp 		int64  `json:"create_timestamp"`
	TimeToLive      		string `json:"time_to_live"`
	MsgType         		string `json:"msg_type"`
	DeliveryRate    		string `json:"delivery_rate"`
	ID              		string `json:"id"`
	Resolved        		int  `json:"resolved"`
	AppNotRegister  		int `json:"app_not_register"`
	MsgDisplay				int32 `json:"msg_display"`
}

type InvalidRegIDsResult struct {
	Result
	Data struct {
		List []string `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type AliasesOfRegIDResult struct {
	Result
	Data struct {
		List []string `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type TopicsOfRegIDResult struct {
	Result
	Data struct {
		List []string `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type XmCallBody struct {
	BarStatus string `json:"barStatus,omitempty"`
	Type      int `json:"type,omitempty"`
	Targets   string `json:"targets,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}