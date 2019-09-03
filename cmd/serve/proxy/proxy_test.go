package proxy

import (
	"strings"
	"testing"

	_ "github.com/sdhealth/rclone/backend/local"
	"github.com/sdhealth/rclone/fs/config/configmap"
	"github.com/sdhealth/rclone/fs/config/obscure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestRun(t *testing.T) {
	opt := DefaultOpt
	cmd := "go run proxy_code.go"
	opt.AuthProxy = cmd
	p := New(&opt)

	t.Run("Normal", func(t *testing.T) {
		config, err := p.run(map[string]string{
			"type": "ftp",
			"user": "me",
			"pass": "pass",
			"host": "127.0.0.1",
		})
		require.NoError(t, err)
		assert.Equal(t, configmap.Simple{
			"type":  "ftp",
			"user":  "me-test",
			"pass":  "pass",
			"host":  "127.0.0.1",
			"_root": "",
		}, config)
	})

	t.Run("Error", func(t *testing.T) {
		config, err := p.run(map[string]string{
			"error": "potato",
		})
		assert.Nil(t, config)
		require.Error(t, err)
		require.Contains(t, err.Error(), "potato")
	})

	t.Run("Obscure", func(t *testing.T) {
		config, err := p.run(map[string]string{
			"type":     "ftp",
			"user":     "me",
			"pass":     "pass",
			"host":     "127.0.0.1",
			"_obscure": "pass,user",
		})
		require.NoError(t, err)
		config["user"] = obscure.MustReveal(config["user"])
		config["pass"] = obscure.MustReveal(config["pass"])
		assert.Equal(t, configmap.Simple{
			"type":     "ftp",
			"user":     "me-test",
			"pass":     "pass",
			"host":     "127.0.0.1",
			"_obscure": "pass,user",
			"_root":    "",
		}, config)
	})

	const testUser = "testUser"
	const testPass = "testPass"

	t.Run("call", func(t *testing.T) {
		// check cache empty
		assert.Equal(t, 0, p.vfsCache.Entries())
		defer p.vfsCache.Clear()

		passwordBytes := []byte(testPass)
		value, err := p.call(testUser, testPass, passwordBytes)
		require.NoError(t, err)
		entry, ok := value.(cacheEntry)
		require.True(t, ok)

		// check hash is correct in entry
		err = bcrypt.CompareHashAndPassword(entry.pwHash, passwordBytes)
		require.NoError(t, err)
		require.NotNil(t, entry.vfs)
		f := entry.vfs.Fs()
		require.NotNil(t, f)
		assert.Equal(t, "proxy-"+testUser, f.Name())
		assert.True(t, strings.HasPrefix(f.String(), "Local file system"))

		// check it is in the cache
		assert.Equal(t, 1, p.vfsCache.Entries())
		cacheValue, ok := p.vfsCache.GetMaybe(testUser)
		assert.True(t, ok)
		assert.Equal(t, value, cacheValue)
	})

	t.Run("Call", func(t *testing.T) {
		// check cache empty
		assert.Equal(t, 0, p.vfsCache.Entries())
		defer p.vfsCache.Clear()

		vfs, vfsKey, err := p.Call(testUser, testPass)
		require.NoError(t, err)
		require.NotNil(t, vfs)
		assert.Equal(t, "proxy-"+testUser, vfs.Fs().Name())
		assert.Equal(t, testUser, vfsKey)

		// check it is in the cache
		assert.Equal(t, 1, p.vfsCache.Entries())
		cacheValue, ok := p.vfsCache.GetMaybe(testUser)
		assert.True(t, ok)
		cacheEntry, ok := cacheValue.(cacheEntry)
		assert.True(t, ok)
		assert.Equal(t, vfs, cacheEntry.vfs)

		// Test Get works while we have something in the cache
		t.Run("Get", func(t *testing.T) {
			assert.Equal(t, vfs, p.Get(testUser))
			assert.Nil(t, p.Get("unknown"))
		})

		// now try again from the cache
		vfs, vfsKey, err = p.Call(testUser, testPass)
		require.NoError(t, err)
		require.NotNil(t, vfs)
		assert.Equal(t, "proxy-"+testUser, vfs.Fs().Name())
		assert.Equal(t, testUser, vfsKey)

		// check cache is at the same level
		assert.Equal(t, 1, p.vfsCache.Entries())

		// now try again from the cache but with wrong password
		vfs, vfsKey, err = p.Call(testUser, testPass+"wrong")
		require.Error(t, err)
		require.Contains(t, err.Error(), "incorrect password")
		require.Nil(t, vfs)
		require.Equal(t, "", vfsKey)

		// check cache is at the same level
		assert.Equal(t, 1, p.vfsCache.Entries())

	})

}
