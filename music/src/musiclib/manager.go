package musiclib

import (
    "errors"
)
df
type MusicManager struct {
	musics []Music
}

/**
   初始化音乐库
*/
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music,0)}
}

/**
   音乐库长度
*/
func (m *MusicManager) Len() int {
	return len(m.musics)
}

/**
   根据索引得到音乐
*/

func (m *MusicManager) Get(index int) {
	if index <0 ||index>=len(m.musics) {
		return nil,errors.New("索引越界")
	}
	return &m.musics[index],nil
}

/**
   根据歌名获取歌曲
*/

func (m *MusicManager) Find(name string) {
	if len(m.musics) == 0 {
		return nil
	}
	for _,m:range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

/**
   添加歌曲
*/
func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics,*music)
}
