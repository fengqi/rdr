package dump

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/xueqiu/rdr/decoder"
	"os"
	"time"
)

// Keys is function for command `keys`
// output all keys in rdbfile(s) get from args
func Keys(c *cli.Context) {
	if c.NArg() < 1 {
		_, _ = fmt.Fprintln(c.App.ErrWriter, "keys requires at least 1 argument")
		_ = cli.ShowCommandHelp(c, "keys")
		return
	}

	noExpire := c.Bool("no-expire")

	for _, filepath := range c.Args() {
		if _, err := os.Lstat(filepath); err != nil {
			continue
		}

		keyDecoder := decoder.NewDecoder()
		go Decode(c, keyDecoder, filepath)
		fmt.Println("database,type,key,size_in_bytes,encoding,num_elements,len_largest_element,expiry")
		for e := range keyDecoder.Entries {
			if e.Expiry.IsZero() {
				fmt.Printf("%d,%s,%s,%d,%s,%d,%d,%s\n", e.Db, e.Type, e.Key, e.Bytes, e.Encoding, e.NumOfElem, e.LenOfLargestElem, "")
			} else {
				if noExpire && time.Now().After(e.Expiry) {
					continue
				}
				fmt.Printf("%d,%s,%s,%d,%s,%d,%d,%s\n", e.Db, e.Type, e.Key, e.Bytes, e.Encoding, e.NumOfElem, e.LenOfLargestElem, e.Expiry)
			}
		}
	}
}
