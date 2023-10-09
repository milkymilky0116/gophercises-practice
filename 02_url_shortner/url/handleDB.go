package urlshortner

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/milkymilky0116/gophercises-practice/02_url_shortner/db"
	"go.etcd.io/bbolt"
)

type dbConfig struct {
	dbname     string
	bucketName string
	path       string
	url        string
}

func (app *AppConfig) CmdDB(w io.Writer, args []string) error {
	var c dbConfig
	fs := flag.NewFlagSet("DB Handler", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.dbname, "dbname", "test.db", "Set database name that you want to parse.")
	fs.StringVar(&c.bucketName, "bname", "urlshortner", "Set bucket name")
	fs.StringVar(&c.path, "path", "/dbtest", "Set path")
	fs.StringVar(&c.url, "url", "https://github.com", "Set url that you want to redirect")
	err := fs.Parse(args)
	if err != nil {
		switch {
		case err.Error() == flag.ErrHelp.Error():
			return flag.ErrHelp
		default:
			return err
		}
	}
	db, err := db.InitDB(c.dbname)
	if err != nil {
		return err
	}
	app.DB = db
	// //err = app.SetBucket(c.bucketName)
	// if err != nil {
	// 	return err
	// }
	err = app.SetPath(c.bucketName, c.path, c.url)
	if err != nil {
		return err
	}
	app.GetPath(c.bucketName)
	if err != nil {
		return err
	}
	return nil
}

func (app *AppConfig) SetBucket(bucketName string) error {
	err := app.DB.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (app *AppConfig) SetPath(bucketName string, path string, url string) error {
	err := app.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(path), []byte(url))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (app *AppConfig) GetPath(bucketName string) error {
	err := app.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for path, url := c.First(); path != nil; path, url = c.Next() {
			newPath := string(path)
			newUrl := string(url)
			fmt.Println(newPath)
			fmt.Println(newUrl)
			app.Mux.HandleFunc(newPath, func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, newUrl, http.StatusPermanentRedirect)
			})
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
