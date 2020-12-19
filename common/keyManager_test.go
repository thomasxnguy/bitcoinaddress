package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("keyManager", func() {
	var (
		keyManager = NewKeyManager("grace feed hazard slot orient plunge slush exhibit drip casino party polar nasty income afford", "test", false)
	)

	Describe("GetSegWitAddressForAccountAt", func() {
		Context("At index 0", func() {
			It("return address at index 0", func() {
				account, _ := keyManager.GetSegWitAddressForAccountAt(0)
				Expect(account).To(Equal("38dcnCjWsjFBBC5tmHFsQJMCHGSc2EZvRK"))
			})
		})
	})

	Describe("GetSegWitAddressForAccountAt", func() {
		Context("At index 1", func() {
			It("return address at index 1", func() {
				account, _ := keyManager.GetSegWitAddressForAccountAt(1)
				Expect(account).To(Equal("3PSwT9XCceh5xboKR5hEchsFxB8W2t3BRV"))
			})
		})
	})

	Describe("GetSegWitAddressForAccountAt", func() {
		Context("At index 2", func() {
			It("return address at index 2", func() {
				account, _ := keyManager.GetSegWitAddressForAccountAt(2)
				Expect(account).To(Equal("3572f7kiocdGYAtfDSiWeEEQyzqKkkmzeB"))
			})
		})
	})

	Describe("GetSegWitAddressForAccountAt", func() {
		Context("At index 3", func() {
			It("return address at index 3", func() {
				account, _ := keyManager.GetSegWitAddressForAccountAt(3)
				Expect(account).To(Equal("35Zu8pdfYhvqJbi1UUCw7YWRCtwKmLKpZv"))
			})
		})
	})

	Describe("GetSegWitAddressForAccountAt", func() {
		Context("At index 4", func() {
			It("return address at index 4", func() {
				account, _ := keyManager.GetSegWitAddressForAccountAt(4)
				Expect(account).To(Equal("3NwtvdnjfxPwEgwx7FBhBgW3ZeqDC2ZmqB"))
			})
		})
	})

	Describe("GetNativeSegWitAddressForAccountAt", func() {
		Context("At index 0", func() {
			It("return address at index 0", func() {
				account, _ := keyManager.GetNativeSegWitAddressForAccountAt(0)
				Expect(account).To(Equal("bc1qchtqmagpdmn8ceavrmg427dylpn0gfx3l2zsrl"))
			})
		})
	})

	Describe("GetNativeSegWitAddressForAccountAt", func() {
		Context("At index 1", func() {
			It("return address at index 1", func() {
				account, _ := keyManager.GetNativeSegWitAddressForAccountAt(1)
				Expect(account).To(Equal("bc1qr3dytznhzmp3hv7c4h7dt7mmqmy6qrs30ktaxt"))
			})
		})
	})

	Describe("GetNativeSegWitAddressForAccountAt", func() {
		Context("At index 2", func() {
			It("return address at index 2", func() {
				account, _ := keyManager.GetNativeSegWitAddressForAccountAt(2)
				Expect(account).To(Equal("bc1qm3lxavzd2vcz4as95lzs9snxmjuqr09rdfue23"))
			})
		})
	})

	Describe("GetNativeSegWitAddressForAccountAt", func() {
		Context("At index 3", func() {
			It("return address at index 3", func() {
				account, _ := keyManager.GetNativeSegWitAddressForAccountAt(3)
				Expect(account).To(Equal("bc1qq09vnyyv3mgxktx5lk80j5faah4l04f8uad3dk"))
			})
		})
	})

	Describe("GetNativeSegWitAddressForAccountAt", func() {
		Context("At index 4", func() {
			It("return address at index 4", func() {
				account, _ := keyManager.GetNativeSegWitAddressForAccountAt(4)
				Expect(account).To(Equal("bc1qwf7f5uqfn9eczhk3wmyd8arardvj5qhura8ugk"))
			})
		})
	})
})
