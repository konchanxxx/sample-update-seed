package seed

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
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

	collectionName := strings.Replace(fileName, ".yaml", "", 1)
	err = putMulti(ctx, collectionName, seeds)
	if err != nil {
		return errors.Wrap(err, ErrLoad.Error())
	}
	log.Printf("Updated seed data: %#v", seeds)

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

func putMulti(ctx context.Context, collectionName string, seeds []map[string]string) error {
	projectID := os.Getenv("GCP_PROJECT_ID")
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	for _, seed := range seeds {
		_, err := c.Collection(collectionName).Doc(seed["id"]).Set(ctx, seed)
		if err != nil {
			return err
		}
	}

	return nil
}
