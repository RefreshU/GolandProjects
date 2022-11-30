package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

const ece_event_topic = "ece_event_c"
const order_topic = "mysql2es_t_orders"
const ods_t_customer_crm = "obs_t_customer_crm"
const stock_app_c_user_info = "StockAppCUserInfoPushTopic"
const all_app_c_user_info = "AllAppCUserInfoPushTopic"
type ReqstMsg struct {
	EceMsg string `json:"ece_msg"`
	OrderMsg string `json:"order_msg"`
	CustomerMsg string `json:"customer_msg"`
	StockCUserInfo string `json:"stock_c_user_info"`
}

func main() {
	producer, err := newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	//consumer, err := sarama.NewConsumer(brokers, nil)
	//if err != nil {
	//	fmt.Println("Could not create consumer: ", err)
	//}

	//subscribe(topic, consumer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello Sarama!") })

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		//err := r.ParseForm()
		if err != nil {
			log.Printf("error: {%s}", err.Error())
		}
		var reqMsg ReqstMsg
		log.Println(json.NewDecoder(r.Body).Decode(&reqMsg))
		//log.Printf("msg is %s", r.FormValue("msg"))
		//msg := prepareMessage(topic, r.FormValue("msg"))
		//eceMsg := `{"url_time_unix":1662691129697,"p__city":"重庆市","p_product_id":"p_620db51ee4b054255d9eb33b","project":"c_production","type":"track","p_resource_type":"2","p_is_paied":1,"p__receive_time":1662691129705,"p_browser_env":"wechat","p__url":"https://appeegq9yqr6896.h5.xiaoeknow.com/p/course/audio/a_620db846e4b04d7e2fcca29d?type=2&pro_id=p_620db51ee4b054255d9eb33b&auto=true","p__model":"SM-G9280","p__url_path":"/p/course/audio/a_620db846e4b04d7e2fcca29d","day":19244,"login_id":"u_60264bfab9998_O3XJ6IdKFC","p__ip":"183.70.67.35","create_time":1662691129705,"p_user_agent":"Mozilla/5.0 (Linux; Android 7.0; SM-G9280 Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4313 MMWEBSDK/20220805 Mobile Safari/537.36 MMWEBID/3175 MicroMessenger/8.0.27.2220(0x28001B53) WeChat/arm64 Weixin NetType/4G Language/zh_CN ABI/arm64","p_platform":"h5","p__browser_version":"8.0.27.2220","p_from_share":0,"p__lib_version":"0.29.1","p__manufacturer":"Samsung","p__is_first_day":false,"p__lib":"js","p_c_user_id":"u_60264bfab9998_O3XJ6IdKFC","p__timezone_offset":-480,"distinct_id":"u_60264bfab9998_O3XJ6IdKFC","p_shop_version_type":4,"track_id":"460718284","p_session_id":"sid_1831fa8ba56257d355133ff2734fb6cad75d","p__lib_method":"code","p__network_type":"4g","month_id":632,"p_sharer_id":"","p_resource_id":"a_620db846e4b04d7e2fcca29d","p__referrer":"https://appeegq9yqr6896.h5.xiaoeknow.com/p/course/audio/a_620db849e4b0beaee42acc65?type=2&pro_id=p_620db51ee4b054255d9eb33b&auto=true","p__browser":"wchat","p__latest_traffic_source_type":"直接流量","p_pv_id":"pv_1832004dd8717ed8b2f301638b46ccf3ab15","url_ip":"183.70.67.35","p__province":"重庆市","p_share_user_id":"","event":"view_page","p_user_channel":"","key":"view_page_C#h5#audio#audio_info_null488a5130-d5ff-458e-838f-445fdc25c59f","p__screen_height":712,"p_abtest_cookie":0,"p__os":"Android","p__latest_search_keyword":"未取到值_直接打开","p__latest_referrer":"","p__referrer_host":"appeegq9yqr6896.h5.xiaoeknow.com","week_id":2749,"p_page_id":"C#h5#audio#audio_info","p__country":"中国","p_l_program":"knowledge_shop","p_app_id":"appeEgq9yQR6896","_track_id":460718284,"url_project":"c_production","p_page_module":"H5","p__os_version":"7.0","anonymous_id":"183159cbbc712-051d1d747e39a8-730d4d3a-284800-183159cbbca0","p_introduce_way":"-1","time":1662691129697,"original_id":"","p__screen_width":400}`
		//orderMsg := `{"data":{"app_id":"appibud2bsi9969","content_app_id":null,"order_id":"o_1662691208_631aa788ea1e2_03898497","user_id":"u_62c00a3e4f09c_PfZKd5idcP","pay_way":0,"payment_type":3,"resource_type":6,"resource_id":"p_62f7dc45e4b0c94264874692","product_id":"p_62f7dc45e4b0c94264874692","count":1,"channel_id":"","channel_info":"","share_user_id":"u_61b5cdf75c8fc_PlAIJDyqpj","share_type":5,"purchase_name":"Y20造句场景篇（新版）","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appibud2bsi9969/image/b_u_61b480c52c675_HgS56Cr8/l2beo49k0317.jpeg","cu_id":"","cou_price":0,"discount_id":"","discount_price":0,"price":9000,"order_state":0,"goods_type":0,"ship_state":0,"out_order_id":null,"transaction_id":null,"wx_app_type":1,"period":null,"use_collection":2,"settle_status":0,"distribute_type":1,"que_check_state":0,"distribute_price":null,"distribute_percent":null,"superior_distribute_user_id":null,"superior_distribute_price":null,"superior_distribute_percent":null,"related_id":"","is_renew":0,"created_at":"2022-09-09 10:40:08","updated_at":"2022-09-09 10:40:09","source":0,"agent":"Mozilla/5.0 (Linux; Android 10; WLZ-AL10 Build/HUAWEIWLZ-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4313 MMWEBSDK/20220805 Mobile Safari/537.36 MMWEBID/9870 MicroMessenger/8.0.27.2220(0x28001B3F) WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64","pay_time":"1970-01-01 08:00:00","settle_time":"1970-01-01 08:00:00","refund_time":"1970-01-01 08:00:00","refund_money":null,"invalid_time":"2022-09-09 12:40:08"},"op":"-U"}`
		msg := prepareMessage(order_topic, reqMsg.OrderMsg)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Fprintf(w, "%s error occured.", err.Error())
		} else {
			fmt.Fprintf(w, "Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
		}
		ece_msg := prepareMessage(ece_event_topic, reqMsg.EceMsg)
		partition1, offset1, err1 := producer.SendMessage(ece_msg)
		if err1 != nil {
			fmt.Fprintf(w, "%s error occured.", err1.Error())
		} else {
			fmt.Fprintf(w, "Message was saved to partion: %d.\nMessage offset is: %d.\n", partition1, offset1)
		}
		customer_msg := prepareMessage(ods_t_customer_crm, reqMsg.CustomerMsg)
		partition2, offset2, err2 := producer.SendMessage(customer_msg)
		if err2 != nil {
			fmt.Fprintf(w, "%s error occured.", err1.Error())
		} else {
			fmt.Fprintf(w, "Message was saved to partion: %d.\nMessage offset is: %d.\n", partition2, offset2)
		}
		stock_c_user_info := prepareMessage(all_app_c_user_info, reqMsg.StockCUserInfo)
		partition3, offset3, err3 := producer.SendMessage(stock_c_user_info)
		if err3 != nil {
			fmt.Fprintf(w, "%s error occured.", err1.Error())
		} else {
			fmt.Fprintf(w, "Message was saved to partion: %d.\nMessage offset is: %d.\n", partition3, offset3)
		}
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, html.EscapeString(getMessage())) })

	log.Fatal(http.ListenAndServe(":8081", nil))
}

