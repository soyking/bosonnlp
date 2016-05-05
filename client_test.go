package bosonnlp

import "testing"

var client = NewBosonNLPClient("your token")

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
