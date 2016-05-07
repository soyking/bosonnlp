package bosonnlp

import "testing"
import "strings"

var client = NewBosonNLPClient("your token")

func TestBosonNLPClient_SentimentAnalysis(t *testing.T) {
	resp, err := client.SentimentAnalysis([]string{"他是个傻逼", "美好的世界"})
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatal("data's length is not 2")
		} else {
			t.Logf("%v\n", resp[0])
			t.Logf("%v\n", resp[1])
		}
	}
}

func TestBosonNLPClient_WeiboSentimentAnalysis(t *testing.T) {
	resp, err := client.WeiboAutoSentimentAnalysis([]string{"他是个傻逼", "美好的世界"})
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatal("data's length is not 2")
		} else {
			t.Logf("%v\n", resp[0])
			t.Logf("%v\n", resp[1])
		}
	}
}

func TestBosonNLPClient_NerAnalysis(t *testing.T) {
	resp, err := client.NerAnalysis([]string{"对于该小孩是不是郑尚金的孩子，目前已做亲子鉴定，结果还没出来", "纪检部门仍在调查之中。成都商报记者 姚永忠"})
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatalf("data's length is not 2,is %d \n", len(resp))
		} else {
			t.Logf("%v\n", resp[0])
			t.Logf("%v\n", resp[1])
		}
	}
}

func TestBosonNLPClient_DepparserAnalysis(t *testing.T) {
	resp, err := client.DepparserAnalysis([]string{"我以最快的速度吃了午饭", "我以最快的速度吃了午饭"})
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatalf("data's length is not 2,is %d \n", len(resp))
		} else {
			t.Logf("%v\n", resp[0])
			t.Logf("%v\n", resp[1])
		}
	}
}

func TestBosonNLPClient_KeywordsAnalysis(t *testing.T) {
	q := map[string]string{"top_k": "10", "segmented": ""}
	resp, err := client.KeywordsAnalysis([]string{"病毒式媒体网站：让新闻迅速蔓延", "病毒式媒体网站：让新闻迅速蔓延"}, q)
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatalf("data's length is not 2,is %d \n", len(resp))
		} else {
			t.Logf("%v\n", resp[0])
			t.Logf("%v\n", resp[1])
		}
	}
}

func TestBosonNLPClient_ClassifyAnalysis(t *testing.T) {
	resp, err := client.ClassifyAnalysis([]string{"俄否决安理会谴责叙军战机空袭阿勒颇平民", "邓紫棋谈男友林宥嘉：我觉得我比他唱得好"})
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatalf("data's length is not 2,is %d \n", len(resp))
		} else {
			t.Logf("%v\n", NewsType[resp[0]])
			t.Logf("%v\n", NewsType[resp[1]])
		}
	}
}

func TestBosonNLPClient_SuggestAnalysis(t *testing.T) {
	q := map[string]string{"top_k": "2"}
	resp, err := client.SuggestAnalysis("粉丝", q)
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatalf("data's length is not 2,is %d \n", len(resp))
		} else {
			t.Logf("%v\n", resp[0])
			t.Logf("%v\n", resp[1])
		}
	}
}

func TestBosonNLPClient_TagAnalysis(t *testing.T) {
	q := map[string]string{"space_mode": "0", "oov_level": "4", "t2s": "0", "special_char_conv": "0"}
	resp, err := client.TagAnalysis([]string{"人民法院案件受理制度改革", "下月起法院将有案必立"}, q)
	if err != nil {
		t.Fatal(err)
	} else {
		if len(resp) != 2 {
			t.Fatal("data's length is not 2")
		} else {
			t.Logf("%v\n", resp[0].Tag)
			t.Logf("%v\n", resp[0].Word)
		}
	}
}

func TestBosonNLPClient_TimeAnalysis(t *testing.T) {
	resp, err := client.TimeAnalysis("2013年二月二十八日下午四点三十分二十九秒")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("timestamp: %s type: %s\n", resp.Timestamp, resp.Type)
	}
}

