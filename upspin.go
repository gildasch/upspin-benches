package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"upspin.io/client"
	"upspin.io/config"
	_ "upspin.io/transports"
	"upspin.io/upspin"
)

type UpspinAccesser struct {
	upspin.Client
	root string
}

func NewUpspinAccesser(confPath string, root string) (UpspinAccesser, error) {
	cfg, err := config.FromFile(confPath)
	if err != nil {
		return UpspinAccesser{}, err
	}

	client := client.New(cfg)
	if client == nil {
		return UpspinAccesser{}, errors.New("client could be initialized")
	}

	for i := 0; i < 10; i++ {
		err = client.Delete(upspin.PathName(root + "/benches/" + strconv.Itoa(i)))
		if err != nil {
			fmt.Println("ignoring error on delete:", err)
		}
	}
	err = client.Delete(upspin.PathName(root + "/benches"))
	if err != nil {
		fmt.Println("ignoring error on delete:", err)
	}

	_, err = client.MakeDirectory(upspin.PathName(root + "/benches"))
	if err != nil {
		return UpspinAccesser{}, err
	}

	return UpspinAccesser{client, root}, nil
}

func (ua UpspinAccesser) persist(ts TestStruct) error {
	path := upspin.PathName(ua.root + "/" + strconv.Itoa(ts.ID))
	f, err := ua.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(ts)
}

func (ua UpspinAccesser) get(id int) (TestStruct, error) {
	path := upspin.PathName(ua.root + "/" + strconv.Itoa(id))
	f, err := ua.Open(path)
	if err != nil {
		return TestStruct{}, err
	}

	var ts TestStruct
	err = json.NewDecoder(f).Decode(&ts)
	return ts, err
}
