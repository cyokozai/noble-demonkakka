package main

import (
	"fmt"
	"os"
	"os/exec"
	"math/rand"
    "time"
	"flag"
)

const face string = "./demonkakka.txt"
const version string = "1.0.0"

func kakka_face() {
	_, err := exec.Command("clear").Output()
	if err != nil {
		fmt.Println("Error: ", err)

		return
	}

	file, err := os.Open(face)
	if err != nil {
		fmt.Println("Error: ", err)

		return
	}

	defer file.Close()

	buf := make([]byte, 1024)

	for {
		n, _ := file.Read(buf)

		if n == 0 {
			break
		}

		os.Stdout.Write(buf[:n])
	}
}

func message() {
	rand.Seed(time.Now().UnixNano())

	messages := []string{
		"お前を蝋人形にしてやろうか！",
		"お前は自分のガキに『人間』って名付けるのか？",
		"常識は破っても構わないが、非常識であってはならない",
		"そうだと思ったら行動せよ",
		"その人が自力ではどうすることもできないコンプレックスを使って人を罵ってはいけない",
		"リハーサルでは自分が世界で一番下手だと思え、ステージでは自分が世界で一番偉いと思え",
		"納得してないのに、納得したと言うな",
		"出身は地獄だ",
		"いい感じだ！まだまだイケるぞ！もっと良くなるぞ！いい声出してたぜ！",
		"吾輩が嫉妬を覚えたアーティストは数少ないが、そのうちの1人がDJ OZMAだ",
		"日本人の力士たちがなんで出世できないか”っていうひとつの別の理由に、忘れ去っている何か…大和魂みたいなものがあるんじゃないですか？",
		"昨今、蝋人形にして欲しい奴が多すぎて、地獄の蝋人形工場はそりゃあ大忙しでてんやわんやさ。よって、お前は蝋人形にはしてやらない。",
		"相撲は300年見てる。",
		"商売目的でやりやがってあんなの音楽じゃねェ！",
		"速くて激しいけれど、♪うまく歌ったってしょうがないじゃーん。上手に歌うことになんの意味があるんだ～。 ってのがパンク。",
		"エンターテインメントはお客さんに説教をする場ではない",
		"この日は銀行に行った。吾輩は悪魔だが、人間界で生活する限り、こういう雑務は避けられないものだ",
		"インプットが多岐にわたると、アウトプットが充実する。吾輩は日々感じている",
		"観にくるだけの客は許されない。黒ミサに参加しなくてはならない。会場に来たからには、ミサをつくるメンバーとしてそこにいなくてはならない",
		"パワーは減っているけど、技術は上がっている",
	}

	fmt.Println("フハハハハハハ！")
	fmt.Println(messages[rand.Intn(len(messages))])
}

func main() {
    flag.Parse()

	if flag.Arg(0) == "" {
		message()
    } else if flag.Arg(0) == "version" || flag.Arg(0) == "v" {
        fmt.Println("demonkakka version", version)
	} else if flag.Arg(0) == "face" || flag.Arg(0) == "f" {
		kakka_face()
	} else {
        fmt.Println("おい、何かがおかしいぞ！")
    }
}