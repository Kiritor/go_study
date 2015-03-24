package musiclib

import (
    "errors"
)
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

func (m *MusicManager) Get(index int) (music *Music,err error){
	if index <0 ||index>=len(m.musics) {
		return nil,errors.New("索引越界")
	}
	return &m.musics[index],nil
}

/**
   根据歌名获取歌曲
*/

func (m *MusicManager) Find(name string) (music *Music) {
	if len(m.musics) == 0 {
		return nil
	}
	for _,m:=range m.musics {
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

/**
   删除歌曲
*/

func (m *MusicManager) Remove(index int) *Music {
	if index< 0 || index >=len(m.musics) {
		return nil
	}
	removeMusic :=&m.musics[index]
	if index<len(m.musics) -1 {
		//删除中间元素
		m.musics = append(m.musics[:index-1],m.musics[index+1:]...)
	}else if index==0 {
		//删除唯一一个元素(清空音乐库)，能够执行到这一步意味着只有一个元素
		m.musics = make([]Music,0)
	}else {
		m.musics=m.musics[:index-1]
	}
	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) *Music {
    if len(m.musics) == 0 {
        return nil
    }

    for i, v := range m.musics {
        if v.Name == name {
            return m.Remove(i)
        }
    }
    return nil
}
