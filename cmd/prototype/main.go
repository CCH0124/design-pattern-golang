package main

import (
	"designpattern/pkg/creational/prototype"
	"log"
)

func main() {
	// 1. 初始化針對 *model.VirtualMachine 的泛型註冊表
	registry := prototype.NewPrototypeRegistry[*prototype.VirtualMachine]()

	// 2. 建立並註冊標準黃金映像檔 (Golden Image)
	goldenWebImage := &prototype.VirtualMachine{
		CPU:      2,
		MemoryGB: 4,
		OS:       "Ubuntu 22.04 LTS",
		Tags: map[string]string{
			"role": "web",
		},
	}
	registry.Register("web-server-base", goldenWebImage)

	log.Println("=== [Registry] Golden Template Registered ===")
	log.Printf("Original Golden Template: %s\n\n", goldenWebImage)

	// 3. 從 Registry 生產出 VM1 並進行客製化微調
	vm1, err := registry.Get("web-server-base")
	if err != nil {
		panic(err)
	}
	vm1.CPU = 4
	vm1.Tags["env"] = "prod"

	// 4. 從 Registry 生產出 VM2 並進行客製化微調
	vm2, err := registry.Get("web-server-base")
	if err != nil {
		panic(err)
	}
	vm2.Tags["env"] = "dev"
	vm2.Tags["tier"] = "frontend"

	// 5. 驗證結果
	log.Println("=== [Verification] Checking Status ===")
	log.Printf("Golden Template : %s\n", goldenWebImage)
	log.Printf("Instance VM1    : %s\n", vm1)
	log.Printf("Instance VM2    : %s\n", vm2)

	if &goldenWebImage.Tags == &vm1.Tags || &vm1.Tags == &vm2.Tags {
		log.Fatalf("[ERROR] Checking if memory address overlaps")
	} else {
		log.Println("[SUCCESS] Checking if memory address overlaps")
	}
}
