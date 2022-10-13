package dump

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/xueqiu/rdr/decoder"
	"github.com/xueqiu/rdr/utils"
	"os"
	"regexp"
	"sync"
	"time"
)

// HashFields 统计 hash 字段出现的次数
func HashFields(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		_, _ = fmt.Fprintln(c.App.ErrWriter, "keys requires at least 1 argument")
		_ = cli.ShowCommandHelp(c, "hash-fields")
		return
	}

	compile, err := regexp.Compile(args[0])
	if err != nil {
		_ = fmt.Errorf("compile %s err: %s\n", args[0], err.Error())
		return
	}

	noExpire := c.Bool("no-expire")

	wg := sync.WaitGroup{}
	wg.Add(c.NArg() - 1)
	counter := make(map[string]int, 0)
	keys := make(map[string]int, 0)
	for _, file := range args[1:] {
		if _, err := os.Lstat(file); err != nil {
			continue
		}

		rdbDecoder := decoder.NewDecoder()
		go Decode(c, rdbDecoder, file)
		for e := range rdbDecoder.Entries {
			if noExpire && !e.Expiry.IsZero() && time.Now().After(e.Expiry) {
				continue
			}

			if e.Type == "hash" {
				if compile.MatchString(e.Key) {
					keys[utils.KeyPrefixDistinct(e.Key)] += 1
					for _, field := range e.HashMembers {
						counter[field] += 1
					}
				}
			}
		}
		wg.Done()
	}
	wg.Wait()

	for key, count := range keys {
		fmt.Printf("pattern: %s, key: %s, count: %d\n", args[0], key, count)
	}
	fmt.Println()
	for field, count := range counter {
		fmt.Printf("pattern: %s, field: %s, count: %d\n", args[0], field, count)
	}
}
