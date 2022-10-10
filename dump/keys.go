package dump

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/xueqiu/rdr/decoder"
)

// Keys is function for command `keys`
// output all keys in rdbfile(s) get from args
func Keys(c *cli.Context) {
	if c.NArg() < 1 {
		_, _ = fmt.Fprintln(c.App.ErrWriter, "keys requires at least 1 argument")
		_ = cli.ShowCommandHelp(c, "keys")
		return
	}
	for _, filepath := range c.Args() {
		keyDecoder := decoder.NewDecoder()
		go Decode(c, keyDecoder, filepath)
		fmt.Println("key,bytes,type,num_of_elem,len_of_large_elem")
		for e := range keyDecoder.Entries {
			fmt.Printf("%s,%d,%s,%d,%d\n", e.Key, e.Bytes, e.Type, e.NumOfElem, e.LenOfLargestElem)
		}
	}
}
