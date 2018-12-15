package seed

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	// ErrLoad はseedデータのロードに失敗したエラー
	ErrLoad = errors.New("sample-update-seed-seed: Failed to load")
)

// Load はseedをfirestoreにロードする
func Load(ctx context.Context, bucketID string, fileName string) error {
	buf, err := readSeed(ctx, bucketID, fileName)
	if err != nil {
		return errors.Wrap(err, ErrLoad.Error())
	}

	seeds, err := mapSeed(buf)
	if err != nil {
		return errors.Wrap(err, ErrLoad.Error())
	}

	log.Printf("Output seed data: %#v", seeds)

	// TODO: firestoreにseedを反映する処理追加

	return nil
}

func readSeed(ctx context.Context, bucketID string, fileName string) ([]byte, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	r, err := c.Bucket(bucketID).Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func mapSeed(buf []byte) ([]map[string]string, error) {
	src := make([]map[string]string, 0)
	err := yaml.Unmarshal(buf, &src)
	if err != nil {
		return nil, err
	}

	return src, nil
}
