package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar authAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLはErrNoAvatarURLを返すべきです")
	}
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url", testUrl}

}
