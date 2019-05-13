package main

import (
	"github.com/kr/pretty"
	"github.com/naoki-kishi/feeder"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestEuphoFetch(t *testing.T) {
	// Set up mock server
	htmlFile, err := os.Open("eupho_test.html")
	if err != nil {
		t.Fatal("Failed to open test html file.")
	}
	bytes, _ := ioutil.ReadAll(htmlFile)
	response := &Response{
		Path:        "/news",
		ContentType: "text/html",
		Body:        string(bytes),
	}
	server := NewMockServer(response)
	defer server.Close()

	const DATE_LAYOUT = "2006.01.02"
	publishedString := []string{"2019.05.10", "2019.05.10", "2019.05.04", "2019.05.02", "2019.04.26"}
	published := []*time.Time{}

	for _, str := range publishedString {
		t, _ := time.Parse(DATE_LAYOUT, str)
		published = append(published, &t)
	}

	expected := &feeder.Items{
		[]*feeder.Item{{
			Title: "『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』大ヒット御礼舞台挨拶＆『リズと青い鳥』特別上映会開催決定!!",
			Link: &feeder.Link{
				Href: "http://anime-eupho.com/news/?id=348",
				Rel:  "",
			},
			Source: nil,
			Author: nil,
			Description: `<p>『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』大ヒットを記念して、大ヒット御礼舞台挨拶＆『リズと青い鳥』特別上映会を行うことが決定いたしました！！<br/>
                                    下記詳細をご確認の上、ぜひご参加ください！</p>

                                <p style="margin-top: 35px;">■舞台挨拶詳細■</p>

                                <dl>
                                    <dt>【実施日】</dt>
                                    <dd>6月2日（日）</dd>
                                </dl>

                                <dl style="margin-top: 50px;">
                                    <dt>【実施劇場】</dt>
                                    <dd><a href="https://www.smt-cinema.com/site/kyoto/" target="_blank">MOVIX京都</a></dd>
                                </dl>

                                <dl style="margin-top: 15px;">
                                    <dt>【実施時間】</dt>
                                    <dd>『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』8：30の回上映後／11：30の回上映前</dd>
                                </dl>

                                <dl style="margin-top: 15px;">
                                    <dt>【登壇者】</dt>
                                    <dd>黒沢ともよ、朝井彩加、豊田萌絵、安済知佳（予定・敬称略）</dd>
                                </dl>

                                <dl style="margin-top: 50px;">
                                    <dt>【実施劇場】</dt>
                                    <dd><a href="https://www.smt-cinema.com/site/marunouchi/" target="_blank">丸の内ピカデリー</a></dd>
                                </dl>

                                <dl style="margin-top: 15px;">
                                    <dt>【実施時間】</dt>
                                    <dd>『リズと青い鳥』16：20の回上映後／『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』19：25の回上映前</dd>
                                </dl>

                                <dl style="margin-top: 15px;">
                                    <dt>【登壇者】</dt>
                                    <dd>黒沢ともよ、朝井彩加、豊田萌絵、安済知佳（予定・敬称略）</dd>
                                </dl>
                                <dl style="margin-top: 50px;">
                                    <dt style="padding-left: 1em; text-indent: -1em;">※登壇者は、予告なく変更となる場合がございます。</dt>
                                    <dt style="padding-left: 1em; text-indent: -1em;">※『リズと青い鳥』特別上映会は、丸の内ピカデリー16：20回のみの実施です。<br/>MOVIX京都8：30回・11：30回、丸の内ピカデリー19：25回は『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』の上映となります。</dt>
                                </dl>

                                <dl style="margin-top: 50px;">
                                    <dt>【チケット販売】</dt>
                                    <dd>各劇場販売システムにて発売予定</dd>
                                    <dd style="padding-left: 1em; text-indent: -1em;">※詳細な販売方法等は劇場HPにてご確認ください。</dd>
                                </dl>

                                <dl style="margin-top: 25px;">
                                    <dt>【販売スケジュール】</dt>
                                    <dd>■販売：劇場HP・窓口・自動券売機にて販売</dd>
                                    <dd>〈MOVIX京都〉</dd>
                                    <dd>インターネット先行販売：5月10日（金）24：00（＝5月11日（土）0：00）以降順次</dd>
                                    <dd>劇場窓口・自動券売機販売：5月11日（土）劇場オープン時より</dd>
                                    <dd style="padding-left: 1em; text-indent: -1em;">※詳細な販売方法等は劇場HPにてご確認ください。</dd>

                                    <dd style="margin-top: 12.5px;">〈丸の内ピカデリー〉</dd>
                                    <dd>インターネット先行販売：5月13日（月）19：00以降順次</dd>
                                    <dd>劇場窓口：5月14日（火）劇場オープン以降（残席がある場合のみ）</dd>
                                    <dd style="padding-left: 1em; text-indent: -1em;">※販売方法等は劇場HPにてご確認ください。</dd>
                                </dl>

                                <dl style="margin-top: 25px;">
                                    <dt>【料金】</dt>
                                    <dd>2,000円均一（税込）</dd>
                                    <dd style="padding-left: 1em; text-indent: -1em;">※前売券・割引券等、利用不可</dd>
                                    <dd style="padding-left: 1em; text-indent: -1em;">※特別興行のため招待券・株主優待券等無料鑑賞、利用不可</dd>
                                </dl>

                                <ul style="margin-top: 25px; padding-left: 1em; text-indent: -1em;">
                                    <li>【注意事項】</li>
                                    <li>※開場・開演時間及び出演者・イベント内容は、予告なく変更となる場合がございます。予めご了承下さい。</li>
                                    <li>※特別興行の為、各種招待券等の無料鑑賞は、ご利用いただけません。</li>
                                    <li>※全席指定・定員入替制での上映となります。</li>
                                    <li>※転売目的でのご購入は固くお断りいたします。</li>
                                    <li>※いかなる事情が生じましても、ご購入・お引換後の鑑賞券の変更や払い戻しはできません。</li>
                                    <li>※場内でのカメラ（携帯電話を含む）・ビデオによる撮影・録画・録音等は、固くお断りいたします。保安上、入場時に手荷物検査を行う場合がございますこと、予めご了承ください。</li>
                                    <li>※会場内ではマスコミ各社の取材による撮影、記録撮影が行われ、テレビ・雑誌・ホームページ等にて放映・掲載される場合がございます。また、イベントの模様が後日販売されるDVD商品等に収録される場合があります。あらかじめご了承ください。お客様のこの催事における個人情報（肖像権）は、このイベントに入場されたことにより、上記に使用されるということにご同意頂けたものとさせて頂きます。</li>
                                </ul>


                            `,
			Id:      "/news/?id=348",
			Updated: nil,
			Created: published[0],
			Content: "",
		},
			{
				Title: "雑誌掲載のお知らせ",
				Link: &feeder.Link{
					Href: "http://anime-eupho.com/news/?id=347",
					Rel:  "",
				},
				Source: nil,
				Author: nil,
				Description: `<p>以下の4誌に『劇場版 響け！ユーフォニアム〜誓いのフィナーレ〜』の記事が掲載されています。<br/>是非ご覧ください！</p>
                                <ul style="margin-top: 20px;">
                                    <li>●4月30日(火)発売</li>
                                    <li>・ <a href="http://gs.dengeki.com/" target="_blank">電撃G’sマガジン 6月号</a></li>
                                    <li>・ <a href="https://hon.gakken.jp/magazine/08643" target="_blank">メガミマガジン 6月号</a></li>
                                </ul>
                                <ul style="margin-top: 20px;">
                                    <li>●5月10日(金)発売</li>
                                    <li>・ <a href="http://cho-animedia.jp/" target="_blank">月刊アニメディア 6月号</a></li>
                                    <li>・ <a href="https://webnewtype.com/magazine/" target="_blank">月刊ニュータイプ 6月号</a></li>
                                </ul>`,
				Id:      "/news/?id=347",
				Updated: nil,
				Created: published[1],
				Content: "",
			},
			{
				Title: "4週目～5週目入場者プレゼント配布決定!!",
				Link: &feeder.Link{
					Href: "http://anime-eupho.com/news/?id=346",
					Rel:  "",
				},
				Source: nil,
				Author: nil,
				Description: `<p>この度、追加の入場者プレゼントの配布が決定いたしました！</p>
                                <p>第4、5週目に本編の35mmフィルムから4コマを切り出したカットフィルムを、数量限定にて劇場への来場者にプレゼントいたします。</p>

                                <ul style="margin-top: 35px; padding-left: 1em; text-indent: -1em;">
                                    <li>・4週目 5/10～：35mmフィルム</li>
                                    <li style="margin-top: 10px;">・5週目 5/17～：35mmフィルム</li>
                                </ul>

                                <p style="margin-top: 20px;"><img alt="" src="/img/news/190509-novelty.jpg" width="700"/></p>

                                <ul style="margin-top: 35px; padding-left: 1em; text-indent: -1em;">
                                    <li>※ご入場時に、お一人様につき一点の特典をお渡し致します。</li>
                                    <li>※ランダム配布となりますので絵柄はお選びいただけません。</li>
                                    <li>※数量限定の為、無くなり次第終了となります。</li>
                                    <li>※チケット購入特典ではございません。ご入場を伴わない配布はお断りさせて頂きます。<br/>あしからずご了承くださいませ。</li>
                                    <li>※配布に関する特別な指定がある場合を除きます。</li>
                                </ul>`,
				Id:      "/news/?id=346",
				Updated: nil,
				Created: published[2],
				Content: "",
			},
			{
				Title: "Twitter感想投稿キャンペーン開催決定！",
				Link: &feeder.Link{
					Href: "http://anime-eupho.com/news/?id=344",
					Rel:  "",
				},
				Source: nil,
				Author: nil,
				Description: `<p>「劇場版 響け！ユーフォニアム～誓いのフィナーレ～」公開を記念して、Twitter感想投稿キャンペーンを開催いたします！<br/>
                                    ハッシュタグ「#誓いのフィナーレ感想」をつけて映画の感想を投稿すると、抽選で豪華な賞品をプレゼント！<br/>
                                    皆さまのご感想をお待ちしてます！</p>

                                <ul style="margin-top: 50px; padding-left: 1em; text-indent: -1em;">
                                    <li>●プレゼント</li>
                                    <li>・キャストサイン入りメインビジュアルポスター　5名様</li>
                                    <li>・キャストサイン入りキービジュアル第2弾ポスター　5名様</li>
                                    <li>・キャストサイン入りアフレコ台本　5名様</li>
                                </ul>

                                <ul style="margin-top: 25px; padding-left: 1em; text-indent: -1em;">
                                    <li>●参加方法</li>
                                    <li>1.「響け！ユーフォニアム」公式Twitterアカウント（<a href="https://twitter.com/anime_eupho" target="_blank">@anime_eupho</a>）をフォローしてください。</li>
                                    <li>2.ハッシュタグ「#誓いのフィナーレ感想」をつけて、映画を見た感想を投稿してください。</li>
                                    <li>3.キャンペーン応募完了です！</li>
                                </ul>

                                <p style="font-weight: bold;"><a href="https://twitter.com/intent/tweet?hashtags=%e8%aa%93%e3%81%84%e3%81%ae%e3%83%95%e3%82%a3%e3%83%8a%e3%83%bc%e3%83%ac%e6%84%9f%e6%83%b3" rel="”nofollow”" target="_blank">〈 #誓いのフィナーレ感想 をつけてツイートする 〉</a></p>

                                <ul style="margin-top: 25px;">
                                    <li>●投稿期間</li>
                                    <li>2019年４月26日（金）〜5月31日（金）まで</li>
                                </ul>

                                <ul style="margin-top: 25px;">
                                    <li>●プレゼント発送予定時期</li>
                                    <li>2019年6月以降</li>
                                </ul>

                                <ul style="margin-top: 25px;">
                                    <li>●当選発表</li>
                                    <li>厳選なる抽選の上、ご当選者様には公式アカウントよりTwitterダイレクトメッセージでご案内いたします。</li>
                                    <li>当選発表はダイレクトメッセージのご連絡をもって代えさせていただきます。</li>
                                </ul>

                                <ul style="margin-top: 25px; padding-left: 1em; text-indent: -1em;">
                                    <li>●注意事項</li>
                                    <li>・Twitterアカウントを非公開にしている場合、リツイートを確認することができないため、応募対象外とさせていただきます。</li>
                                    <li>・当選者にのみダイレクトメッセージで当選通知をお送りするため、受け取れるように設定をお願いいたします。</li>
                                    <li>・当選時にダイレクトメッセージにてお伺いした氏名・送付先住所については、キャンペーン以外の目的では使用いたしません。</li>
                                    <li>・抽選結果に関するお問い合わせにはお答えできません。</li>
                                    <li>・また、本キャンペーンの内容は予告なく変更または中止となる場合がございますので、予めご了承ください。</li>
                                    <li>・当選後の連絡が取れないなどの場合は、当選を取り消しとさせていただく場合があります。</li>
                                    <li>・発送後の商品破損、紛失につきましては責任を負いかねますので、予めご了承ください。</li>
                                </ul>`,
				Id:      "/news/?id=344",
				Updated: nil,
				Created: published[3],
				Content: "",
			},
			{
				Title: "『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』スタッフトーク付き上映会　実施決定！",
				Link: &feeder.Link{
					Href: "http://anime-eupho.com/news/?id=345",
					Rel:  "",
				},
				Source: nil,
				Author: nil,
				Description: `<p>この度、『劇場版 響け！ユーフォニアム～誓いのフィナーレ～』のスタッフトーク付き上映会の実施が決定致しました！</p>
                                <p>登壇者は、石原立也監督をはじめ作品制作の中心を担ったメンバーとなっております。是非この機会に劇場へお越しください。</p>
                                <p>以下詳細となっております。</p>

                                <dl style="margin-top: 25px;">
                                    <dt>【開催日時】</dt>
                                    <dd>2019年5月10日(金)　19:00～の回、上映終了後スタッフトーク</dd>
                                </dl>
                                <dl style="margin-top: 25px;">
                                    <dt>【登壇者】※予定</dt>
                                    <dd>石原立也（監督）</dd>
                                    <dd>池田晶子（キャラクターデザイン・総作画監督）</dd>
                                    <dd>西屋太志（総作画監督）</dd>
                                    <dd style="margin-top: 15px;">※登壇者は、予告なしに変更する場合がございます。</dd>
                                </dl>
                                <dl style="margin-top: 25px;">
                                    <dt>【劇場】</dt>
                                    <dd><a href="https://www.smt-cinema.com/site/kyoto/" target="_blank">MOVIX京都</a></dd>
                                </dl>
                                <dl style="margin-top: 25px;">
                                    <dt>【チケット料金】</dt>
                                    <dd>金額：劇場通常料金</dd>
                                    <dd>※特別興行につき、各種招待券等無料でのご鑑賞は頂けません。</dd>
                                </dl>
                                <dl style="margin-top: 25px;">
                                    <dt>【チケット販売方法】</dt>
                                    <dd>・インターネット先行販売：</dd>
                                    <dd>4月26日(金)24:00（＝4月27日（土）0：00）</dd>
                                    <dd>以降順次、劇場ホームページにて販売順次開始</dd>
                                    <dd style="margin-top: 15px;">・劇場窓口・自動券売機販売：</dd>
                                    <dd>4月27日(土)劇場オープンより窓口にて販売開始（残席がある場合のみ）</dd>
                                    <dd>※販売方法等は劇場HPにてご確認ください。</dd>
                                </dl>
                                <ul style="margin-top: 25px; padding-left: 1em; text-indent: -1em;">
                                    <li>◆注意事項</li>
                                    <li>※特別興行につき、招待券・株主優待券などはご使用いただけません。</li>
                                    <li>※お電話でのご予約は承っておりません。</li>
                                    <li>※転売・転用目的の購入は固くお断りいたします。</li>
                                    <li>※ご購入後の払い戻し、座席変更は承っておりません。</li>
                                    <li>※いかなる場合においても途中入場はお断りさせていただきますので、ご了承ください。</li>
                                    <li>※登壇者は予告なく変更する場合がございます。予めご了承ください。</li>
                                    <li>※場内での撮影（カメラ付携帯電話を含む）および録音は固くお断りしております。</li>
                                    <li>※撮影機材の持ち込みもご遠慮頂きますようお願いいたします。</li>
                                    <li>※当日マスコミ取材が入る場合がございます。予めご了承ください。</li>
                                </ul>`,
				Id:      "/news/?id=345",
				Updated: nil,
				Created: published[4],
				Content: "",
			}}}

	fetcher := NewEuphoFetcher(server.URL + "/news")
	got, err := fetcher.Fetch()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*expected, *got) {
		diffs := pretty.Diff(*expected, *got)
		t.Log(pretty.Println(diffs))
		t.Error("Failed to convert Html to Item.")
	}
}