func TestBosonNLPClient_SummaryAnalysis(t *testing.T) {
	resp, err := client.SummaryAnalysis("", strings.Replace(strings.Replace(`腾讯科技讯（刘亚澜）10月22日消息，
        前优酷土豆技术副总裁黄冬已于日前正式加盟芒果TV，出任CTO一职。
        资料显示，黄冬历任土豆网技术副总裁、优酷土豆集团产品技术副总裁等职务，
        曾主持设计、运营过优酷土豆多个大型高容量产品和系统。
        此番加入芒果TV或与芒果TV计划自主研发智能硬件OS有关。
        今年3月，芒果TV对外公布其全平台日均独立用户突破3000万，日均VV突破1亿，
        但挥之不去的是业内对其技术能力能否匹配发展速度的质疑，
        亟须招揽技术人才提升整体技术能力。
        芒果TV是国内互联网电视七大牌照方之一，之前采取的是“封闭模式”与硬件厂商预装合作，
        而现在是“开放下载”+“厂商预装”。
        黄冬在加盟土豆网之前曾是国内FreeBSD（开源OS）社区发起者之一，
        是研究并使用开源OS的技术专家，离开优酷土豆集团后其加盟果壳电子，
        涉足智能硬件行业，将开源OS与硬件结合，创办魔豆智能路由器。
        未来黄冬可能会整合其在开源OS、智能硬件上的经验，结合芒果的牌照及资源优势，
        在智能硬件或OS领域发力。
        公开信息显示，芒果TV在今年6月对外宣布完成A轮5亿人民币融资，估值70亿。
        据芒果TV控股方芒果传媒的消息人士透露，芒果TV即将启动B轮融资。`, "\n", "", -1), " ", "", -1), 0.2, false)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("summary:%s\n", resp)
	}
}

func TestGenerateTaskID(t *testing.T) {
	t.Logf("id:%s\n", GenerateID())
}

func TestBosonNLPClient_Cluster(t *testing.T) {
	task := Task{
		ID:   "abc",
		Text: "#新闻追踪#：【冀中星被移送检察院审查起诉】首都机场公安分局对冀中星爆炸案侦查终结，目前已移送朝阳检察院审查起诉。7月20日18时24分，冀中星在首都机场T3航站楼B口外引爆自制炸药。案发当天除冀中星左手腕因被炸截肢外，无其他人伤亡。7月29日，冀中星因涉嫌爆炸罪被批捕。http://t.cn/zQHjr0S",
	}
	tasks := []Task{}
	for i := 0; i < 100; i++ {
		tasks = append(tasks, task)
	}
	println(len(tasks))

	err := client.Cluster(tasks)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("ok")
	}
}

func TestBosonNLPClient_Comments(t *testing.T) {
	text1 := "#新闻追踪#：【冀中星被移送检察院审查起诉】首都机场公安分局对冀中星爆炸案侦查终结，目前已移送朝阳检察院审查起诉。7月20日18时24分，冀中星在首都机场T3航站楼B口外引爆自制炸药。案发当天除冀中星左手腕因被炸截肢外，无其他人伤亡。7月29日，冀中星因涉嫌爆炸罪被批捕。http://t.cn/zQHjr0S"
	text2 := "#豫广微新闻#【首都机场爆炸案嫌犯冀中星 移送检方审查起诉】 据报道，首都机场公安分局对冀中星爆炸案侦查终结，目前已移送朝阳检察院审查起诉。7月20日，山东籍男子冀中星在首都机场T3航站楼引爆自制炸药，案发当天除冀中星左手腕因被炸截肢外，无其他人伤亡"
	task := client.NewCommentsTask()
	task.AddText(text1)
	task.AddText(text2)

	resp, err := client.Comments(task)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v",resp)
	}
}
