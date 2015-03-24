package musiclib

import(
   "testing"
)

func TestOps (t *testing.T) {
    mm :=NewMusicManager()
	if mm == nil {
		t.Error("初始化音乐库失败!")
	}
	if mm.Len()!=0 {
		t.Error("初始化音乐库失败!不为空!")
	}
	m0 :=&Music{
		"1","突然好想你","五月天","D:music","MP3"}
	mm.Add(m0)
	//测试Add
	if mm.Len()!= 1 {
		t.Error("MusicManager.Add,失败!")
	}
	//测试Find
	m :=mm.Find(m0.Name)
	if m==nil {
		t.Error("MusicManager.Find,失败!")
	}
	m,err:=mm.Get(0)
	if m==nil {
		t.Error("MusicManager.Get,失败",err)
	}
	m = mm.Remove(0)
	if m==nil|| mm.Len()!=0 {
		t.Error("MusicManager.Remove,失败",err)
	}
}
