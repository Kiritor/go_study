package music

import "fmt"

/**
   播放接口
*/
type Player interface {
	Play(source string)
}

func Play(source,mtype string) {
	var p Player
    switch mtype {
		case "MP3":
		     p = &MP3Player{}
	    case "WAV":
		     p = &WAVPlayer{}
		default:
		     fmt.Println("不支持音乐类型",mtype)
		     return
	}
	p.Play(source)
}
