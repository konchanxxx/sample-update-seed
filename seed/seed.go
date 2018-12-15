package seed

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	// ErrLoad はseedデータのロードに失敗したエラー
	ErrLoad = errors.New("sample-update-seed-seed: Failed to load")
)

// Load はseedをfirestoreにロードする
func Load(ctx context.Context, fileName string) error {
	buf, err := readSeed(ctx, fileName)
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

func readSeed(ctx context.Context, fileName string) ([]byte, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucketName := os.Getenv("GCS_BUCKET_NAME")
	r, err := c.Bucket(bucketName).Object(fileName).NewReader(ctx)
	fmt.Println(bucketName)
	fmt.Println(fileName)
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
